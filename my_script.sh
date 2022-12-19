#!/bin/sh
echo "runnig script"

read folderName number

go run golang/script.go $folderName $number

read -p "do you want to open vs code (y/n)" openVsCode

if [ $openVsCode = "y" ]; then
    echo "opening vscode.."
    cd $folderName
    code .
fi
