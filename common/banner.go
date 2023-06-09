package common

import (
	"fmt"
	"log"
	"os"
)

func ReadBanner(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(file))
}
