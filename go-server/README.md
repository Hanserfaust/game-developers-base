# Go server examples
In this section we focus on server concurrency using Go (or Golang).

In some cases we re-use and even build on the Python cocos2d-based clients. In that case
we assume you have worked through the python-cocos2d examples and have the "pyg" virtual
environment available.

## Notes and discoveries
During development of the examples, I discovered possible improvements, note here:

- More modern/better binary serialization? https://capnproto.org/

## Go environment
Recommend using Visual Studio Code for Go development.

For mac, just install the golang using

    brew install golang

## Go Resources
I'll just drop links to useful resources related to Go development in general here:

Standard library: https://pkg.go.dev/std

Curated list of popular Go libraries: https://github.com/avelino/awesome-go

Add this to your shell rc of choice (like ~/.bashrc )

        export GOPATH=$HOME/go
        export PATH=$PATH:$GOPATH/bin

## Go Module init
Note, each example is set up as its own go module to get its dependecy tracking right etc., like so:

    go mod init gdb/golangserver/exampleX

And after adding new dependencies:

    go tidy

## Go examples

### Example 0
Simple Go hello-world example to ensure your GO-environment is working.

### Example 1
Cocos 2D client and a Go-based server echoing a message back to the client(s).

It is similar to Example 4) of the python-cocos2d example, but this one instead tests Proto Buffers
as Wire format.

https://medium.com/@hugovs/the-need-for-speed-experimenting-with-message-serialization-93d7562b16e4

One difference between MessagePack (MP) and ProtoBuffers (PB) is that PB can generate "model stubs"
for the supported languages, based on a **common schema**. Sometimes this is NOT what you want (and
in that case I would recommend MessagePack), but in our case, this is exactly what want. 

Why?  Because we expect or game message data structures to evolve frequently and we want to be able to
update both the server and the client quickly without having to manually update the message source
code on each side for each change we make.

We will to define our message protocol using .proto files, and from that we will generate stubs in
the language used by the server (Go) and for the client(s), Python and Possibly javascript.

Read more on Protocol Buffers here:

https://protobuf.dev/

#### Message packet structure
While the protocol buffers specify the messages and allows for effective serial- and deserialization
between the client and server, we still need to consider how to send the messages effectively.

A common strategy is to delimit each packet with the packet size, and since we are sending packets
of different types, we also add a single byte to tell the de-serializing side what type to create:

**[1 byte packet length] + [1 byte packet type] + [N bytes packet data]**

A receiver read strategy will then be:

1. Read 1 byte determining packet size.
2. Read 1 byte determining packet type.
3. Read [packet size] bytes into a buffer.
4. Instantiate message of [packet type].

The above structure is defined in the messages.proto as the type **GameMessage** which is the generic
message-type encapsulating the type and the actual packet. It does not include the packet size

#### Go support
Download the sources (I used the "ALL" version) and build them using make && make install

And we also need these:

        go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

The example contains the file messages.proto, compile it to GO using one of the below:

        protoc --go_out=. *.proto

#### Python support
Googles Python generator is quite crappy, this one is much better.

https://github.com/danielgtaylor/python-betterproto 

So install that (in the correct virtual env, still "pyg") and have a brief look at that page for 
some basic instruction on usage.

        pip install "betterproto[compiler]"
        pip install betterproto    

To generate the Python message classes we do:

        protoc --python_betterproto_out=. messages.proto

### Updating message model
The neat thing now of course is, that each time we want to update a message, or add a new one we
just generate the message stubs for the client and server with a single command

Essentially combining the above calls to protoc like this, with the addition of generating
the respective stubs in the correct sub-directory:

        protoc --go_out=./server --python_betterproto_out=./client messages.proto

And why not add that command to a simple shell script [gen_stubs.py](./example1/gen_stubs.sh).

        hanseklund@Hans-MacBook-Pro example1 % ./gen_stubs.sh
        Writing __init__.py
        Writing messages.py

### Python client message serialization
Just a brief pointer on how ridiculously great this is, in the [mousy_game_client.py](./example1/mousy_game_client.py):

        mouse_move_message = messages.MouseMove(x, y, left_click=False, right_click=False)        
        MessageTCPClient.send(bytes(mouse_move_message))

In those two lines, from a shared model (with the Go-server) we instantiate a Python message
object and just serialize and send it over the wire.

### Go server message de-serialization
And, on the server-side we just do this

### Example 3
Same as example 1, but a script spawning any number of headless clients just moving a virtual cursor
sending the same types of packages to the server and echoing all coordinates out to all clients.

This is to experiment with scalability.