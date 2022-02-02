package sysdbdriver

import (
	"database/sql"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/mattn/go-oci8"
	"os"
	"strconv"
)

func OracleTestWithOCI8() {
	//Reading the Database Environment variables.
	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("192.168.1.10/pdev")
	db, err := sql.Open("oci8", username+"/"+password+"@"+database)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Printf("Error connecting to the database: %s\n", err)
		return
	}

	rows, err := db.Query("select 2+2 from dual")
	if err != nil {
		fmt.Println("Error fetching addition")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var a string = "Hello!"
	for rows.Next() {
		var sum int
		rows.Scan(&sum)
		result := strconv.Itoa(sum)
		a = "2 + 2 always equals: " + result
		fmt.Println(a)
	}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": a})
	}))

}
