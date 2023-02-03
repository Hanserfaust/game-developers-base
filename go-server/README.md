# Go server examples

## Go environment
Recoomend using Visual Studio Code for Go development.

For mac, just install the golang using

    brew install golang


## Go Resources

Standard librar: https://pkg.go.dev/std

Curated list of popular Go libraries: https://github.com/avelino/awesome-go

Add this to your shell rc of choice (like ~/.bashrc )

        export GOPATH=$HOME/go
        export PATH=$PATH:$GOPATH/bin

## Go server examples
Note, each example is set up as its own go module to get its dependecy tracking right etc., like so:

    go mod init gdb/golangserver/exampleX

And after adding new dependencies:

    go tidy


### Example 0
Simple Go hello-world example

### Example 1
Cocos 2D client and a Go-based server echoing a message back to the client(s).

It is similar to Example 4) of the python-cocos2d example, but this one instead tests Proto Buffers
as Wire format.

https://medium.com/@hugovs/the-need-for-speed-experimenting-with-message-serialization-93d7562b16e4

The difference between MessagePack (MP) and ProtoBuffers (PB) is that PB can generate "model stubs"
for the supported languages, based on a common schema. Sometimes this is NOT what you want, but in
our case, this is exactly what want. The message model must be pre-defined to ensure the server
and client is in sync for the over-the-wire (OTW) communication.

### The .proto format
We need to define our OTW protocol using .proto files, and from them we can generate go stubs from
the common message format defined in messages.proto.

Read more on Protocol Buffers here:

https://protobuf.dev/

Download the sources (I used the "ALL" version) and build them using make && make install

And we also need these:

        go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

The example contains the file messages.proto, compile it to GO using one of the below:

        protoc --go-grpc_out=. *.proto
        protoc --go_out=. *.proto

and for python we do:

        protoc --python_out=. *.proto

Now, this can be combined into this to be run after each change to the. proto-file.

        protoc --go_out=. --python_out=. *.proto

### Example 3
Same as example 1, but a script spawning any number of headless clients just moving a virtual cursor
sending the same types of packages to the server and echoing all coordinates out to all clients.

This is to experiment with scalability.