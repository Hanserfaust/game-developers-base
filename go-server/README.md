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
This is fairly extensive, read on in its own [README.md](./example1/README.md).