#!/bin/sh -xe

# introduction: http://qbsuranalang.blogspot.com/2017/05/openssl-certificate-chain-included-san.html
# opensll command to pem/der/pfx format: https://ssorc.tw/7142/openssl-%E6%8C%87%E4%BB%A4-command-line-%E8%BD%89%E6%AA%94-pem-der-p7b-pfx-cer/
# verify cert by golang: https://golang.org/pkg/crypto/x509/#example_Certificate_Verify

openssl req -x509 -sha256 -newkey rsa:4096 -keyout ca.key -out ca.crt -days 356 -nodes -subj '/CN=Doro Cert Authority'

openssl x509 -in ca.crt -text -noout

ls -l ca.*
