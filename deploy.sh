#! /bin/sh

kill -9 $(pgrep 17disney-etl)
cd /data/deb-go/
git pull https://github.com/xanke/17disney-etl

cd 17disney-etl
./17disney-etl $