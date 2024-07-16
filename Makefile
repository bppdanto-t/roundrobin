routing: go run cmd/routing/main.go & sleep 2

server: go run cmd/simple-service/main.go
server2: go run cmd/simple-service/main.go -port=8101
server3: go run cmd/simple-service/main.go -port=8102

run: routing server server2 server3