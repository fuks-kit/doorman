# Fuks Doorman System

## Install doorman cli

```shell
go install cmd/doorman/doorman.go
go install cmd/rfid/rfid.go
go install cmd/door/door.go
```

## Create new service

```shell
sudo cp doorman.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo service doorman start
```

Run ```sudo systemctl enable doorman``` to start the doorman service on startup

## Setup Google Workspace credentials

TODO

## Cross compile executables

```shell
rm -rf build
mkdir -p build
GOOS=linux GOARCH=arm go build -o build/fuks_doorman cmd/doorman/doorman.go
```
