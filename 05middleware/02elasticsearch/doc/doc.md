# 安装
```yaml
version: '3'
services:
  elasticsearch:
    image: elasticsearch:8.1.1
    container_name: es
    volumes: 
      - /home/xuwuruoshui/go-video/es:/usr/share/elasticsearch/plugins
    ports:
      - 9200:9200
      - 9300:9300
    environment: 
      - discovery.type=single-node
      - ELASTIC_PASSWORD=root
  kibana:
    image: kibana:8.1.0
    container_name: kibana
    depends_on:
      - elasticsearch
    ports:
      - 5601:5601
```

# xpack访问
```shell
docker cp es:/usr/share/elasticsearch/config/certs/http_ca.crt .

# curl访问
curl --cacert http_ca.crt -u elastic https://localhost:9200

# 浏览器(用户名:elastic 密码:root)
https://localhost:9200

# 验证码
# 按界面提示到elasticsearch容器中执行一个命令获取密钥,填入kibana
# 然后再到kibana容器中执行一个命令获取验证码
```

# ik分词器
```shell
# 下载放到容器内的 /usr/share/elasticsearch/plugins
# 解压并删除zip文件
https://github.com/medcl/elasticsearch-analysis-ik/releases
```
