FROM golang:1.24

ADD ./bin/app /app
ADD ./config.yaml /app/config.yaml

ENV LOMS_CONF="/app/config.yaml"

CMD ["/app/loms"]