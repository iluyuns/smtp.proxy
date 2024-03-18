FROM iuupub/alpine_shanghai:3.19

COPY ./smtp.proxy.server.app /var/bin/smtp.proxy.server.app

ARG token="123456"
ARG host="smtp.exmail.qq.com"
ARG port="587"
ARG username="value"
ARG password="123456"

ENV SMTP_PROXY_TOKEN=${token}
ENV SMTP_PROXY_HOST=${host}
ENV SMTP_PROXY_PORT=${port}
ENV SMTP_PROXY_USERNAME=${username}
ENV SMTP_PROXY_PASSWORD=${password}

EXPOSE 80

ENTRYPOINT [ "/var/bin/smtp.proxy.server.app" ]
