FROM golang:1.16.5 as builder
WORKDIR /code
COPY go.mod go.sum /code/
RUN go version \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod download
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o httpservers .

FROM alpine:3.7
COPY --from=builder /code/httpservers /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/httpservers"]




