#! /bin/sh

kill -9 $(pgrep webserver)
cd /data/deb-go/
git pull https://github.com/xanke/17disney-etl

cd webserver
./webserver $