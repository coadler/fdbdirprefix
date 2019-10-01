package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
)

func main() {
	fdb.MustAPIVersion(610)
	db := fdb.MustOpenDefault()

	dir, err := directory.CreateOrOpen(db, []string{os.Args[1]}, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(hex.EncodeToString(dir.Bytes()))
}
