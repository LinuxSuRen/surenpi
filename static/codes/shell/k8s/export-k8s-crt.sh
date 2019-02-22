#!/bin/sh
# You need to execute this file in the k8s nodes
cat /etc/kubernetes/admin.conf | grep certificate-authority-data | awk '{print $2}' | base64 --d > ca.crt
cat /etc/kubernetes/admin.conf | grep client-certificate-data | awk '{print $2}' | base64 --d > client.crt
cat /etc/kubernetes/admin.conf | grep client-key-data | awk '{print $2}' | base64 --d > client.key
openssl pkcs12 -export -out cert.pfx -inkey client.key -in client.crt -certfile ca.crt
tar czvf ca.tr.gz ca.crt client.key client.crt
rm -rf ca.crt client.key client.crt