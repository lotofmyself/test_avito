FROM debian

RUN apt-get update && apt-get install -y nginx 
COPY nginx.conf /etc/nginx
EXPOSE 8089
CMD ["nginx","-g","daemon off;"]
