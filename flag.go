package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	token   string
	admin   int64
	promote bool
	version bool
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	flag.StringVar(&token, "token", "", "token of telegram bot")
	flag.Int64Var(&admin, "admin", 0, "admin of bot")
	flag.BoolVar(&promote, "promote", false, "promote option")
	flag.BoolVar(&version, "version", false, "version of regulatebot")
	flag.Parse()
}

func printVersion() {
	if version {
		fmt.Println(Version())
		os.Exit(0)
	}
}

func checkFlag() {
	if len(token) == 0 || admin == 0 {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
}
