import socketserver

from events import ClientEvent


class GameEventHandler(socketserver.BaseRequestHandler):
    """
    The request handler class for our server.

    It is instantiated once per connection to the server, and must
    override the handle() method to implement communication to the
    client.
    """

    def handle(self):
        """

            Handler for a client

        :return:
        """
        print("Client Connected!")

        while True:
            # self.request is the TCP socket connected to the client
            # recv(LARGER THAN SINGLE PACKET), so 1024 works.
            data = self.request.recv(1024)

            if data:
                client_event = ClientEvent.factory(bin_packet=data)
                print("EVENT (len=%s): %s" % (len(data), client_event))


def main():
    host, port = "localhost", 7777

    print("Server Started, listening for connections on port %s" % port)

    # Create the server, binding to localhost on port 9999
    with socketserver.TCPServer((host, port), GameEventHandler) as server:
        # Activate the server; this will keep running until you
        # interrupt the program with Ctrl-C
        server.serve_forever()


if __name__ == "__main__":
    main()
