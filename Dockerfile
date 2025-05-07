FROM golang:1.24.3-alpine3.21 AS build-env
 
ENV APP_NAME=demo-api
ENV CMD_PATH=main.go
 
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

# Build application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH
 
## Run Stage ##
FROM alpine:3.21.3
 
ENV APP_NAME=demo-api
COPY --from=build-env /$APP_NAME .
 
EXPOSE 8081
 
CMD ["./$APP_NAME"]