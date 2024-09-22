import sys

if __name__ == "__main__":
    if sys.argv[1] == "server":
        from src.app.server.server import server
        server()

    if sys.argv[1] == "consumer":
        from src.app.consumer.consumer import consumer
        consumer()
