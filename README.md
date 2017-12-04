# sf-code-marathon
the code marathon program

## prepare for development environment

### 1. mysql
run mysql service
>sudo docker pull mysql:5.7

>sudo docker run -p 3306:3306 --name marathon-mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7

connect to mysql
>sudo docker run -it --rm --link marathon-mysql:mysql mysql:5.7 sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p123456'

>sudo docker container start marathon-mysql

```sql
create database marathon;
use marathon;
drop table if exists `m_user`;
create table `m_user`(
    `id` int primary key not null auto_increment, 
    `name` varchar(16)
)Engine=Innodb DEFAULT CHARSET=utf8;
```
### 2. redis
run redis service
>sudo docker pull redis:3.2

>sudo docker run -p 6379:6379 --name marathon-redis -d redis:3.2 redis-server --appendonly yes

>sudo docker container start marathon-redis

connect to redis
>sudo docker run -it --rm --link marathon-redis:redis redis:3.2 redis-cli -h redis -p 6379

### 3. mongodb
run mongodb service
>sudo docker pull mongo:3.4

>sudo docker run -p 27017:27017 --name marathon-mongodb -d mongo:3.4

>sudo docker container start marathon-mongodb

connect to mongodb
>sudo docker run -it --link marathon-mongodb:mongo --rm mongo:3.4 sh -c 'exec mongo "$MONGO_PORT_27017_TCP_ADDR:$MONGO_PORT_27017_TCP_PORT/marathon"'


### 4. zookeeper
run zookeeper service
>sudo docker pull zookeeper:3.4

>sudo docker run -p 2181:2181 --name marathon-zookeeper --restart always -d zookeeper:3.4

connect to zookeeper 
>sudo docker run -it --rm --link marathon-zookeeper:zookeeper zookeeper:3.4 zkCli.sh -server zookeeper