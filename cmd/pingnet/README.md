
## 编译
### windows
```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```
### linux
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
```

## 使用
```bash
pingnet 3.1.20.0/24
```