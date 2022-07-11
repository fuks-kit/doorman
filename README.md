# Fuks Doorman System

## Setup Google Workspace credentials

TODO

## Cross compile executables

```shell
rm -rf build
mkdir -p build
GOOS=linux GOARCH=arm go build -o build/fuks_doorman cmd/doorman/doorman.go
```
