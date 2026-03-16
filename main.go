package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Fprintln(os.Stderr, "usage: hashpwd <password>")
		os.Exit(1)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(os.Args[1]), bcrypt.MinCost)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	parts := strings.SplitN(string(hash), "$", 4)
	result := strings.NewReplacer(".", "", "/", "").Replace(parts[3])
	fmt.Println(result)
}
