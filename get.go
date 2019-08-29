package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

type emp struct {
	tableName struct{} `sql:"emp"`
	ID        int      `sql:"user_id"`
	Password  string   `sql:"password"`
	Email     string   `sql:"email"`
	//HouseNo   int      `sql:"house_no"`
}

func connection() (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		User:     "root",
		Password: "tolexo",
		Database: "postgres",
		Addr:     "localhost:5432",
	})
	return
}
func empdetail() (e []emp, err error) {
	conn := connection()

	err = conn.Model(&e).Select()
	if err != nil {
		return
	}
	return
}
func goget(c *gin.Context) {
	data, err := empdetail()
	if err == nil {
		// v, err := json.Marshal(data)
		if err == nil {
			c.JSON(200, gin.H{
				"message": data,
			})
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", goget)
	r.Run(":5000") // listen and serve on 0.0.0.0:8080
}
