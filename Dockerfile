FROM golang:1.24.3-alpine3.21 AS build-env
ENV APP_NAME=demo-api
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -v -o /demo-api main.go

FROM alpine:3.21.3
WORKDIR /
RUN ln -sf /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime

COPY --from=build-env /demo-api ./

EXPOSE 8080

ENTRYPOINT ["./demo-api"]
