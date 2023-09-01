import pika

CONNECTION_STRING = "amqp://guest:guest@localhost:5672/"
QUEUE_NAME = "test_queue"
EXCHANGE_NAME = "test_exchange"
ROUTING_KEY = "test_routing_key"


if __name__ == "__main__":
    params = pika.URLParameters(CONNECTION_STRING)
    connection = pika.BlockingConnection(params)
    channel = connection.channel()

    channel.exchange_declare(exchange=EXCHANGE_NAME, exchange_type="direct")
    channel.queue_declare(queue=QUEUE_NAME)
    channel.queue_bind(
        exchange=EXCHANGE_NAME,
        queue=QUEUE_NAME,
        routing_key=ROUTING_KEY,
    )

    channel.basic_publish(
        exchange=EXCHANGE_NAME,
        routing_key=ROUTING_KEY,
        body=b"Hello World!",
    )

    channel.close()
    connection.close()
