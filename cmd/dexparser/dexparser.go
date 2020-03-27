package main

import (
	"fmt"
	"github.com/smm-goddess/dexparser/dex/items"
	"io/ioutil"
	"log"
)

const (
	dexPath = "/home/neal/work/apks/classes.dex"
)

func main() {
	dexBytes, err := ioutil.ReadFile(dexPath)
	if err != nil {
		log.Fatal("error read dex file")
	}
	header := items.ParseHeader(dexBytes)
	stringIds := items.ParseStringIds(dexBytes, header.StringIdsOff, header.StringIdsSize)
	fmt.Printf("%x", dexBytes[stringIds[0].StringDataOff:stringIds[0].StringDataOff+4])
}
