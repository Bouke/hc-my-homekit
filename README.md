Homekit
=======

My DIY HomeKit / HomeControl accessories.

Deploy
======

    GOOS=linux GOARCH=arm GOARM=5 go build
    ssh pi@10.0.1.10 rm homekit
    scp homekit pi@10.0.1.10:~/homekit
    ssh pi@10.0.1.10 sudo systemctl restart homekit
