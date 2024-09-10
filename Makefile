protoc: 
	protoc --go_out=. proto/*.proto

msg: protoc
	go run genMsg.go