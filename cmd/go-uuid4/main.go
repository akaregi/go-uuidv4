package main

import (
	"fmt"
	"github.com/akarregi/go-uuidv4/pkg/uuid"
)

func main() {
	id, err := uuid.NewUUID()

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(uuid.Stringify(id))
}
