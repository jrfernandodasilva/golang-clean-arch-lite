#!/bin/sh

if ! command -v go; then
    apk update

    # Add apps
    apk add curl wget nano git go

    # Add go binaries dir to $PATH
    export GOPATH=$(go env GOPATH)
    export PATH=$PATH:$GOPATH/bin
fi

# Enable nano highlight
if ! ls ~/.nanorc; then
    cd /tmp
    git clone https://github.com/scopatz/nanorc.git
    cd /tmp/nanorc && sh install.sh
fi

# Install CompileDaemon
if ! command -v CompileDaemon; then
    cd /tmp
    git clone -b v1.4.0 https://github.com/githubnemo/CompileDaemon.git
    cd CompileDaemon
    go build -o CompileDaemon .
    chmod +x CompileDaemon
    mv CompileDaemon /usr/local/bin
fi

# Hot reload
cd /app
CompileDaemon -build="go build -o clean-arch-lite main.go" -command="./clean-arch-lite"