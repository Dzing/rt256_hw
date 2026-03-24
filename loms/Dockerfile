FROM golang:1.25.6

ADD ./build/bin/ /app
ADD ./conf/config.yaml /app/loms-config.yaml

ENV LOMS_CONF="/app/loms-config.yaml"

CMD ["/app/route-loms"]