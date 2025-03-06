
## 介绍
使用go模拟logstash实现类似功能，完成数据的ETL(消费、处理、写入)

## 安装


## 运行


## 配置



### kafka
```yaml
inputs:
  - name: kafka
    config:
      topic: "topic_testxxx"
      kafka_config:
        bootstrap.servers: "10.2.2.2:8080,10.2.2.3:8080"
        group.id: "gostash-biflow-test"
        broker.address.family: "v4"
        session.timeout.ms: 6000
        enable.auto.commit: true
        enable.auto.offset.store: false
        auto.offset.reset: "earliest"
          #security.protocol: "SASL_PLAINTEXT"
          #sasl:
          #mechanisms: "PLAIN",
        #username: kafkaConf.Username,
        #password: kafkaConf.Password,
#  参考 https://github.com/edenhill/librdkafka/tree/master/CONFIGURATION.md

```

## 插件

### input插件

### output插件
### 注意事项
- output Start()实现的时候一定要判断，event, ok := <-inChan，否则外部channel关闭就会无限收收到0值，要处理好。