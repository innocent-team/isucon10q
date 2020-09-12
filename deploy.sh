#! /bin/bash

set -ex

cd `dirname $0`

cd /home/isucon/isuumo/webapp/go
go build
sudo systemctl restart isuumo.go
