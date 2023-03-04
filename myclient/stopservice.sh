sudo rm  /etc/systemd/system/deyichain.service
PID_DEYICHAIN=$(ps -ef | grep deyichain | grep -v grep | awk '{print $2}')
sudo kill -9 $PID_DEYICHAIN
sudo systemctl daemon-reload
