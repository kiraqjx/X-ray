# X-ray

## Introduce
Network traffic proxy tool

## Use

#### Start server
```shell
cd server
go run ./server.go
```

#### Start the test client
```shell
cd client
go run ./client.go
```
The server terminal then displays the proxy port  
like: Proxy from 127.0.0.1:53804 to 127.0.0.1:8082

The client terminal also displays the proxy port  
like: 8082 -> 127.0.0.1:53804

#### Test proxy
```shell
telnet <server ip> 53804
```
The client terminal then displays "hello world"
