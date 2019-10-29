FROM golang:1.13.3-alpine
ADD . /app
WORKDIR /app
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod download \
    && go build -o server ./server
CMD ["./server"]
