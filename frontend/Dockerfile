FROM nginx:1.29.0-alpine

RUN rm /etc/nginx/conf.d/default.conf

COPY nginx.conf.template /etc/nginx/templates/default.conf.template

COPY index.html /usr/share/nginx/html/
COPY style.css /usr/share/nginx/html/
COPY script.js /usr/share/nginx/html/

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]