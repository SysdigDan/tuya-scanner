FROM golang:1.17.1-alpine3.14 as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN go get ./...
RUN CGO_ENABLED=0 go build -o tuya-scanner -a -ldflags '-extldflags "-static"' ./cmd/worker/

FROM alpine:3.14

RUN addgroup -S worker && adduser -S worker -G worker

WORKDIR /app

COPY --from=builder build/tuya-scanner .

EXPOSE 9265/tcp
EXPOSE 6666/udp
EXPOSE 6667/udp

USER worker

CMD ["/app/tuya-scanner"]
