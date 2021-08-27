# Toolbox
在线连接MySql、Redis、Zookeeper、Kafka、RabbitMQ、Elasticsearch等

# 打包发布

### 前端
前端工程使用vue开发，打开html目录，编译项目文件，拷贝dist目录下所有问题至web\src\main\webapp目录（删除原有文件）

### 服务端
服务端使用java开发，包管理使用maven，在根目录install即可，将会生成release目录，该目录下可以执行放至服务器运行