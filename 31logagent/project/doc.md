# 日志系统(整合kafka)
## 安装kafka和zookeeper
```shell
# 添加一下配置
version: "3"
services:
  zookeeper:
    image: 'bitnami/zookeeper:latest'
    ports:
      - '2181:2181'
    environment:
      - ALLOW_ANONYMOUS_LOGIN=yes
  kafka:
    image: 'bitnami/kafka:latest'
    ports:
      - '9092:9092'
    environment:
      - KAFKA_BROKER_ID=1
      - KAFKA_CFG_LISTENERS=PLAINTEXT://:9092
      - KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://192.168.0.110:9092
      - KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181
      - ALLOW_PLAINTEXT_LISTENER=yes
    depends_on:
      - zookeeper
      
 # 启动kafka和zookeeper(默认kafka的配置路径为/opt/bitnami/kafka,消息存储路径为/bitnami/kafka/data)
 docker-compose up -d
 
 # 测试,从kafka自带的消费者中读取
cd /opt/bitnami/kafka/data
kafka-console-consumer.sh --bootstrap-server=192.168.0.110:9092 --topic=web_log --from-beginning
```




## etcd安装
```shell
# docker-compose
version: "3.0"

networks:
  etcd-net:           # 网络
    driver: bridge    # 桥接模式

volumes:
  etcd1_data:         # 挂载到本地的数据卷名
    driver: local
  etcd2_data:
    driver: local
  etcd3_data:
    driver: local
###
### etcd 其他环境配置见：https://doczhcn.gitbook.io/etcd/index/index-1/configuration
###
services:
  etcd1:
    image: bitnami/etcd:latest  # 镜像
    container_name: etcd1       # 容器名 --name
    restart: always             # 总是重启
    networks:
      - etcd-net                # 使用的网络 --network
    ports:                      # 端口映射 -p
      - "20000:2379"
      - "20001:2380"
    environment:                # 环境变量 --env
      - ALLOW_NONE_AUTHENTICATION=yes                       # 允许不用密码登录
      - ETCD_NAME=etcd1                                     # etcd 的名字
      - ETCD_INITIAL_ADVERTISE_PEER_URLS=http://etcd1:2380  # 列出这个成员的伙伴 URL 以便通告给集群的其他成员
      - ETCD_LISTEN_PEER_URLS=http://0.0.0.0:2380           # 用于监听伙伴通讯的URL列表
      - ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379         # 用于监听客户端通讯的URL列表
      - ETCD_ADVERTISE_CLIENT_URLS=http://etcd1:2379        # 列出这个成员的客户端URL，通告给集群中的其他成员
      - ETCD_INITIAL_CLUSTER_TOKEN=etcd-cluster             # 在启动期间用于 etcd 集群的初始化集群记号
      - ETCD_INITIAL_CLUSTER=etcd1=http://etcd1:2380,etcd2=http://etcd2:2380,etcd3=http://etcd3:2380 # 为启动初始化集群配置
      - ETCD_INITIAL_CLUSTER_STATE=new                      # 初始化集群状态
    volumes:
      - etcd1_data:/bitnami/etcd                            # 挂载的数据卷

  etcd2:
    image: bitnami/etcd:latest
    container_name: etcd2
    restart: always
    networks:
      - etcd-net
    ports:
      - "20002:2379"
      - "20003:2380"
    environment:
      - ALLOW_NONE_AUTHENTICATION=yes
      - ETCD_NAME=etcd2
      - ETCD_INITIAL_ADVERTISE_PEER_URLS=http://etcd2:2380
      - ETCD_LISTEN_PEER_URLS=http://0.0.0.0:2380
      - ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379
      - ETCD_ADVERTISE_CLIENT_URLS=http://etcd2:2379
      - ETCD_INITIAL_CLUSTER_TOKEN=etcd-cluster
      - ETCD_INITIAL_CLUSTER=etcd1=http://etcd1:2380,etcd2=http://etcd2:2380,etcd3=http://etcd3:2380
      - ETCD_INITIAL_CLUSTER_STATE=new
    volumes:
      - etcd2_data:/bitnami/etcd

  etcd3:
    image: bitnami/etcd:latest
    container_name: etcd3
    restart: always
    networks:
      - etcd-net
    ports:
      - "20004:2379"
      - "20005:2380"
    environment:
      - ALLOW_NONE_AUTHENTICATION=yes
      - ETCD_NAME=etcd3
      - ETCD_INITIAL_ADVERTISE_PEER_URLS=http://etcd3:2380
      - ETCD_LISTEN_PEER_URLS=http://0.0.0.0:2380
      - ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379
      - ETCD_ADVERTISE_CLIENT_URLS=http://etcd3:2379
      - ETCD_INITIAL_CLUSTER_TOKEN=etcd-cluster
      - ETCD_INITIAL_CLUSTER=etcd1=http://etcd1:2380,etcd2=http://etcd2:2380,etcd3=http://etcd3:2380
      - ETCD_INITIAL_CLUSTER_STATE=new
    volumes:
      - etcd3_data:/bitnami/etcd
      
# 启动
docker-compose up -d
docker exec -it etcd1 /bin/bash

# 测试
etcdctl put a "hello world"
etcdctl get a

# 远程测试
# 先下载etcd 
# https://github.com/etcd-io/etcd/releases/
etcdctl.exe --endpoints=http://192.168.0.110:20004 put b "666"
etcdctl.exe --endpoints=http://192.168.0.110:20004 get b
etcdctl.exe --endpoints=http://192.168.0.110:20004 del b
```

