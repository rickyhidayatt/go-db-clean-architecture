package delivery

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rickyhidayatt/config"
	"github.com/rickyhidayatt/repository"
	"github.com/rickyhidayatt/usecase"
)

var Config = config.NewConfig()                         // ngambil koneksi dari package config
var Db = Config.DbConnect()                             // konekin ke DB dari config
var CustomerRepo = repository.NewCustomerRepository(Db) // panggil repository insert data
var CustomerUseCase = usecase.NewCustomerUseCase(CustomerRepo)

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
	CustomerCli(Db)
	// JoinProdukStoreCLI(Db)
	// getDataCustomer()

	// ProductCLI(Db)
}

func StoreCLI(db *sqlx.DB) {
	productRepo := repository.NewProductRepository(db)
	productUc := usecase.NewProductUseCase(productRepo)

	storeRepo := repository.NewStoreRepository(db)
	storeUseCase := usecase.NewStoreUseCase(storeRepo, productUc)

	// GET ALL STORE
	storeProducts, err := storeUseCase.GetAllStoreWithProducts()
	if err != nil {
		panic(err.Error())
	}

	for _, product := range storeProducts {
		fmt.Println(product)
	}

	router := gin.Default()
	router.GET("/customers", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data KEY": storeProducts,
		})
	})

	// jalankan gin
	errGin := router.Run()
	if err != nil {
		panic(errGin)
	}
}

// uji coba pake join
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
		fmt.Printf("ID: %s, Nama: %s, Stock: %d, Harga: %d, Created At: %s, Produk_ID : %s\n", product.Id, product.Name, product.Stock, product.Price, product.CreatedAt, product.StoreId)
	}
}

// ==============

func ProductCLI(db *sqlx.DB) {
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
		fmt.Println("CreatedAt : ", v.StoreId)

		fmt.Println()
	}

}

func CustomerCli(db *sqlx.DB) {
	customerRepo := repository.NewCustomerRepository(db)
	customerUseCase := usecase.NewCustomerUseCase(customerRepo)

	//GET LIST CUSTOMER
	customers, err := customerUseCase.GetAllCustomer()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/customers", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": customers,
		})
	})
	errGin := router.Run()
	if err != nil {
		panic(errGin)
	}

	// percobaan nyari customers by id
	router.GET("/customers/:id", func(ctx *gin.Context) {
		customerID := ctx.Param("id")
		if err != nil {
			panic(err)
		}

		customer, err := customerUseCase.GetCustomerByID(customerID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Customer not found",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"data": customer,
		})
	})

	err = router.Run()
	if err != nil {
		panic(err)
	}
	// for _, customer := range customers {
	// 	fmt.Println(strings.Repeat("=", 20))
	// 	fmt.Println("ID :", customer.Id)
	// 	fmt.Println("NAME :", customer.Name)
	// 	fmt.Println("BALANCE :", customer.Balance)
	// }
}
