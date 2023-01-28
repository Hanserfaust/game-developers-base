import unittest

import multiqueue
import queue

from messages import GameMessage, MouseMoveGameMessage



class TestGameMessages(unittest.TestCase):

    def test_message_serialization_using_args(self):

        #
        # Given
        #
        orig_message = MouseMoveGameMessage(10, 20)
        print("Original message: " + str(orig_message))
        client_id = 1234
        bin_packet = orig_message.serialize(client_id)
        # print(bin_packet)

        #
        # When
        #
        # > using the abstract ClientEvent, will create a MouseMoveClientEvent
        mouse_message = GameMessage.factory(bin_packet)

        #
        # Then
        #
        # > deserialized message should match the original message
        self.assertEqual(orig_message.name, mouse_message.name)
        self.assertEqual(orig_message.data, mouse_message.data)
        self.assertEqual(client_id, mouse_message.client_id)
        self.assertEqual(orig_message.as_dict(), mouse_message.as_dict())

        print("Deserialized message: " + str(mouse_message))
        print(orig_message.as_dict())

    def test_message_serialization_using_kwargs(self):

        #
        # Given
        #
        orig_message = MouseMoveGameMessage(x=10, y=20)
        print("Original message: " + str(orig_message))
        client_id = 1234
        bin_packet = orig_message.serialize(client_id)
        # print(bin_packet)

        #
        # When
        #
        # > using the abstract ClientEvent, will create a MouseMoveClientEvent
        mouse_message = GameMessage.factory(bin_packet)

        #
        # Then
        #
        # > deserialized message should match the original message
        self.assertEqual(orig_message.name, mouse_message.name)
        self.assertEqual(orig_message.data, mouse_message.data)
        self.assertEqual(client_id, mouse_message.client_id)
        self.assertEqual(orig_message.as_dict(), mouse_message.as_dict())

        print("Deserialized message: " + str(mouse_message))
        print(orig_message.as_dict())


class TestMultiQueue(unittest.TestCase):

    def test_multiqueue(self):

        q1 = queue.Queue()
        q2 = queue.Queue()

        multi_queue = multiqueue.MultiQueue([q1, q2])

        q1.put(1)
        q2.put(2)
        q1.put(3)
        q2.put(4)

        res = 0
        while not res == 4:
            while not multi_queue.empty():
                res = multi_queue.get()[1]
                print("returning result =", res)
