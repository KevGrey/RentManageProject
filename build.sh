#!/bin/bash

echo "请选择要编译的系统环境:"
echo "1. Windows_amd64"
echo "2. linux_amd64"

read -p "请输入数字: " action

if [ "$action" == "1" ]; then
    echo "编译Windows版本64位"
    export CGO_ENABLED=0
    export GOOS=windows
    export GOARCH=amd64
    go build -o project_user/target/project-user.exe project_user/main.go
    go build -o project_api/target/project-api.exe project_api/main.go
elif [ "$action" == "2" ]; then
    echo "编译Linux版本64位"
    export CGO_ENABLED=0
    export GOOS=linux
    export GOARCH=amd64
    go build -o project_user/target/project-user project_user/main.go
    go build -o project_api/target/project-api project_api/main.go
    go build -o project-project/target/project-project project-project/main.go
else
    echo "无效的选项"
fi
