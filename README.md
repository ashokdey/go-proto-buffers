# Using Protobuff 3 in Golang

This repo contains code manipulation of prot bufers v3 using golang

## Quick Notes

If you want to compile proto files to **golang** and you're using **zhs** instead of **bash** in Ubuntu/MacOS, add the following lines in you `.zshrc`. Follow the steps:

- `vim ~/.zhsrc`
- Add the following:
  - `export GOPATH=$HOME/go`
  - `export GOBIN=$GOPATH/bin`
  - `export PATH=$PATH:$GOPATH/bin`
- Now you can safely execute `./genarate.sh` to convert the `proto` file `golang code`
