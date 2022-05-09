package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	key := make([]byte, 64)

// 	if _, erro := rand.Read(key); erro != nil {
// 		log.Fatal(erro)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(stringBase64)
// }

func main() {
	config.Load()
	r := router.Generate()

	fmt.Printf("Listening on port: %d", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
