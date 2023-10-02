# Setting up local kafka environment
## Get kafka
1. Download the latest kafka and extract it 
2. navigate in to the extracted directory

## Start the Kafka Environment 
### Kafka with ZooKeeper
1. Run the following commands to lauch a default service
2. `bin/zookeeper-server-start.sh config/zookeeper.properties`
3. `bin/kafka-server-start.sh config/server.properties`

## Create a topic
1. Run the following command
2. `bin/kafka-topics.sh --create --topic {topic name} --bootstrap-server {address}`