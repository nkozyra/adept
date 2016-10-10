//	Adept binary wrapper
package main

import (
	"fmt"

	"github.com/nkozyra/adept/src"
)

const (
	configLocation = "config/config.json"
)

func main() {
	var x adept.Question
	fmt.Println(x)

	conf := adept.LoadConfig(configLocation)
	fmt.Println(conf)
}
