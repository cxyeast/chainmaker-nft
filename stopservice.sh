sudo rm  /etc/systemd/system/ddechain.service
PID_DEYICHAIN=$(ps -ef | grep ddechain | grep -v grep | awk '{print $2}')
sudo kill -9 $PID_DEYICHAIN
sudo systemctl daemon-reload
