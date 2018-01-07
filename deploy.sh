#! /bin/sh
kill -9 $(pgrep webserver)
cd /data/dev-go/17disney-etl
git pull
go build

cd webserver
./webserver &