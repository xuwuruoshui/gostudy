# Docker-compose安装rabbitmq
```yaml
version: '3'
services:
    rabbitmq:
        image: rabbitmq:latest
        container_name: rabbitmq
        restart: always
        hostname: myRabbitmq
        ports:
          - 15672:15672
          - 5672:5672
        # volumes:
        #   - ./data:/var/lib/rabbitmq
        environment:
          - RABBITMQ_DEFAULT_USER=root
          - RABBITMQ_DEFAULT_PASS=root
        command:
          /bin/bash -c "redis-server /usr/local/etc/redis/redis.conf "
```

# 配置
```shell
# 开启web界面
rabbitmq-plugins enable rabbitmq_shovel rabbitmq_shovel_management
# 显示channel页
echo management_agent.disable_metrics_collector = false >
/etc/rabbitmq/conf.d/management_agent.disable_metrics_collector.conf
```




