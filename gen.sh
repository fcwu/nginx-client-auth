#!/bin/sh

# CERTDIR=.
# 
# openssl genrsa -des3 -passout pass:xxxx -out /tmp/server.pass.key 2048
# # fix permission on osx
# rm -f ${CERTDIR}/default.key
# openssl rsa -passin pass:xxxx -in /tmp/server.pass.key -out ${CERTDIR}/default.key
# chmod 400 ${CERTDIR}/default.key
# rm /tmp/server.pass.key
# openssl req -new -key ${CERTDIR}/default.key -out /tmp/server.csr -subj "/C=TW/ST=Taiwan R.O.C/L=Taipei/O=MOXA/OU=SYS Department/CN=moxa.com"
# openssl x509 -req -sha256 -days 3650 -in /tmp/server.csr -signkey ${CERTDIR}/default.key -out ${CERTDIR}/default.crt

# Generate the CA Key and Certificate
openssl req -x509 -sha256 -newkey rsa:4096 -keyout ca.key -out ca.crt -days 356 -nodes -subj '/CN=Fern Cert Authority'
openssl req -new -newkey rsa:4096 -keyout server.key -out server.csr -nodes -subj '/CN=10.144.48.106'
openssl x509 -req -sha256 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt -extensions req_ext -extfile v3.ext
openssl req -new -newkey rsa:4096 -keyout client.key -out client.csr -nodes -subj '/CN=10.144.48.106'
# openssl x509 -req -sha256 -days 365 -in client.csr -CA ca.crt -CAkey ca.key -set_serial 02 -out client.crt -extensions req_ext -extfile v3.ext
# openssl req -new -newkey rsa:4096 -keyout ubuntu.dorowu.com.key -out ubuntu.dorowu.com.csr -nodes -subj '/CN=ubuntu.dorowu.com'
# openssl x509 -req -sha256 -days 365 -in ubuntu.dorowu.com.csr -CA ca.crt -CAkey ca.key -set_serial 03 -out ubuntu.dorowu.com.crt -extensions req_ext -extfile v3.ext
