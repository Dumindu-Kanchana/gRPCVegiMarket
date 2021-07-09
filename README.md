# gRPCVegiMarket
This is a simple client/server application which uses RPC to perform operations

These codes are developed using following versions.

go version - 1.16.5 <br/>
protoc version - 3.17.3

#### useful docs

https://grpc.io/docs/languages/go/quickstart/ <br/>
https://developers.google.com/protocol-buffers/docs/gotutorial <br/>
https://github.com/protocolbuffers/protobuf/blob/master/examples <br/>

## _Instructions to run the clients & server._

run server from the home (vegimarket.com)

```sh
  go run marketServer/marketServer.go
```
run addVegetableClient

```sh
go run marketClient/addVegetableClient.go potato 4.54 3.4
```

run updateVegetableClient

```sh
go run marketClient/updateVegetableClient.go carrot 4.54 -1
```

if -1 was specified, the values will not get updated.