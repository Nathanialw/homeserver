git checkout -- *
git pull
chmod a+x webserver/app/main
chmod 777 webserver/db/homeserver.sqlite3
chmod 777 webserver/db
sudo service homeserver restart
sudo service homeserver status