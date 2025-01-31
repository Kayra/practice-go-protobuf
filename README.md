# practice-go-protobuf

This is the code written while roughly following the [Protocol Buffers (ProtoBuf) with GoLang tutorial](https://medium.com/trendyol-tech/protocol-buffers-protobuf-with-golang-41d0d332745d).

Technologies used:

- go: 1.23.5 darwin/arm64
- protobuf: 29.3
- protoc-gen-go: 1.36.4

## Set up and run

Get the protobuf package:

```bash
cd main; go get github.com/golang/protobuf; go mod tidy
```

Run the code to see the output:

```bash
cd main; go run .

Original person:  name:"Rie" age:31 email:"rie@example.com" email:"rie@gmail.com"
Serialized person:  [10 3 82 105 101 16 31 26 15 114 105 101 64 101 120 97 109 112 108 101 46 99 111 109 26 13 114 105 101 64 103 109 97 105 108 46 99 111 109]
Buff person:  name:"Rie" age:31 email:"rie@example.com" email:"rie@gmail.com"
```

## Useful commands

Compile protocol buffers:

```bash
cd person; protoc --go_out=. person.proto
```