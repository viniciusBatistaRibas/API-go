package main

import (
	"fmt"

	"github.com/viniciusBatistaRibas/go-rest-api.git/database"
	"github.com/viniciusBatistaRibas/go-rest-api.git/routes"
)

func main() {
	database.Conceta_BD()
	fmt.Println("iniciando o server rest")
	routes.HandleRequest()
}
