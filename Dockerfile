# Environment
FROM golang:1.16 as build-env

RUN mkdir -p /opt/http
WORKDIR /opt/http
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /opt/service/http

# Release
FROM alpine:latest
COPY --from=build-env /opt/service/http /bin/http
ENTRYPOINT ["/bin/http"]