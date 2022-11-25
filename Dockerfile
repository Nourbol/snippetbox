FROM golang:1.19.3-alpine
WORKDIR /app
COPY . .
RUN go mod download
ENTRYPOINT ["go", "run", "./cmd/web"]