package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/dict"
	"log"
	"os"
)

var (
	ls   = flag.Bool("l", false, "list dictionaries")
	dn   = flag.String("d", "*", "specify dictionary")
	addr = flag.String("s", "dict.org:2628", "specify dict server")
)

func main() {
	flag.Parse()
	client, err := dict.Dial("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	if *ls {
		dicts, err := client.Dicts()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		for _, dict := range dicts {
			fmt.Println(dict.Name + ":" + dict.Desc)
		}
		os.Exit(0)
	}
	for _, arg := range flag.Args() {
		defs, err := client.Define(*dn, arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		for _, def := range defs {
			fmt.Println(string(def.Text))
		}
	}
}
