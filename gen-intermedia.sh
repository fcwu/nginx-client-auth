#!/bin/sh -xe

openssl req -new -newkey rsa:4096 -keyout intermedia.key -out intermedia.csr -nodes -subj '/CN=intermedia'
openssl x509 -req -sha256 -days 365 -in intermedia.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out intermedia.crt -extensions req_ext -extfile intermedia.ext

openssl verify -CAfile ca.crt intermedia.crt

openssl x509 -in intermedia.crt -text -noout

ls -l intermedia.*
