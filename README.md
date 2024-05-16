# go-websocket

`go-websocket` is a proof of concept implementation of a Go websocket microservice running behind an nginx proxy

`docker-compose.yml` sets up the go websocket, a php server, and nginx to proxy between them

`GET /ws/sub/:topic` allows a client to subscribe to a topic

`POST /ws/pub/:topic` allows a service to publish data to a topic

When we publish to a topic, the data is broadcasted to the corresponding subscribers

## Quickstart

In your terminal run

```sh
git clone github.com/sebar-io/go-websocket
cd /go-websocket
docker compose up
```

Open [localhost/?topic=a](http://localhost/?topic=a) in a browser and open the page's console

Use your terminal to publish data to the topic

```sh
curl --location 'localhost/ws/pub/a' \
--header 'Content-Type: application/json' \
--data '{
    "foo": "bar"
}'
```

Observe your browser's console log the data