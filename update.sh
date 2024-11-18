

git pull

go build main.go

mkdir -p /etc/iputility

cp main /etc/iputility/iputility

cp iputility.service /etc/systemd/system/iputility.service

systemctl daemon-reload

systemctl restart iputility

systemctl enable iputility

echo "Update complete"