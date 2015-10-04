
SRC_FILES = server/*.go

all: $(SRC_FILES)
	go build -o serverChat server/main.go
