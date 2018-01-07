#! /bin/sh
kill -9 $(pgrep webserver)
cd /data/dev-go/17disney-etl
git pull

cd webserver
go build
./webserver &