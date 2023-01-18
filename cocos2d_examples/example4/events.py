import msgpack


class ClientEvent(object):
    """
        Abstract base-class representing events on a game client.
        Data is any dict. The contents are considered event specific.
    """

    # Below must be defined by subclasses
    name = None
    data_keys = []

    # Set on server side upon deserialization
    client_id = None

    def __init__(self, *data):
        # Data is a tuple representing the keys in data_keys
        assert len(data) == len(self.data_keys)
        self.data = data

    def __str__(self):
        if self.client_id:
            return "Client %s: %s data=(%s)" % (self.client_id, self.name, self.data)
        else:
            return "%s: %s" % (self.name, self.data)

    def as_dict(self):
        return dict(zip(self.data_keys, self.data))

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

        :param bin_packet:
        :return:
        """
        packet = msgpack.unpackb(bin_packet, raw=False)
        name, client_id, data = packet[0], packet[1], packet[2]

        # Instantiate new event object
        e = events_registry[name](*tuple(data))
        e.client_id = client_id

        return e


class HelloServerEvent(ClientEvent):
    name = 'HELLO_SERVER'
    data_keys = ['message']


class MouseMoveClientEvent(ClientEvent):
    # Data: x, y
    # Notes: Clicks are not MouseMove events.
    name = 'MOUSE_MOVE'
    data_keys = ['x', 'y']


events_registry = {
    HelloServerEvent.name: HelloServerEvent,
    MouseMoveClientEvent.name: MouseMoveClientEvent
}
