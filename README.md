# gotools


## Compiling from Windows to Linux

https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04

### Powershell

```
$env:GOOS = "linux"
$env:GOARCH = "amd64"
```


```
$env:GOOS = "linux"
$env:GOARCH = "arm"
$env:GOARM = "5"
```