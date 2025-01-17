#!/bin/bash

cd webserver/src/main
go build -buildvcs=false -o ../../app/main
cd ../../app

pip install -r ../scripts/requirements.txt

#checking it the arguemnt is a valid port number
if [[ "$1" =~ ^[0-9]+$ ]] && [ "$1" -ge 1 ] && [ "$1" -le 65535 ]; then
    port=$1
else
    echo "Invalid port number. Please provide a port number between 1 and 65535. Defaulting to 10002."
    port=10002
fi

#downloads and converts the IMDB text database into a friendlier format
./main --install

#run on port $port
./main $port
xdg-open http://localhost:$port