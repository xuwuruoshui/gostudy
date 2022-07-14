# 安装gogs

```sh
docker volume create --name gogs-data
docker run --name=gogs -p 10022:22 -p 10880:3000 -v gogs-data:/data gogs/gogs
```
http://xxxxxxxx:10880, 配置好数据库,想简单就sqlite
http://xxxxxxxx:10880/user/settings/applications, 去授权应用获取下一步要的token
可以随意创建一个仓库,比如drone-test

tip: 本地的地址ip,需要在`/data/gogs/conf/app.ini`加一个
```ini
[security]
LOCAL_NETWORK_ALLOWLIST = ip或domain
```



# 安装drone
```
docker run \
  --volume=/var/lib/drone:/data \
  --env=DRONE_AGENTS_ENABLED=true \
  --env=DRONE_GOGS_SERVER=http://yourip:10880 \
  --env=DRONE_RPC_SECRET=super-duper-secret \
  --env=DRONE_SERVER_HOST=yourip:8080 \
  --env=DRONE_GIT_ALWAYS_AUTH=true \
  --env=DRONE_SERVER_PROTO=http \
  --env=DRONE_USER_CREATE=username:{your admin name},admin:true \
  --publish=8080:80 \
  --publish=8443:443 \
  --restart=always \
  --detach=true \
  --name=drone \
  drone/drone

```

# 安装drone runner
```sh
  docker run --detach \
  --volume=/var/run/docker.sock:/var/run/docker.sock \
  --env=DRONE_TMATE_ENABLED=true \
  --env=DRONE_RPC_PROTO=http \
  --env=DRONE_RPC_HOST=yourip:8080 \
  --env=DRONE_RPC_SECRET=super-duper-secret \
  --env=DRONE_RUNNER_CAPACITY=2 \
  --env=DRONE_RUNNER_NAME=my-first-runner \
  --publish=3000:3000 \
  --restart=always \
  --name=runner \
  drone/drone-runner-docker
```
访问 drone.local.com:8080, 可以看到drone-test


# 编写相关文件,提交到gogs

.drone.yml(drone自动化配置)
```yml
kind: pipeline
type: docker
name: build
steps:
- name: 1.编译文件
  image: golang:alpine
  pull: if-not-exists # always never
  commands:
    - export GOARCH=amd64
    - export GOOS=linux
    - export GO111MODULE=auto
    - export GOPROXY=https://goproxy.cn
    - go build -o main
    - ls
- name: 3.传输文件
  image: appleboy/drone-scp
  settings:
    host: 120.78.159.42
    username: root
    password: 
      from_secret: ssh_password
    port: 22
    target: /app/drone-test
    #tmp目录是远程主机的
    source: ./
    #需要传输的文件夹，当前目录
- name: 3.打包docker镜像
  pull: if-not-exists
  image: appleboy/drone-ssh
  settings:
    host: 120.78.159.42
    username: root
    password: 
      # 从drone仓库配置中秘密空间读取密码
      from_secret: ssh_password
    port: 22
    script:
      - cd /app/drone-test
      - chmod +x deploy.sh
      - ./deploy.sh
```

deploy.sh(部署脚本)
```sh
# !/bin/bash
echo =======暂停容器=======
docker stop `docker ps -a | grep drone-test | awk '{print $1}' `
echo =======暂停旧容器和镜像=======
docker rm -f `docker ps -a | grep drone-test | awk '{print $1}' `
docker rmi `docker images | grep drone-test | awk '{print $3}' `
echo =======开始构建新镜像=======cd /app/drone-test
docker build -t drone-test:latest .
echo =======开始部署应用=======
docker run  -p 4567:4567 --name drone-test  -d drone-test:latest 
echo =======部署成功=======
```

goweb(简单的服务器)
```go
package main

import (
	"fmt"
	"net/http"
)

/**
* @creator: xuwuruoshui
* @date: 2022-07-06 15:51:58
* @content:
 */

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4567", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ff")
	w.Write([]byte("ffff"))
}
```

dockerfile(镜像文件)
```dockerfile
FROM golang:alpine

RUN  mkdir -p /drone-test
ADD  ./main /drone-test
WORKDIR /drone-test


EXPOSE 4567

ENTRYPOINT ["/drone-test/main"]
```