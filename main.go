package main

import (
	"fmt"

	_ "github.com/go-zookeeper/zk"
)

func ExtFunction() {
	fmt.Println("Printing from package")
}
