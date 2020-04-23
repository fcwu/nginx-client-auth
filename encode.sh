#!/bin/bash

echo "some data" > sign.txt

# openssl enc -nosalt -aes-256-cbc -pbkdf2 -k hello-aes -P
key=`openssl enc -nosalt -aes-256-cbc -k hello-aes -pbkdf2 -P | sed -e '/^key/!d' -e 's/^key=//'`
iv=`openssl enc -nosalt -aes-256-cbc -k hello-aes -pbkdf2 -P | sed -e '/^iv/!d' -e 's/^iv =//'`

openssl enc -aes-256-cbc -k $key -iv $iv -pbkdf2 -in sign.txt -out sign.txt.enc

openssl enc -d -aes-256-cbc -k $key -iv $iv -pbkdf2 -in sign.txt.enc
