FROM golang:1.18-alpine3.16 as builder

WORKDIR /app
COPY . .
RUN go build -o go-fiber-hello


FROM scratch as runner
WORKDIR /app
EXPOSE 3000
COPY --from=builder /app/go-fiber-hello ./
COPY --from=builder /app/public ./public

CMD ["./go-fiber-hello"]