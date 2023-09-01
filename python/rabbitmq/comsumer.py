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

    channel.basic_consume(
        queue=QUEUE_NAME,
        on_message_callback=lambda ch, method, properties, body: print(body),
        auto_ack=True,
    )

    channel.start_consuming()
    connection.close()
