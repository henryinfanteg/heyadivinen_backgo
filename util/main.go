package main

import (
	"fmt"

	random "github.com/henryinfanteg/heyadivinen_backgo/util/random"
	security "github.com/henryinfanteg/heyadivinen_backgo/util/security"
)

func main() {
	fmt.Println("Hello util")
	fmt.Println(random.GenerateRandomString(1, 20))
	fmt.Println(security.EncryptHMACSHA512("123456", "XXXXX"))
}
