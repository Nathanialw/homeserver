git checkout -- *
git pull
chmod a+x webserver/app/main
chmod 777 webserver/db/homeserver.sqlite3
sudo service homeserver restart
sudo service homeserver status