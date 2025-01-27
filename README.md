# practice-go-protobuf

This is the code written while roughly following the [protobuf getting started tutorial](https://protobuf.dev/getting-started/gotutorial/).

Technologies used:

- go: 1.23.5 darwin/arm64
- protobuf: 29.3
- protoc-gen-go: 1.36.4

## Set up

Install go protobuf plugin:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## Useful commands

Compile protocol buffers:

```bash
protoc --proto_path=./src --go_out=./src/go ./src/addressbook.proto
```