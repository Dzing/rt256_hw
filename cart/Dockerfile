FROM golang:1.25.6

ADD ./build/bin/  /app
ADD ./conf/config.yaml /app/cart-config.yaml

ENV CART_CONF="/app/cart-config.yaml"

CMD ["/app/route-cart"]
