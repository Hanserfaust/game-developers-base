import random
import socket

# TODO: Find better solution to this, works for now
from packet import build_packet


class MessageTCPClient(object):
    """
        Singleton representing/abstracting interactions with the server.

        Will _NOT_ connect upon creation.

        Does NOT assume/know of the specific message format. It just sends and receives
        a string of bytes, optionally prefixing the send operation with the size of the
        given buffer.

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
        if not MessageTCPClient._instance:
            MessageTCPClient._instance = MessageTCPClient(**kwargs)
        return MessageTCPClient._instance

    def connect(self):
        self.sock.connect(self.addr)

    def _send(self, packet):
        try:
            # is a high-level Python-only method that sends the entire buffer you pass or throws
            # an exception. It does that by calling socket.send until everything has been sent
            # or an error occurs.
            self.sock.sendall(packet)
        except socket.error as e:
            print(e)

    @staticmethod
    def send(game_message):
        """
        Works for now, but not sure if this module/function should call into the packet module.

        :param game_message:
        :return:
        """

        packet = build_packet(game_message)

        MessageTCPClient.get_instance()._send(packet)
