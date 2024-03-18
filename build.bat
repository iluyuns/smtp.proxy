set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64

go build -ldflags "-X main.version=1.0.3" -o smtp.proxy.server.app ./server/smtp.proxy.server.app.go

docker build -t iuupub/smtp.proxy:1.0.3 .