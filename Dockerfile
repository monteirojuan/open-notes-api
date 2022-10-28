FROM golang:1.18 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest as runner

COPY --from=builder /app .

EXPOSE 8000

ENTRYPOINT [ "./main" ]