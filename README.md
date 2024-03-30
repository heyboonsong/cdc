## create trable users

```
CREATE TABLE users (
    id        serial primary key,
    first_name     VARCHAR(40) not null,
    last_name      VARCHAR(40) not null,
    age            INTEGER not null,
    address        VARCHAR(255) not null
)

```

## create a connector that captures changes from a PostgreSQL database and publishes them to Kafka topics.

```
curl --location 'http://localhost:8083/connectors' \
 --header 'Accept: application/json' \
 --header 'Content-Type: application/json' \
 --data '{
    "name": "debezium-connector",
    "config": {
        "connector.class": "io.debezium.connector.postgresql.PostgresConnector",
        "database.hostname": "192.168.1.86",
        "database.port": "5443",
        "database.user": "postgres",
        "database.password": "P123!",
        "database.dbname": "data-liberation",
        "database.server.id": "184054",
        "table.include.list": "public.users",
        "topic.prefix": "debezium-topic"
    }
}'
```

## run existing-system

```
cd /existing-system
go run main.go
```

## run worker

```
cd /worker
yarn install
yarn dev
```

## create users

```
curl --request POST \
  --url http://localhost:8000/users \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/8.6.1' \
  --data '{
	"first_name": "Test",
	"last_name": "Test",
	"age": 18,
	"address": "Test"
}'
```
