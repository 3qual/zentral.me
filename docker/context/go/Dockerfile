FROM golang:1.20-alpine

WORKDIR /app

CMD ["sh", "-c", "go run main.go | awk '{ print strftime(\"[%Y-%m-%d %H:%M:%S]\"), $0 }' | tee -a /logs/go.log"]

EXPOSE 8080
