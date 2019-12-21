# Golang gRPC API using PostgreSQL
  
### To start this API run `go run main.go` in the `server` and `client` file.  

#### Installation:  
Golang: https://golang.org/dl/  

Protocol Buffers: https://github.com/protocolbuffers/protobuf/releases  
  
> Recompile `.proto` file?  
> Open the `.proto` file and copy the command.  
  
> Already in this project >> `third_party` folder.  
> Only replace if you want to use another version of protoc  
> 1. Download the `protoc-{version}-win64.zip`  
> 2. Copy the folder inside the `include` folder  
> 3. Delete the current content from the `third_party` folder  
> 4. Paste it inside the `third_party` folder that is contained in this project  
>  
> **IMPORTANT:** Add `protoc` to your global path  
 > 1. Copy the `protoc.exe` from inside the `bin` folder > 2. Paste it inside `$GOPATH\go\bin` (a.e `C:\Users\{user}\go\bin`)  

Golang Protocol Buffers Compiler: https://github.com/golang/protobuf  

PostgreSQL: https://www.postgresql.org/download/ 

gRPC: `go get -u google.golang.org/grpc`  
  
Gin (Routing): `go get -u github.com/gin-gonic/gin`  
  
Other `go get -u {dependencies}` used? Just download them.
