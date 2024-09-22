import os
import pika


class Broker:
    def __init__(self):
        self.connection = self.__new_connection()

    def __new_connection(self):
        return pika.BlockingConnection(
            pika.ConnectionParameters(
                host=os.getenv("RMQ_HOST"),
                port=os.getenv("RMQ_PORT"),
            )
        )

    def update_connection(self):
        if self.connection.is_closed:
            self.connection = self.__new_connection()
