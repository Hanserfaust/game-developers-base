import socketserver

from messages import GameMessage, HelloServerGameMessage, MouseMoveGameMessage


class GameMessageHandler(socketserver.BaseRequestHandler):
    """
    The request handler class for our server.

    It is instantiated once per connection to the server, and must
    override the handle() method to implement communication to the
    client.
    """

    def handle(self):
        """
            Handler for a client.

            NOTE: This server is NOT threaded, so the while-loop below will block the thread
            forver. The server will NOT be able to server multiple clients.

            This will be improved on in the next example.

        :return:
        """
        print("Client Connected!")

        messages = 0
        #
        # TODO: While "connected" etc.
        #
        while True:
            #
            # self.request is the last TCP socket connected to the client
            # recv(LARGER THAN SINGLE PACKET), so 1024 works.
            #
            try:
                data = self.request.recv(1024)

                if data:
                    messages += 1
                    gm = GameMessage.factory(bin_packet=data)
                    print("EVENT %s (len=%s): %s" % (messages, len(data), gm))

                    match gm.name:

                        case HelloServerGameMessage.name:
                            pass

                        case MouseMoveGameMessage.name:
                            pass

                    #
                    # Further work:
                    #
                    # Ok, got message from client, could check that game ID is correct
                    # for this socket etc.
                    #
                    # This is a good place to just announce that we got data (Observer pattern)
                    # and just move one and let the Observer decide how this event affected
                    # the server state.
                    #
                    # This is thing we will focus on in later examples.
                    #

            except ConnectionResetError as e:
                print("ConnectionResetError: %s" % e)


def main():
    host, port = "localhost", 7777

    print("Server Started, listening for connections on port %s" % port)

    # Create the server, binding to localhost on port 9999
    with socketserver.TCPServer((host, port), GameMessageHandler) as server:
        # Activate the server; this will keep running until you
        # interrupt the program with Ctrl-C
        server.serve_forever()


if __name__ == "__main__":
    main()
