FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go mod init test

RUN go get github.com/1Password/connect-sdk-go

RUN go mod tidy

RUN go build

CMD ["go", "run", "main.go"]
