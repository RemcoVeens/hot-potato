trow:
    AMQP_URL=amqp://potato:secret@localhost:5672/ go run ./cmd/insert_potato
play:
    docker compose up --build --no-attach rabbitmq --scale player=3
