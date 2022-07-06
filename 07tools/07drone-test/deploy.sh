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
