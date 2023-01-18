import random


class Client(object):
    """
        Singleton representing/abstracting interactions with the server.
    """

    _instance = None

    # TODO: Supplied from server upon connect
    id = random.randint(0, 2**31)

    @staticmethod
    def get():
        if not Client._instance:
            Client._instance = Client()
        return Client._instance

    def connect(self):
        pass

    def send_event(self, client_event, asynch=True):
        pass
