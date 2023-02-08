import random
import socket


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

    def send(self, packet):
        try:
            # is a high-level Python-only method that sends the entire buffer you pass or throws
            # an exception. It does that by calling socket.send until everything has been sent
            # or an error occurs.
            self.sock.sendall(packet)
        except socket.error as e:
            print(e)

    def receive(self):
        rest_size = self.sock.recv(1)
        rest_size = int(rest_size[0])
        type_and_message = self.sock.recv(rest_size)
        message_type = type_and_message[0]
        message = type_and_message[1:]

        print("Received packet of type: %s, len: %s" % (message_type, len(message)))

        return message_type, message
