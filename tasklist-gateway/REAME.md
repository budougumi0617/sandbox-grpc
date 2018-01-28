# TaskList with Gateway sever

Simple gRPC server wtih REST proxy server.

See below

https://github.com/grpc-ecosystem/grpc-gateway


# Example
Start gRPC server.

```bash
$ go run ./server/main.go
```

Start gateway server.

```bash
$ go run ./gateway/main.go
```

Try REST API from other terminal.

```bash
$ curl -D - -X GET http://localhost:8080/v1/tasklist-gateway/task
HTTP/1.1 200 OK
Content-Type: application/json
Grpc-Metadata-Content-Type: application/grpc
Date: Sun, 28 Jan 2018 04:03:14 GMT
Transfer-Encoding: chunked

{"result":{"id":20,"title":"List Tasks 1","detail":"Implement interface"}}
{"result":{"id":21,"title":"List Tasks 2","detail":"Implement interface"}}
{"result":{"id":22,"title":"List Tasks 3","detail":"Implement interface"}}

$ curl -D - -X GET http://localhost:8080/v1/tasklist-gateway/task/11
HTTP/1.1 500 Internal Server Error
Content-Type: application/json
Trailer: Grpc-Trailer-Content-Type
Date: Sun, 28 Jan 2018 04:03:29 GMT
Transfer-Encoding: chunked

{"error":"Not find Task","code":2}Grpc-Trailer-Content-Type: application/grpc

$ curl -D - -X GET http://localhost:8080/v1/tasklist-gateway/task/10
HTTP/1.1 200 OK
Content-Type: application/json
Grpc-Metadata-Content-Type: application/grpc
Date: Sun, 28 Jan 2018 04:03:33 GMT
Content-Length: 72

{"id":10,"title":"Implement gRPC server","detail":"Implement interface"}%
```

gRPC server prints below logs.

```bash
$ go run ./server/main.go
2018/01/28 13:00:48 gRPC Server started: localhost:50051
2018/01/28 13:03:14 ListTasks in gPRC server
2018/01/28 13:03:29 GetTask in gPRC server
2018/01/28 13:03:33 GetTask in gPRC server
```

