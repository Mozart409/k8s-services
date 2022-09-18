FROM golang:1.18-alpine3.16 as builder

WORKDIR /go/src/app
COPY . .
RUN go build -o go-web-server


FROM alpine:3.16
WORKDIR /app
COPY --from=builder /go/src/app ./
EXPOSE 3000
CMD ["./go-web-server"]