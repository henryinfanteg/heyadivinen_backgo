package main

import (
	"fmt"

	random "gitlab.com/adivinagame/backend/maxadivinabackend/util/random"
	security "gitlab.com/adivinagame/backend/maxadivinabackend/util/security"
)

func main() {
	fmt.Println("Hello util")
	fmt.Println(random.GenerateRandomString(1, 20))
	fmt.Println(security.EncryptHMACSHA512("123456", "XXXXX"))
}
