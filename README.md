Tested for lunix running Ubuntu

To run locally:

Required libraries:
    
    sudo apt update
    sudo apt install golang-go
    sudo apt install python3-pip
    sudo apt install xdg-utils

To build and install:

    git clone https://github.com/Nathanialw/homeserver.git
    cd homeserver
    chmod +x build.sh
    ./build.sh

if xdg-open fails to find a default browser, assuming you are using the default port, you can manually navigate to:

    http://localhost:10002

Will accept a port as an argument:

    cd homserserver/webserver/app
    ./main 4000 #runs on port 4000
