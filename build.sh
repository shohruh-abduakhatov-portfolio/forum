#!/bin/bash

go install
go build --tags=sqlite_userauth
sudo ./forum.com