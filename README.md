Homekit
=======

My DIY HomeKit / HomeControl accessories.

Installation
============

Save the following file as ``/etc/systemd/system/homekit.service``:

    [Unit]
    Description=homekit bridge
    Wants=network-online.target
    After=network-online.target

    [Service]
    ExecStart=/home/pi/homekit
    User=pi
    Restart=on-failure
    Environment=ICY_USERNAME=XXX
    Environment=ICY_PASSWORD=XXX
    WorkingDirectory=/home/pi

    [Install]
    WantedBy=multi-user.target

Run:

    sudo systemctl enable homekit.service

Deploy
======

    make

