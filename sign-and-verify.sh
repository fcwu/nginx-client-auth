#!/bin/bash -ex

echo "some data..." > sign.txt

openssl dgst -sha256 -sign final.key -out sign.txt.sha256 sign.txt

openssl dgst -sha256 -verify  <(openssl x509 -in final.crt  -pubkey -noout) -signature sign.txt.sha256 sign.txt

sed -i 's/some/some1/' sign.txt

openssl dgst -sha256 -verify  <(openssl x509 -in final.crt  -pubkey -noout) -signature sign.txt.sha256 sign.txt
