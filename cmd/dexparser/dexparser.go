package main

import (
	"github.com/smm-goddess/dexparser/dex"
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
	dex.ParseDexFile(dexBytes)
}
