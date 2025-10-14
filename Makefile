.PHONY: tidy build run stop ps

tidy: ; go mod tidy
build: ; go build -o bin/users-server ./cmd/server
run:   ; GRPC_ADDR=:50051 go run ./cmd/server
ps:    ; /usr/bin/ps -W | grep users-server || true
stop:  ; /usr/bin/ps -W | awk '/users-server/ {print $$1}' | xargs -r kill
