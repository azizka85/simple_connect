# simple_connect
 Using gRPC on the browser and server, unary and server streaming

# Installation
- install go compiler
- install nodejs and npm
- install additional tools: https://connect.build/docs/go/getting-started/

# How to use
- run go project: `go run .`
- change to client directory: `cd ./client`
- install client libraries: `npm i`
- run client: `npm start`

# Switch from http/1 to http/2
- in `main.go`, comment from line 150 to 155, and uncomment from 142 to 149
- in `./client/.env`, change from `http` to `https` for `API_GRPC`
