#!/bin/bash
cd src
pwd
go build main.go
mv main ../build
scp /home/matheus/Documents/Dev/Go/DNSFilter/build/main matheus@10.0.1.189:/home/matheus
