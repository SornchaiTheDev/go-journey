package main

import (
	"fmt"

	"github.com/SornchaiTheDev/username_password/models"
)

func main() {
	hachi := models.Tamagotchi{Name : "Hachi"}

	fmt.Println(hachi.Name)
}
