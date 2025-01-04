FROM golang:1.23.4

COPY . /

WORKDIR /

RUN go build main.go

CMD ["./main"]

EXPOSE 8081
