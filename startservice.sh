sudo cp ddechain.service  /etc/systemd/system
sudo systemctl daemon-reload
sudo systemctl start ddechain
sudo systemctl enable ddechain
sudo systemctl status ddechain
