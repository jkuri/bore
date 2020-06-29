FROM golang:1.14-alpine as build

RUN apk add --no-cache git make ca-certificates alpine-sdk

COPY . /app

WORKDIR /app

RUN make install_dependencies && make

FROM scratch

ENV BORE_DOMAIN=jankuri.com BORE_HTTPADDR=0.0.0.0:80

COPY --from=build /etc/ssl/certs /etc/ssl/certs
COPY --from=build /app/build/bore-server /bore-server

EXPOSE 80 2200

CMD ["/bore-server", "-config", "/bore/bore-server.yaml"]
