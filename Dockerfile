FROM nginx:1.19.0
EXPOSE 80

RUN rm /etc/nginx/conf.d/default.conf
ADD server.conf /etc/nginx/conf.d/server.conf
