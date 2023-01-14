FROM golang:1.19 as builder

COPY . .

RUN unset GOPATH

RUN go install .
RUN go build -o app

EXPOSE 7001

RUN useradd ava
USER ava

CMD ["./app"]