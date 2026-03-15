FROM golang:1.21
WORKDIR /app
COPY main.go .
RUN go mod init ric-security-xapp
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
