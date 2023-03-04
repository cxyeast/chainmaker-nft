sudo cp deyichain.service  /etc/systemd/system
sudo systemctl daemon-reload
sudo systemctl start deyichain
sudo systemctl enable deyichain
sudo systemctl status deyichain
