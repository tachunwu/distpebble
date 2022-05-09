go run main.go -port 30001 -nats localhost:4001 -cluster localhost:4002,localhost:4003
go run main.go -port 30002 -nats localhost:4002 -cluster localhost:4001,localhost:4003
go run main.go -port 30003 -nats localhost:4003 -cluster localhost:4001,localhost:4002