FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go get .
RUN go build -o our-home-server .

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/our-home-server .
EXPOSE 3001
CMD ["./our-home-server"]