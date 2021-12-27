# README

The is a repo of go-present slides for various topics I have given.

| Title | Summary | Presented Date | Presentation Location | Notes |
|--|--|--|--|--|
| Introduction to Generic Programming in Go | A motivation for what generics is and why we might want them in Go. Followed by a summary of the proposed generics feature set as of April 2021 with examples | May 2021 | Comcast Internal Go Club | The examples in these slides may not be runnable they were using an unofficial syntax |


### Running Using Go present

These slides are intended to by viewed using the go tool `go present`

> There is a viewable PDF in the respective directories if that is better for you. Note The format is inferior

To view the presentations:

Fetch [go-present](https://pkg.go.dev/golang.org/x/tools/present)
```
go get -u golang.org/x/tools/cmd/present
```

run go present in this directory:
```
present
```

A webpage will be hosted at [http://127.0.0.1:3999](http://127.0.0.1:3999)
