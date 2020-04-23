#!/bin/bash -xe

openssl req -new -newkey rsa:4096 -keyout final.key -out final.csr -nodes -subj '/CN=final'

openssl x509 -req -sha256 -days 365 -in final.csr -CA intermedia.crt -CAkey intermedia.key -CAcreateserial -out final.crt -extensions req_ext -extfile final.ext

openssl verify -CAfile <(cat ca.crt intermedia.crt) final.crt

openssl x509 -in final.crt -text -noout

ls -l final.*
