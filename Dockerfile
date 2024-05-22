FROM golang:alpine AS builder
WORKDIR /build/
COPY . .
RUN go mod download
RUN go build -o pass-gen cmd/main/main.go

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /build/pass-gen .
EXPOSE 80
CMD ["./pass-gen"]
