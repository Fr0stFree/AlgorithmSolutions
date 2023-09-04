import pika

CONNECTION_STRING = "amqps://student:XYR4yqc.cxh4zug6vje@rabbitmq-exam.rmq3.cloudamqp.com/mxifnklj"
QUEUE_NAME = "exam"
MESSAGE = "Hi CloudAMQP, this was fun!"
ROUTING_KEY = "efa5b7f6-353f-4246-a32f-e7b6161449e4"
EXCHANGE_NAME = "exchange.efa5b7f6-353f-4246-a32f-e7b6161449e4"


if __name__ == "__main__":
    params = pika.URLParameters(CONNECTION_STRING)
    connection = pika.BlockingConnection(params)
    channel = connection.channel()

    channel.exchange_declare(
        exchange=EXCHANGE_NAME,
        exchange_type="direct",
        passive=False,
    )

    channel.queue_declare(queue=QUEUE_NAME, durable=True)
    channel.queue_bind(queue=QUEUE_NAME, exchange=EXCHANGE_NAME, routing_key=ROUTING_KEY)

    channel.basic_publish(
        exchange=EXCHANGE_NAME,
        routing_key=ROUTING_KEY,
        body=MESSAGE,
        properties=pika.BasicProperties(delivery_mode=2),
    )

    print("Message sent")
    channel.close()
    connection.close()

