FROM alpine

MAINTAINER ChenPeng <chenpeng@ghostcloud.cn>

LABEL Vendor="Ghostcloud" \
    Name="server" \
    Version="1.0.0" \
    Date="07/01/2016"

COPY bin /

EXPOSE 8080

ENTRYPOINT ["/ms-todo"]
