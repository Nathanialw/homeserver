git checkout -- *
git pull
sudo chmod a+x webserver/app/main
sudo chmod 777 webserver/db/homeserver.sqlite3
sudo chmod 777 webserver/db
sudo service homeserver restart
sudo service homeserver status