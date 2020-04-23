#!/bin/sh

docker rm -f mynginx3; 
docker run --name mynginx3 \
    -v ${PWD}/html:/usr/share/nginx/html:ro \
    -v ${PWD}/conf:/etc/nginx/conf.d \
    -v ${PWD}:/etc/nginx/client_certs/ \
    -p 10082:80 -p 10083:443 -d nginx
