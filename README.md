# Project presents implementation of different gRPC communications.

### Makefile uses for **Windows**  OS only  
  
    

First of all you should build project
```
make <communication_type>
```

This command will genereate Go code from `.proto` file and `bin` folder will be created with folder `<communication_type>` inside, which will contain two binary files `client.exe` and `server.exe`.

To start application you must run `server.exe` file with port number in command line arguments:
```
./bin/<communication_type>/server <port_number>
```
And after run `client.exe` with the same port number.  
```
./bin/<communication_type>/client <port_number>
```