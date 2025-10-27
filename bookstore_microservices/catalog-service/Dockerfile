FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o catalog

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/catalog /app/catalog
COPY catalog.json /app/catalog.json
EXPOSE 4000
CMD ["/app/catalog"]
