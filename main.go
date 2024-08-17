package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bradenhilton/mozillainstallhash"
)

func main() {
	if len(os.Args) > 2 {
		log.Fatalf("too many args - got: %d, want: 1", len(os.Args)-1)
	}

	path := os.Args[1]
	path = strings.TrimSuffix(path, "/")
	path = strings.TrimSuffix(path, "\\")
	
	hash, err := mozillainstallhash.MozillaInstallHash(path)
	if err != nil {
		log.Fatalf("hashing: %s", path)
	}

	fmt.Print(hash)
}
