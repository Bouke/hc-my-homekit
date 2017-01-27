deploy:
	GOOS=linux GOARCH=arm GOARM=7 go build
	-ssh pi@pi3.local rm homekit
	scp hc-my-homekit pi@pi3.local:~/homekit
	ssh pi@pi3.local sudo systemctl restart homekit
