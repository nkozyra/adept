//	Adept binary wrapper
package main

import (
	"fmt"

	"github.com/nkozyra/adept/src"
)

const (
	configLocation = "config/config.json"
)

//	Loads config
func init() {

}

//	Starts the magic
func main() {
	fmt.Println("Adept server starting")
	adept.LoadConfig(configLocation)
}
