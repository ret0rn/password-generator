FROM golang:1.18.3
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
EXPOSE 80
RUN go build -o /password-generator cmd/main/main.go
CMD ["/password-generator"]
