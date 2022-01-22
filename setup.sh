#!/bin/bash
rm -rf .git README.md
go mod init
git init
rm "$0"
