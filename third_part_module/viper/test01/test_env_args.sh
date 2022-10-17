#!/usr/bin/env bash
export MYSQL_IP=192.168.1.100
export MYSQL_PORT=3306
export redis_port=6666

go build  -gcflags "all=-N -l" -o go_build_test01

if [ "$1" = "debug" ]; then
  dlv --listen=":2345"  --headless=true --api-version=2 --check-go-version=false --only-same-user=false exec ./go_build_test01 -- --mq.ip 192.168.1.101 --mq.port 3389
else
  ./go_build_test01 --mq.ip 192.168.1.101 --mq.port 3389
fi
