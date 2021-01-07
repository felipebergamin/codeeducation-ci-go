FROM golang:latest AS builder

WORKDIR /app
COPY . .
RUN go build -o main

FROM busybox:latest

COPY --from=builder /app/main /app/main

ENTRYPOINT [ "/app/main" ]
