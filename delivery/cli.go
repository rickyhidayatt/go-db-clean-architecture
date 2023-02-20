package delivery

import (
	"fmt"

	"github.com/rickyhidayatt/config"
	"github.com/rickyhidayatt/model"
	"github.com/rickyhidayatt/repository"
	"github.com/rickyhidayatt/usecase"
)

// jembatan dari beberapa package buat di tampilin di CLI
func Run() {

	config := config.NewConfig()                         // ngambil koneksi dari package config
	db := config.DbConnect()                             // konekin ke DB dari config
	customerRepo := repository.NewCustomerRepository(db) // panggil repository insert data
	customerUseCase := usecase.NewCustomerUseCase(customerRepo)

	// ==========================================
	//insert data
	// var customer = model.Customer{
	// 	Name:    "Jhon Lenon",
	// 	Balance: 700000,
	// }
	// err := customerUseCase.RegisterCustomer(&customer)

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Success Create New Customer")
	// }
	// ==============================================

	// LIHAT SEMUA DATA CUSTOMER ======================
	// customer, err := customerUseCase.GetAllCustomer()
	// if err != nil {
	// 	panic(err)
	// }
	// for _, v := range customer {
	// 	fmt.Println(v)
	// }
	// =================================================

	// // LIhat data By ID ++++++++++++++++++++++
	// customerCheckByID, err := customerUseCase.GetCustomerByID(1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customerCheckByID)

	// // +++++++++++++++++++++++++

	//Update
	// var customer = model.Customer{
	// 	Id:      "2",
	// 	Name:    "Jhon Update",
	// 	Balance: 700000,
	// }
	// err := customerUseCase.UpdateCustomer(customer.Id, &customer)

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Success UPDATE  Customer")
	// }

	// ==============================

	// Delete

	var customer = model.Customer{
		Id:      "2",
		Name:    "Jhon Update",
		Balance: 700000,
	}
	err := customerUseCase.DeteCustomerByID(customer.Id, &customer)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Success Delete  Customer")
	}
}
