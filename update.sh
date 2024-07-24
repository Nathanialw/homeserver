git checkout -- *
git pull
cd /var/www
sudo service homeserver restart
sudo service homeserver status
sudo chmod a+x homeserver/webserver/app/main
sudo chmod 777 homeserver/webserver/db/homeserver.sqlite3
sudo chmod 777 homeserver/webserver/db