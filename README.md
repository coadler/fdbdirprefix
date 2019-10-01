# fdbdirprefix

Get the hex prefix of a directory in FoundationDB.

## Installation

Ensure `$GOPATH/bin` is in your `$PATH`

```
GO111MODULE=off go get -u github.com/coadler/fdbdirprefix
cd $GOPATH/src/github.com/coadler/fdbdirprefix
go install
```

Alternatively, you can use [gobin](https://github.com/myitcv/gobin).

```
gobin github.com/coadler/foundationdb
```

## Usage:

`fdbdirprefix [dirName]`


## Flags:

`--apiVersion int` - What FDB API version to use. Default: 610

`--create` - Whether or not to create the directory if it doesn't exist. Default: false
