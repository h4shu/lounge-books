FROM golang:1.22.6 AS builder
COPY . /app
WORKDIR /app
RUN go mod download
RUN go mod verify
RUN go build -o server .

FROM golang:1.22.6
COPY --from=builder /app/server /app/server
WORKDIR /app
EXPOSE 8080
CMD ["/app/server"]