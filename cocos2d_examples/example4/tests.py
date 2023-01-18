import unittest

from events import ClientEvent, MouseMoveClientEvent


class TestEvents(unittest.TestCase):

    def test_event_serialization(self):

        #
        # Given
        #
        orig_event = MouseMoveClientEvent(10, 20)
        print("Original event: " + str(orig_event))
        client_id = 1234
        bin_packet = orig_event.serialize(client_id)
        # print(bin_packet)

        #
        # When
        #
        # > using the abstract ClientEvent, will create a MouseMoveClientEvent
        mouse_event = ClientEvent.factory(bin_packet)

        #
        # Then
        #
        # > deserialized event should match the original event
        self.assertEqual(orig_event.name, mouse_event.name)
        self.assertEqual(orig_event.data, mouse_event.data)
        self.assertEqual(client_id, mouse_event.client_id)
        print("Deserialized event: " + str(mouse_event))
