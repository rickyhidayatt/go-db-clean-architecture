package delivery

import (
	"database/sql"
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
	// pilihanMenu()
	// var pilihan int

	// // for {
	// // 	fmt.Scan(&pilihan)
	// // 	if pilihan == 5 {
	// // 		fmt.Println("Terimkasih sudah menggunakan aplikasi")
	// // 		break
	// // 	}
	// // 	switch pilihan {
	// // 	case 1:
	// // 		addCustomer()
	// // 		pilihanMenu()

	// // 	case 2:
	// // 		getDataCustomer()
	// // 	case 3:
	// // 		updatecust()
	// // 		pilihanMenu()
	// // 	case 4:
	// // 		deleteCust()
	// // 		pilihanMenu()
	// // 	}
	// // }

	// StoreCLI(Db)
	// JoinProdukStoreCLI(Db)
	getDataCustomer()

	// ProductCLI(Db)
}

func StoreCLI(db *sql.DB) {
	storeRepo := repository.NewStoreRepository(db)
	storeUsecase := usecase.NewStoreUseCase(storeRepo)

	stores, err := storeUsecase.GetAllStore()
	if err != nil {
		panic(err.Error())
	}

	for _, store := range stores {
		fmt.Println(store)
	}

	// tinggal rapihkan di terminal nya
}

func JoinProdukStoreCLI(db *sql.DB) {

	joinRepo := repository.JoinStoreProdukRepository(db)
	joinUseCase := usecase.JoinProductUseCase(joinRepo)

	stores, products, err := joinUseCase.JoinProductAndStore()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Data Store:")
	for _, store := range stores {
		fmt.Printf("ID: %s, Siup Number: %s, Nama: %s\n", store.Id, store.SiupNumber, store.Name)
	}

	fmt.Println("Data Product:")
	for _, product := range products {
		fmt.Printf("ID: %s, Nama: %s, Stock: %d, Harga: %d, Created At: %s, Produk_ID : %s\n", product.Id, product.Name, product.Stock, product.Price, product.CreatedAt, product.ProductId)
	}

	// for _, store := range stores {
	// 	fmt.Printf("Data Store:\nID: %s, Siup Number: %s, Nama: %s\n", store.Id, store.SiupNumber, store.Name)
	// 	fmt.Println("Data Produk:")

	// 	// Loop semua produk dari store saat ini
	// 	for _, product := range products {
	// 		// Cetak produk jika id store nya sama dengan id store saat ini
	// 		if product.ProductId == store.Id {
	// 			fmt.Printf("ID: %s, Nama: %s, Stock: %d, Harga: %d, Created At: %s\n", product.Id, product.Name, product.Stock, product.Price, product.CreatedAt)
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// tinggal rapihkan di terminal nya
}

func ProductCLI(db *sql.DB) {
	storeRepo := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUseCase(storeRepo)

	product, err := productUsecase.GetAllProduct()
	if err != nil {
		panic(err.Error())
	}

	for _, v := range product {
		fmt.Println("ID : ", v.Id)
		fmt.Println("Name : ", v.Name)
		fmt.Println("Stock : ", v.Stock)
		fmt.Println("Price : ", v.Price)
		fmt.Println("CreatedAt : ", v.CreatedAt)
		fmt.Println("CreatedAt : ", v.ProductId)

		fmt.Println()
	}

}