## elasticsearch
```shell
# 设置vm.max_map_count
vi /etc/sysctl.conf
vm.max_map_count=262144

sysctl -w vm.max_map_count=262144

# 需要先创建以下文件夹及文件
mkdir /home/xuwuruoshui/data/elasticsearch/conf
vi /home/xuwuruoshui/data/elasticsearch/conf/jvm.options
-Xmx1g
-Xms1g

mkdir /home/xuwuruoshui/data/elasticsearch/data/data01
mkdir /home/xuwuruoshui/data/elasticsearch/data/data02
mkdir /home/xuwuruoshui/data/elasticsearch/data/data03


# docker-compose
version: '2.2'
services:
  es01:
    image: docker.elastic.co/elasticsearch/elasticsearch:7.15.2
    container_name: es01
    environment:
      - node.name=es01
      - cluster.name=es-docker-cluster
      - discovery.seed_hosts=es02,es03
      - cluster.initial_master_nodes=es01,es02,es03
      - bootstrap.memory_lock=true
    ulimits:
      memlock:
        soft: -1
        hard: -1
    volumes:
      - /home/xuwuruoshui/data/elasticsearch/data/data01:/usr/share/elasticsearch/data
      - /home/xuwuruoshui/data/elasticsearch/conf:/usr/share/elasticsearch/config/jvm.options.d/
    ports:
      - 9200:9200
    networks:
      - elastic
  es02:
    image: docker.elastic.co/elasticsearch/elasticsearch:7.15.2
    container_name: es02
    environment:
      - node.name=es02
      - cluster.name=es-docker-cluster
      - discovery.seed_hosts=es01,es03
      - cluster.initial_master_nodes=es01,es02,es03
      - bootstrap.memory_lock=true
    ulimits:
      memlock:
        soft: -1
        hard: -1
    volumes:
      - /home/xuwuruoshui/data/elasticsearch/data/data02:/usr/share/elasticsearch/data
      - /home/xuwuruoshui/data/elasticsearch/conf:/usr/share/elasticsearch/config/jvm.options.d/
    networks:
      - elastic
  es03:
    image: docker.elastic.co/elasticsearch/elasticsearch:7.15.2
    container_name: es03
    environment:
      - node.name=es03
      - cluster.name=es-docker-cluster
      - discovery.seed_hosts=es01,es02
      - cluster.initial_master_nodes=es01,es02,es03
      - bootstrap.memory_lock=true
    ulimits:
      memlock:
        soft: -1
        hard: -1
    volumes:
      - /home/xuwuruoshui/data/elasticsearch/data/data03:/usr/share/elasticsearch/data
      - /home/xuwuruoshui/data/elasticsearch/conf:/usr/share/elasticsearch/config/jvm.options.d/
    networks:
      - elastic

volumes:
  data01:
    driver: local
  data02:
    driver: local
  data03:
    driver: local

networks:
  elastic:
    driver: bridge
```