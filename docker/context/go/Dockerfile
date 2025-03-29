FROM golang:1.24.0-alpine

WORKDIR /app

CMD ["sh", "-c", "go run cmd/server/main.go | awk '{ print strftime(\"[%Y-%m-%d %H:%M:%S]\"), $0 }' | tee -a /logs/go.log"]

EXPOSE ${GO_APP_EXTERNAL_PORT}
