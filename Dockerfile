FROM golang:1.9.0 as builder

# install Dep
RUN go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/nguyenhoaibao/logsvc

COPY Gopkg.toml Gopkg.toml
COPY Gopkg.lock Gopkg.lock

RUN dep ensure --vendor-only

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o logsvc cmd/server/main.go


# use a minimal alpine image
FROM alpine:3.7

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /root

COPY --from=builder /go/src/github.com/nguyenhoaibao/logsvc/logsvc .

CMD ["./logsvc"]
