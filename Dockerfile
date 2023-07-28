FROM golang as builder

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOPROXY=https://goproxy.cn

WORKDIR /app

RUN apk add --no-cache make

COPY go.mod .
COPY go.sum .

RUN make all

FROM alpine

COPY --from=builder /app/bin/retalk /

EXPOSE 3000

CMD ["/bin/retalk"]
