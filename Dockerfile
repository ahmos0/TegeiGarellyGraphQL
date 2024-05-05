FROM golang:1.20 as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN GOOS=linux GOARCH=amd64 go build -mod=readonly -v -o server

CMD ["/app/server"]

EXPOSE 8080

