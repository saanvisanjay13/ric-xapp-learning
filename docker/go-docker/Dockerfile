FROM golang:1.21
WORKDIR /app
COPY . . 
RUN echo 'package main\nimport "fmt"\nfunc main() { fmt.Println("Hello from Go in Docker!") }' > main.go
RUN go mod init hello
RUN go build -o hello .
CMD ["./hello"]

