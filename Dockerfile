FROM golang:1.14-buster as builder

CMD mkdir -p /tmp/excel-import
WORKDIR /tmp/excel-import
COPY . .

RUN go mod tidy
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o main main.go

FROM alpine

COPY --from=builder /tmp/excel-import /
CMD ["/main"]
