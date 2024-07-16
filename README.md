# roundrobin
Simple implementation for round robin algorithm

There are 2 types of service:
- Routing: The routing service that will handle requests using round robin for `/request` API (default port: 8000).
- Simple Server: The server that will be registered to the routing service and contains `/request` API (default port: 8100)

All service's port can be configurable using `-port` flag

## Additional APIs
### Routing
- `POST /register`: To register the service URL (Plaintext body, full base URL with `http://`)

### Simple Server
- `POST /maintenance`: To set the service in maintenance mode (x-www-form-urlencoded, `maintenance` key with boolean value)
- `POST /delay`: To set the `/request` response delay (x-www-form-urlencoded, `delay` key with numeric milliseconds value)

## How to Run
- Install go (1.22.5, lower version might work as well)
- `go mod vendor` & `go mod tidy`
- run the routing service `go run cmd/routing/main.go`
- run the simple service(s) `go run cmd/simple-service/main.go`, `go run cmd/simple-service/main.go -port=8101`, `go run cmd/simple-service/main.go -port=8102`
