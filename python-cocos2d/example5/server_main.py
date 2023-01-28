import socketserver
import threading

from common.messages import GameMessage, HelloServerGameMessage, MouseMoveGameMessage


class GameMessageHandler(socketserver.StreamRequestHandler):
    """
    The request handler class for our server.

    It is instantiated once per connection to the server, and must
    override the handle() method to implement communication to the
    client.
    """

    def send_message(self, game_message):
        #
        # TODO: Lock around connection?
        #
        self.connection.send(game_message)

    def handle(self):
        """
            Handler for a client
        :return:
        """
        print("Client Connected!")

        #
        # TODO: Need to register this GameMessageHandler centrally based on client_id
        #

        messages = 0
        #
        # TODO: While "connected" etc.
        #
        while True:
            # self.request is the last TCP socket connected to the client
            # recv(LARGER THAN SINGLE PACKET), so 1024 works.
            #
            #
            try:
                data = self.request.recv(1024)

                if data:
                    messages += 1
                    gm = GameMessage.factory(bin_packet=data)
                    print("Message %s (len=%s): %s" % (messages, len(data), gm))

                    #
                    # TODO: Publish a "GotClientMessage()" event here.
                    #

                    match gm.name:

                        case HelloServerGameMessage.name:
                            pass

                        case MouseMoveGameMessage.name:
                            #
                            # For this example, just echo back the coordinates a GameMessage
                            #   this is essentially the starting point of the next example.
                            #
                            dd = gm.as_dict()
                            server_gm = MouseMoveGameMessage(dd['x'], dd['y'])
                            ser = server_gm.serialize('SERVER')
                            self.request.sendall(ser)
                            print("Sending from server: %s" % server_gm)

            except ConnectionResetError as e:
                print("ConnectionResetError: %s" % e)


def main():
    host, port = "localhost", 7777

    print("Server Started, listening for connections on port %s" % port)

    #
    # Start the Main thread, this will serve events and the world for now
    # this may be split further in later examples
    #
    threading.Thread()

    #
    # Create the server that will spawn one thread for each client.
    #
    with socketserver.ThreadingTCPServer((host, port), GameMessageHandler) as server:
        # Activate the server; this will keep running until you
        # interrupt the program with Ctrl-C
        server.serve_forever()


if __name__ == "__main__":
    main()
