FROM golang:1.12.2-alpine3.9 AS builder
RUN apk --no-cache add build-base
COPY . /code
RUN mkdir -p /usr/local/go/src/github.com/zjyl1994 && \
    ln -s /code /usr/local/go/src/github.com/zjyl1994/telegram-push-bot && \
    cd /usr/local/go/src/github.com/zjyl1994/telegram-push-bot && \
    CGO_ENABLED=1 go build -a
FROM alpine:latest
RUN apk --no-cache add tzdata ca-certificates libc6-compat libgcc libstdc++
COPY --from=builder /usr/local/go/src/github.com/zjyl1994/telegram-push-bot/telegram-push-bot /app/app
CMD ["/app/app"]