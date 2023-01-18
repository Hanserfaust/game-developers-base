import msgpack

from client import Client


class ClientEvent(object):
    """
        Abstract base-class representing events on a game client.
        Data is any dict. The contents are considered event specific.
    """

    # Must be set if sending to server
    client_id = Client.id

    # Below must be defined by subclasses
    name = None
    data_keys = []

    def __init__(self, data):
        # Data is a list representing the keys in data_keys
        assert(len(data), len(self.data_keys))
        self.data = data

    def __str__(self):
        return "%s: %s" % (self.name, self.data)

    def as_dict(self):
        return dict(zip(self.data_keys, self.data))

    def build_packet(self):
        return [self.name, self.client_id] + self.data

    def serialize(self):
        # First build a packet
        packet = self.build_packet()
        return msgpack.packb(packet, use_bin_type=True)

    @staticmethod
    def factory(bin_packet):
        """
        Given a binary packet, returns a ClientEvent-object of correct concrete class.

        :param bin_packet:
        :return:
        """
        packet = msgpack.unpackb(bin_packet, raw=False)
        name, client_id, data = packet[0], packet[1], packet[2:]

        # Instantiate new event object
        e = events_registry[name](data)
        e.client_id = client_id

        return e


class MouseMoveClientEvent(ClientEvent):
    # Data: x, y
    # Notes: Clicks are not MouseMove events.
    name = 'MOUSE_MOVE'
    data_keys = ['x', 'y']


events_registry = {
    MouseMoveClientEvent.name: MouseMoveClientEvent
}
