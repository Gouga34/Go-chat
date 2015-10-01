COMMON_FILES = common/socket.go common/constants.go server/room/*.go common/user.go

all: serverChat clientChat

serverChat:	server/server.go $(COMMON_FILES)
	go build -o serverChat server/server.go

clientChat: client/client.go $(COMMON_FILES)
	go build -o clientChat client/client.go
