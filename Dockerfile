FROM golang:1.14-alpine as build

RUN apk add --no-cache git make ca-certificates alpine-sdk yarn

COPY . /app

WORKDIR /app

RUN make install_dependencies && make statik_landing && make wire && make build_server

FROM scratch

ENV BORE_DOMAIN=bore.network BORE_HTTPADDR=0.0.0.0:80

COPY --from=build /etc/ssl/certs /etc/ssl/certs
COPY --from=build /app/build/bore-server /bore-server

EXPOSE 80 2200 55000-65000

CMD ["/bore-server", "-config", "/bore/bore-server.yaml"]
