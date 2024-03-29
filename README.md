# Fuks Doorman System

The purpose of the Fuks Doorman system is to allow fuks members access to the fuks office
by using RFID tags in student ID cards.

The system is assembled out of a Raspberry Pi, an RFID reader, a relay module and an electronic door opener.
Members can register the RFID number from their KIT-Card at Google Workspace.
The RFID reader placed at the office door can then read these numbers and pass them to the Raspberry Pi.
There the numbers will be checked and authenticated.
If the number can be matched to a fuks user, the door will be opened by using a simple relay and GPIO pins.

![Example](photo.jpg)

## Credentials and Certificates

This section describes how to generate the required credentials and certificates for the Doorman system. All generated
files can be found in the Google Drive folder `05_Team_IT/02_Interne Projekte/07_Doorman`.

### Google Workspace

The Doorman system needs access to Google Workspace APIs to fetch authorized RFID chip-numbers.
To do this a **service account** with **domain-wide delegation** is required.

The domain-wide delegation needs the following OAuth scopes:

```
https://www.googleapis.com/auth/admin.directory.user,
https://www.googleapis.com/auth/admin.directory.userschema,
https://www.googleapis.com/auth/admin.directory.group,
https://www.googleapis.com/auth/admin.directory.group.member,
https://www.googleapis.com/auth/spreadsheets.readonly
````

Follow the
instructions [here](https://developers.google.com/workspace/guides/create-credentials#create_credentials_for_a_service_account)
to create a new credentials.json.
The generated JSON must be placed under `fuks/credentials.json` and will be included in the compiled GO executable.

> Note that you might need to update the email address field **config.Subject** in `fuks/fuks.go`

### Firebase

The Doorman system uses Firebase to authenticate users. To do this a **service account** is required. To create a new
service account follow the instructions:

1. Go to the [Firebase Console](https://console.firebase.google.com/)
2. Select the project `fuks-app`
3. Go to `Project Settings`
4. Go to `Service accounts`
5. Click on `Generate new private key`
6. Place the downloaded JSON file under `server/firebase-credentials.json`

### TLS Certificates

The Doorman system uses TLS certificates to secure the gRPC connection between the Doorman and the Fuks App.
To generate new certificates follow the instructions:

1. Update parameters in `certificate/Makefile`
2. Run `make` in `certificate/`
3. Update copy the contents of `certificate/doorman-cert.pem` to `fuks_app/lib/services/doorman_cert.dart`
4. (Update common name in `fuks_app/lib/services/doorman.dart`)

## Create a new release

Prepare a new release by following these steps:

1. Update the changelog in `CHANGELOG.md`
2. Update dependencies `go get -u all`
3. Commit changes `git commit -am "Release vX.X.X"`
4. Push changes `git push`
5. Create a new git tag:
    1. `git tag -a vX.X.X -m "Release vX.X.X"`
    2. `git push origin vX.X.X`

## Compile and install Doorman CLI

```shell
git clone https://github.com/fuks-kit/doorman/

git pull
git checkout vX.X.X

go get all

go install cmd/doorman/doorman.go
go install cmd/rfid/doorman_rfid.go
go install cmd/door/door.go
```

> Note: Add ```export PATH=$PATH:$HOME/go/bin/``` to ```.bashrc```

### Create doorman service

```shell
sudo mkdir -p /etc/doorman/
sudo cp config.json /etc/doorman/config.json
sudo cp fallback_access.json /etc/doorman/fallback-access.json

sudo cp doorman.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo service doorman start
```

Run ```sudo systemctl enable doorman``` to start the doorman service on startup

### Cross compile executables

```shell
rm -rf bin
mkdir -p bin
GOOS=linux GOARCH=arm go build -o bin/doorman cmd/doorman/doorman.go
scp bin/doorman fuks@10.0.0.238:~/go/bin
```

## Generate gRPC Definitions

### Dependencies

Before you begin, make sure you have the following dependencies installed:

- Protocol Buffers: Install with Homebrew (macOS) or your preferred package manager.
   ```bash
   brew install protobuf
   ```

- Go Protobuf and gRPC code generation tools:
   ```bash
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   ```

### Generate Code

To generate gRPC definitions, follow these steps:

1. Update APP_DIR Variable
    - In the `proto/Makefile`, update the `APP_DIR` variable to point to the Fuks App directory.

2. Update PROTO_ROOT_DIR Variable
    - If necessary, modify the `PROTO_ROOT_DIR` variable in the `proto/Makefile` to suit your setup.

3. Update gRPC Definitions
    - Make changes to the gRPC definitions in `proto/doorman.proto` as needed.

4. Generate Code
    - Use the following commands to generate the code:
        - Generate Go code:
          ```bash
          make go
          ```
        - Generate Fuks App code:
          ```bash
          make dart
          ```

## Troubleshooting

```shell
# View logs
cat /var/log/doorman.log

# Find Raspberry Pi
sudo nmap -PE 10.0.0.0/24

# Clean up debug procedure
sudo service doorman stop
sudo rm /var/log/doorman.log
sudo rm ~/doorman-recovery.json
```

### GPIO permission

After running `sudo rpi-update` the error `Error: open /dev/gpiomem: permission denied` can be thrown. This can be fixed by:

```shell
# Add user to gpio group
sudo adduser fuks gpio

# Change permissions of /dev/gpiomem
sudo chown root.gpio /dev/gpiomem
sudo chmod g+rw /dev/gpiomem
```
