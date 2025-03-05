
# 介绍
使用go模拟logstash实现类似功能，完成数据的ETL(消费、处理、写入)

# 安装


# 运行


# 配置

# 插件

## input插件

## output插件
### 注意事项
- output Start()实现的时候一定要判断，event, ok := <-inChan，否则外部channel关闭就会无限收收到0值，要处理好。