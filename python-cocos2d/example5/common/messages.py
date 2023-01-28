import msgpack


class GameMessage(object):
    """
        Abstract base-class representing messages sent between client and server.

        The data carried is message specific is given to the constructor as args or
        as kwargs and must match length of the given data_keys (no optional data).

    """

    # Below must be defined by subclasses
    name = None
    data_keys = []

    # Set on server side upon deserialization
    client_id = None

    def __init__(self, *args, **kwargs):
        if args:
            assert len(args) == len(self.data_keys)
            self.data = args
        if kwargs:
            assert len(kwargs.keys()) == len(self.data_keys)
            data = list()
            for key in self.data_keys:
                data.append(kwargs[key])
            self.data = tuple(data)

    def __str__(self):
        if self.client_id:
            return "Client %s: %s data=%s" % (self.client_id, self.name, self.data)
        else:
            return "%s: %s" % (self.name, self.data)

    def as_dict(self):
        return dict(zip(self.data_keys, self.data))

    def d(self):
        return self.as_dict()

    def build_packet(self, client_id):
        return self.name, client_id, self.data

    def serialize(self, client_id):
        # Builds a packet to send over the socket
        packet = self.build_packet(client_id)
        return msgpack.packb(packet, use_bin_type=True)

    @staticmethod
    def factory(bin_packet):
        """
        Given a binary packet, returns a ClientEvent-object of correct concrete class.

        Used after receiving a bin_packet over the wire. Like so:

        received_game_message = GameMessage.factory(bin_packet)

        :param bin_packet:
        :return:
        """
        packet = msgpack.unpackb(bin_packet, raw=False)
        name, client_id, data = packet[0], packet[1], packet[2]

        # Instantiate new message object
        e = message_registry[name](*tuple(data))
        e.client_id = client_id

        return e


# TODO: Separate to-server and to-client messages?

class HelloServerGameMessage(GameMessage):
    name = 'HELLO_SERVER'
    data_keys = ['message']


class GlobalMessage(GameMessage):
    name = 'GLOBAL_MESSAGE'
    data_keys = ['text']


class MouseMoveGameMessage(GameMessage):
    name = 'MOUSE_MOVE'
    data_keys = ['x', 'y']


message_registry = {
    HelloServerGameMessage.name: HelloServerGameMessage,
    MouseMoveGameMessage.name: MouseMoveGameMessage
}
