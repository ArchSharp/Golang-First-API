FROM golang:1.21.1

WORKDIR /app

COPY go.mod .
COPY main.go .

# CGO_ENABLED=0 GOOS=windows 
RUN go build -o bin . 

ENTRYPOINT [ "/app/bin" ]