# Environment
FROM golang:1.16 as build-env

RUN mkdir -p /opt/http_server
WORKDIR /opt/http_server
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /opt/service/http_server

# Release
FROM alpine:latest
COPY --from=build-env /opt/service/http_server /bin/http_server
ENTRYPOINT ["/bin/http_server"]