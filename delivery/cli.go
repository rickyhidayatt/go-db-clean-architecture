package delivery

import (
	"fmt"

	"github.com/rickyhidayatt/config"
	"github.com/rickyhidayatt/repository"
	"github.com/rickyhidayatt/usecase"
)

var Config = config.NewConfig()                         // ngambil koneksi dari package config
var Db = Config.DbConnect()                             // konekin ke DB dari config
var CustomerRepo = repository.NewCustomerRepository(Db) // panggil repository insert data
var CustomerUseCase = usecase.NewCustomerUseCase(CustomerRepo)

// jembatan dari beberapa package buat di tampilin di CLI
func Run() {
	pilihanMenu()
	var pilihan int

	for {
		fmt.Scan(&pilihan)
		if pilihan == 5 {
			fmt.Println("Terimkasih sudah menggunakan aplikasi")
			break
		}
		switch pilihan {
		case 1:
			addCustomer()
			pilihanMenu()

		case 2:
			getDataCustomer()
		case 3:
			updatecust()
			pilihanMenu()
		case 4:
			deleteCust()
			pilihanMenu()
		}
	}
}
