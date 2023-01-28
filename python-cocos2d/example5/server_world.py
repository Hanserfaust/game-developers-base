import time, random, queue

from common.messages import GlobalMessage
from common import multiqueue


#
# Various parts of the game will post events to the event queue
#
to_clients_messages = queue.Queue()

# Maps client_id to a GameMessageHandler
# This should probably be abstracted further, but this for
# having our very simple "world" send some messages to the clients.
player_handlers = dict()


# Started as thread
def client_messenger():

    while True:
        client_message = to_clients_messages.get()




# Just for testing this simple world
def random_event_producer():

    while True:
        time.sleep(random.randint(0, 10))

        to_clients_messages.put(GlobalMessage("Hello from world!"))
