all: serverChat clientChat

serverChat:	server/server.go
	go build -o serverChat server/server.go

clientChat: client/client.go
	go build -o clientChat client/client.go
