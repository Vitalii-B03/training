#!/bin/bash
echo "Compiling started..."
go build main.go
echo "Compiling complete."
go run main.go
echo "Program exited."