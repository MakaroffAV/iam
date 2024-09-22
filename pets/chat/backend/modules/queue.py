from src.config.broker import Broker as ConfigBroker


class Queue:
    def __init__(self, broker: ConfigBroker) -> None:
        self.__broker = broker

    def send(self, queue_name: str, msg_body: str) -> None:
        self.__broker.update_connection()
        conn = self.__broker.connection

        chan = conn.channel()

        chan.queue_declare(queue=queue_name)
        chan.basic_publish(
            exchange="", routing_key=queue_name, body=msg_body,
        )
