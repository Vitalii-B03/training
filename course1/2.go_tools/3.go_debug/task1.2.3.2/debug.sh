#!/bin/bash
echo "Debug started..."
dlv debug main.go
if [ -z "$1" ]; then
  echo "Usage: debug.sh main.go"
    exit 1
    fi

go build -o myprogram "$1"
dlv exec ./myprogram
echo "Debug ended."