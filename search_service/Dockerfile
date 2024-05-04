FROM golang:1.22.1-alpine3.19 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main ./cmd/main.go

FROM alpine:latest AS final

WORKDIR /

COPY --from=build /bin/main /main

EXPOSE 8080

ENTRYPOINT ["/main"]