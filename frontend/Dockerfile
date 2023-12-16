FROM nginx:latest

LABEL maintainer="Sachin Parihar <sachinparihar98120@gmail.com>"

RUN rm -rf /usr/share/nginx/html/*

COPY . /usr/share/nginx/html

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]
