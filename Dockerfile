FROM golang:1.19 AS stage

WORKDIR /app

COPY go.mod go.sum ./
# RUN go mod download

COPY . .

# Build aplikasi Go
RUN go build -o app .
# RUN CGO_ENABLED=0 GOOS=linux go build -o /app

EXPOSE 8080

CMD ["./app"]
