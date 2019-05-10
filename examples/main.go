package main

import (
	"fmt"
	"os"

	foldercompare "github.com/codetraceio/folder-compare"
)

func main() {
	var fSnap1, err1 = foldercompare.CreateSnapshot(os.Args[1])
	if err1 != nil {
		panic(err1)
	}
	var fSnap2, err2 = foldercompare.CreateSnapshot(os.Args[2])
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(foldercompare.CompareSnapshots(*fSnap1, *fSnap2))
}
