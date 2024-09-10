protoc: 
	protoc --go_out=. proto/*.proto

msg:
	go run main.go