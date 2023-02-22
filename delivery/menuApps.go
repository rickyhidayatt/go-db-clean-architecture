package delivery

import (
	"fmt"
	"strings"

	"github.com/rickyhidayatt/model"
)

func pilihanMenu() {
	menu := `

Selamat Datang Di TokoBagus.com 
Pilih menu :

1. Tambah Customer
2. Lihat Data Customer
3. Perbarui Data Customer
4. Hapus Data Customer
5. Exit
Masukan pilihan kamu :
	`
	fmt.Println(menu)

}

func addCustomer() {

	var nama string
	var balance int
	fmt.Print("Masukan Nama Customer :")
	fmt.Scan(&nama)
	fmt.Print("Masukan Jumlah Saldo :")
	fmt.Scan(&balance)

	var customer = model.Customer{
		Name:    nama,
		Balance: balance,
	}

	err := CustomerUseCase.RegisterCustomer(&customer)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Sukses Tambah Customer")
	}
	fmt.Println()
	fmt.Println("Tekan 5 untuk keluar")
}

func getDataCustomer() {

	var pilih int
	fmt.Println("Lihat data customer :\n1. Lihat Semua Data Customer\n2. Berdasarkan Id\n ")

	fmt.Scan(&pilih)

	if pilih == 1 {
		customer, err := CustomerUseCase.GetAllCustomer()
		if err != nil {
			panic(err)
		}
		fmt.Println(strings.Repeat("-", 60))
		fmt.Println("ID\tNama Customer\t             Balance")
		fmt.Println(strings.Repeat("-", 60))
		for _, v := range customer {
			fmt.Printf("%s\t%-10s\t%20d \n", v.Id, v.Name, v.Balance)
		}

	} else if pilih == 2 {
		var idcust string
		fmt.Print("Masukan ID Customer : ")
		fmt.Scan(&idcust)
		customerCheckByID, err := CustomerUseCase.GetCustomerByID(idcust)

		if err != nil {
			panic(err)
		}
		fmt.Println(customerCheckByID)
	} else {
		fmt.Println("Pilihan yang kamu masukan salah")
	}
	fmt.Println()
	fmt.Println("Tekan 5 untuk keluar")

}

func updatecust() {
	var nama, id string
	var balance int
	fmt.Print("Masukan ID Customer :")
	fmt.Scan(&id)
	fmt.Print("Update Nama Customer :")
	fmt.Scan(&nama)
	fmt.Print("Update Jumlah Saldo :")
	fmt.Scan(&balance)

	var customer = model.Customer{
		Id:      id,
		Name:    nama,
		Balance: balance,
	}

	err := CustomerUseCase.UpdateCustomer(customer.Id, &customer)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Sukses Update Data Customer")
	}
	fmt.Println()
	fmt.Println("Tekan 5 untuk keluar")
}

func deleteCust() {

	var id string
	fmt.Print("Masukan ID Customer :")
	fmt.Scan(&id)
	var customer = model.Customer{
		Id: id,
	}
	err := CustomerUseCase.DeteCustomerByID(customer.Id, &customer)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Sukses Hapus Customer")
	}
	fmt.Println()
	fmt.Println("Tekan 5 untuk keluar")
}
