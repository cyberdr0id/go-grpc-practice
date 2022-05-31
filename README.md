# Project presents implementation of different gRPC communications.

### Makefile uses for *Windows*  OS

First of all you should build project.  
> Make sure that [GNU Make](https://www.gnu.org/software/make/) is installed, otherwise follow the commands in the appropriate group
```
make <communication_type>
```

This command will genereate Go code from `.proto` file and `bin` folder will be created with folder `<communication_type>` inside, which will contain two binary files `client.exe` and `server.exe`.

# unary
### Example of *unary* API with server and client implementation. In this case, server will return sum of two interegs sent by client.
---
```
protoc -Iunary/proto --go_opt=module=github.com/cyberdr0id/go-grpc-practice --go_out=. --go-grpc_opt=module=github.com/cyberdr0id/go-grpc-practice --go-grpc_out=. unary/proto/*.proto

go build -o bin/unary/client.exe ./unary/client
go build -o bin/unary/server.exe ./unary/server
```

To start application you must run `server.exe` file with port number in command line arguments:
```
./bin/<communication_type>/server <port_number>
```
And then run `client.exe` with the same port number and two additional integer parameters.
```
./bin/<communication_type>/client <port_number> <int1> <int2>
```  

# sstream
### Example of *server streaming* API with server and client implementation. Here, server decomposes a number into prime factors.
---
```
protoc -Isstream/proto --go_opt=module=github.com/cyberdr0id/go-grpc-practice --go_out=. --go-grpc_opt=module=github.com/cyberdr0id/go-grpc-practice --go-grpc_out=. sstream/proto/*.proto

go build -o bin/sstream/client.exe ./sstream/client
go build -o bin/sstream/server.exe ./sstream/server
```
To start application you must run `server.exe` file with port number in command line arguments:
```
./bin/<communication_type>/server <port_number>
```
And then run `client.exe` with the same port number and one additional integer parameter
```
./bin/<communication_type>/client <port_number> <int1>
```  

# cstream
### Example of *client streaming* API with server and client implementation. In this example, the server will calculate average of the numbers, that will be sent by the client.
---
```
protoc -Icstream/proto --go_opt=module=github.com/cyberdr0id/go-grpc-practice --go_out=. --go-grpc_opt=module=github.com/cyberdr0id/go-grpc-practice --go-grpc_out=. cstream/proto/*.proto

go build -o bin/cstream/client.exe ./cstream/client
go build -o bin/cstream/server.exe ./cstream/server
```
To start application you must run `server.exe` file with port number in command line arguments:
```
./bin/<communication_type>/server <port_number>
```
And then run `client.exe` with the same port number and additional integer parameters
```
./bin/<communication_type>/client <port_number> <int1> <int2> ... <int n>
```  
# bistream
### Example of *bi-directional streaming* API with server and client implementation. In this example, the server will return current max value of the numbers, that will be sent by the client.
---
```
protoc -Ibistream/proto --go_opt=module=github.com/cyberdr0id/go-grpc-practice --go_out=. --go-grpc_opt=module=github.com/cyberdr0id/go-grpc-practice --go-grpc_out=. bistream/proto/*.proto

go build -o bin/bistream/client.exe ./bistream/client
go build -o bin/bistream/server.exe ./bistream/server
```
To start application you must run `server.exe` file with port number in command line arguments:
```
./bin/<communication_type>/server <port_number>
```
And then run `client.exe` with the same port number and additional integer parameters
```
./bin/<communication_type>/client <port_number> <int1> <int2> ... <int n>
```  
