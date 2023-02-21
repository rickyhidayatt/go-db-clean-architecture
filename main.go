package main

import "github.com/rickyhidayatt/delivery"

func main() {
	delivery.Run()

	// r := gin.Default() // menampilkan dengan logger
	// // r := gin.New() // tidak ada midleware secara defaultnya

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Hai Ricky Ganteng",
	// 	})
	// })
	// err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// if err != nil {
	// 	panic(err)
	// }
}

// fungtion handler function yang memiliki parameter cintex

// func greeting(ctx *gin.Context){
// 	name := ctx.Param("name")
// 	ctx.String(200, "Hallo")

// }
