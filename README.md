Tested for lunix running Ubuntu

To run locally:

Required libraries:
    
    sudo apt update
    sudo apt install golang

To build and install:

    git clone git@github.com:Nathanialw/homeserver.git
    go build -buildvcs=false -o ../../app/main
    cd webserver/app
    ./main --install

Defualt runs localhost:10002:

    ./main

Will accept a port as an arguement:

    ./main 4000 #runs on port 4000