package main

import (
	"bookstore/internals/routes"
	"bookstore/pkg"
	"log"
)

// Depedency Injection (DI)

func main() {
	//  Inisialisasi DB
	db, err := pkg.InitMySql()
	if err != nil {
		log.Fatal(err)
		// return
	}
	//  Inisialisasi Router
	router := routes.InitRouter(db)
	//  Inisialisasi Server
	server := pkg.InitServer(router)
	// Jalankan servernya
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
