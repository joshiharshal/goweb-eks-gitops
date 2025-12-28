FROM golang:1.24.2 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o app .

FROM gcr.io/distroless/base
COPY --from=builder /app/app .
COPY --from=builder /app/static ./static
EXPOSE 8080
CMD ["./app"]
