package main

import (
	"fmt"

	"github.com/Eliwelton-The-Espada/fcutils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
