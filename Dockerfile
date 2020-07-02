FROM alpine:3.12.0
EXPOSE 80

WORKDIR /var/app
ADD ipserver ipserver

CMD ./ipserver
