# go-apache-kafka

kafka

ctrl + l

kafka-topics

kafka-topics --create --topic=teste --bootstrap-server=localhost:9092 --partitions=3

kafka-topics --bootstrap-server=localhost:9092 --topic=teste --describe

kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste 

docker exec -it fc2-kafka_kafka_1 bash

kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste 

kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning

kafka-consumer-groups --bootstrap-server=localhost:9092 --group=y --describe

cat /etc/hosts

10.0.0.104      host.docker.internal
127.0.0.1       kubernetes.docker.internal

127.0.0.1       kubernetes.docker.internal docker.host.internal

go mod init github.com/jdoliveirasa/fc2-gokafka

kafka-topics --create --bootstrap-server=localhost:9092 --topic=teste --partitions=3

kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste

kafka-consumer-groups --bootstrap-server=localhost:9092 --group=goapp-group --describe

kafka-connect

mysql -uroot -proot fullcycle

create table categories (
id int auto_increment primary key,
name varchar(255)
);

show tables;

desc categories;
 
insert into categories (name) values ('Eletronicos');
insert into categories (name) values ('Moveis');
insert into categories (name) values ('Veiculos');
insert into categories (name) values ('Celulares');

insert into categories (name) values ('Imoveis');
insert into categories (name) values ('Cozinha');

insert into categories (name) values ('Carrinho');

insert into categories (name) values ('Teste01');
insert into categories (name) values ('Teste02');
insert into categories (name) values ('Teste03');

SELECT * FROM categories;

