package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"go.coder.com/flog"
)

var (
	apiV   int
	create bool
)

func init() {
	flag.IntVar(&apiV, "apiVersion", 610, "FDB API version to use")
	flag.BoolVar(&create, "create", false, "Create directory if it doesn't exist")
}

func main() {
	fdb.MustAPIVersion(apiV)
	db := fdb.MustOpenDefault()

	if len(os.Args) != 2 {
		flog.Error("Please specify one directory")
		os.Exit(1)
	}

	dir, err := directory.Open(db, []string{os.Args[1]}, nil)
	if err != nil {
		flog.Fatal("Failed to open directory: " + err.Error())
	}

	fmt.Println(hex.EncodeToString(dir.Bytes()))
}
