#! /bin/bash

set -ex

cd `dirname $0`
 	
sudo rm -rf /etc/mysql/conf.d
sudo cp -rL conf/mysql/conf.d /etc/mysql/conf.d
sudo rm -rf /etc/mysql/mysql.conf.d
sudo cp -rL conf/mysql/mysql.conf.d /etc/mysql/mysql.conf.d
sudo rm -rf /etc/mysql/my.cnf
sudo cp -rL conf/mysql/my.cnf /etc/mysql/my.cnf

sudo systemctl restart mysql
