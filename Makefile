deploy:
	GOOS=linux GOARCH=arm GOARM=5 go build
	ssh pi@pi.local rm homekit
	scp hc-my-homekit pi@pi.local:~/homekit
	ssh pi@pi.local sudo systemctl restart homekit
