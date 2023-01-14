# https://www.bacancytechnology.com/blog/dockerize-golang-application

# Build Stage
# First pull Golang image
FROM golang:1.19-alpine as build-env
 
# Set environment variable
ENV APP_NAME app
ENV CMD_PATH main.go
 
# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

# Build application
RUN GIN_MODE=release CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH
 
# Run Stage
FROM debian:bullseye-slim
 
# Set environment variable
ENV APP_NAME app
 
# Copy only required data into this image
COPY --from=build-env /$APP_NAME .
COPY data data


# Expose application port
EXPOSE 7001

RUN useradd ava
USER ava

# Start app
CMD ./$APP_NAME