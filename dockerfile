# syntax=docker/dockerfile:1
FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /goweb
EXPOSE 1323
CMD ["/goweb"]