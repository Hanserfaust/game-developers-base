import random
import threading
import socket


class GameClient(threading.Thread):
    """
        Singleton representing/abstracting interactions with the server.

        Will _NOT_ connect upon creation.
    """

    _instance = None

    # TODO: Supplied from server upon connect
    id = random.randint(0, 2**31)

    def __init__(self, hostname="127.0.0.1", port=7777, debug=False):
        self.server = hostname
        self.port = port
        self.addr = (self.server, self.port)
        self.debug = debug

        # SOCK_STREAM means TCP socket
        self.sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    @staticmethod
    def get_instance(**kwargs):
        if not GameClient._instance:
            GameClient._instance = GameClient(**kwargs)
        return GameClient._instance

    def connect(self):
        self.sock.connect(self.addr)

    def _send(self, data):
        try:
            # is a high-level Python-only method that sends the entire buffer you pass or throws
            # an exception. It does that by calling socket.send until everything has been sent
            # or an error occurs.
            self.sock.sendall(data)
        except socket.error as e:
            print(e)

    @staticmethod
    def send_message(game_message):
        # Quick access method for sending data to server
        gc = GameClient.get_instance()
        if gc.debug:
            print("Sending %s" % game_message)
        bin_packet = game_message.serialize(GameClient.id)
        gc._send(bin_packet)
