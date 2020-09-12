#! /bin/bash


sudo mkdir -p /data/logs/mysql
sudo mkdir -p /data/logs/nginx

D=`date -Iminutes`
MF="/data/logs/mysql/mysql-slow_${D}.log"
NF="/data/logs/nginx/access_${D}.log"

set -ex

sudo mv /var/log/mysql/mysql-slow.log $MF
sudo chmod 666 $MF
sudo systemctl restart mysql

sudo mv /var/log/nginx/access.log $NF
sudo chmod 666 $NF
sudo systemctl restart nginx

