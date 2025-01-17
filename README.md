Tested for lunix running Ubuntu

<h1>Linux Home Server</h1>

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
    ./main 4000                         #runs on port 4000

<h1>Usage</h1>

At home I deploy it on a server with a static local address. that way anyone can set it as a bookmark and access it from any device on the network.

Apache2 reverse proxy configuration:

```apache
ServerAdmin webmaster@localhost
ServerName 192.168.0.2
ServerAlias 192.168.0.2
DocumentRoot /var/www/homeserver

ProxyPass / http://127.0.0.1:10002/
ProxyPassReverse / http://127.0.0.1:10002/
