#!/bin/bash

start_or_run () {
    docker inspect hot_potato_rabbitmq > /dev/null 2>&1

    if [ $? -eq 0 ]; then
        echo "Starting Hot potato RabbitMQ container..."
        docker start hot_potato_rabbitmq
    else
        echo "Hot potato RabbitMQ container not found, creating a new one..."
        docker run -d --name hot_potato_rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4.4-management
    fi
}

case "$1" in
    start)
        start_or_run
        ;;
    stop)
        echo "Stopping Hot potato RabbitMQ container..."
        docker stop hot_potato_rabbitmq
        ;;
    logs)
        echo "Fetching logs for Hot potato RabbitMQ container..."
        docker logs -f hot_potato_rabbitmq
        ;;
    *)
        echo "Usage: $0 {start|stop|logs}"
        exit 1
esac
