FROM golang:1.21-alpine as build

RUN apk add --no-cache git make ca-certificates alpine-sdk npm

COPY . /app

WORKDIR /app

RUN make install_dependencies && make statik_landing && make wire && make build_server

FROM alpine:latest

ENV BORE_DOMAIN=bore.digital BORE_HTTPADDR=0.0.0.0:2000

COPY --from=build /etc/ssl/certs /etc/ssl/certs
COPY --from=build /app/build/bore-server /bore-server

EXPOSE 2000 2200 55000-65000

CMD ["/bore-server"]
