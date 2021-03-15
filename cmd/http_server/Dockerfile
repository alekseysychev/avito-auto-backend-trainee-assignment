# Environment
FROM golang:1.16 as build-env

RUN mkdir -p /opt/avito-auto-backend-trainee-assignment

WORKDIR /opt/avito-auto-backend-trainee-assignment

COPY . .

RUN go mod download

RUN pwd

RUN CGO_ENABLED=0 go build -o /opt/service/http_server /opt/avito-auto-backend-trainee-assignment/cmd/http_server

# Release
#FROM alpine:latest
#COPY --from=build-env /opt/service/http_server /bin/http_server
#ENTRYPOINT ["/bin/http_server"]