# users-service

gRPC сервис пользователей (GORM + SQLite).
## Запуск
```bash
make build && ./bin/users-server

gRPC примеры (grpcurl)
grpcurl -plaintext -d '{"email":"alice@example.com","name":"Alice"}' localhost:50051 user.UserService/CreateUser
grpcurl -plaintext -d '{"page":1,"page_size":10}' localhost:50051 user.UserService/ListUsers
grpcurl -plaintext -d '{"id":1}' localhost:50051 user.UserService/GetUser
grpcurl -plaintext -d '{"user":{"id":1,"email":"alice2@example.com","name":"Alice 2"}}' localhost:50051 user.UserService/UpdateUser
grpcurl -plaintext -d '{"id":1}' localhost:50051 user.UserService/DeleteUser

