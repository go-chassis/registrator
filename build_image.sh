#!/bin/bash
set -e 
set -x 
export GOPROXY=https://goproxy.io
GO111MODULE=on go mod vendor
sudo docker build -t gomesh/registrator .

