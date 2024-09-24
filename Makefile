protoc: 
	protoc --go_out=. msg/proto/*.proto

msg: protoc
	go run main.go genMsg

install: msg
	go build -o server main.go