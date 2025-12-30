### Swagger

```
swag init -g ./cmd/server/main.go -o cmd/docs
```

### Run

```
go run cmd/server/main.go
```

### User management

```
go run ./cmd/usercli create -u nabi --admin
go run ./cmd/usercli list
go run ./cmd/usercli change-password -u nabi
go run ./cmd/usercli delete -u nabi
```
