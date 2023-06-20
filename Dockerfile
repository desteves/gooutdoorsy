FROM golang:1.20

WORKDIR /usr/src/app

EXPOSE 8080

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
ENV GODEBUG=netdns=cgo
COPY . .
RUN go build -v -o /usr/local/bin/ ./...

CMD ["/usr/local/bin/gooutdoorsy"]