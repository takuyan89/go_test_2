// package main

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// var Db *sql.DB

// var err error

// type Person struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	Db, err = sql.Open("postgres", "user=test_user dbname=testdb password=password sslmode=disable")
// 	if err != nil {
// 		log.Panicln(err)
// 	}
// 	defer Db.Close()

// 	// cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
// 	// _, err = Db.Exec(cmd, "Nancy", 20)
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }

// 	// cmd := "SELECT * FROM persons where age = $1"
// 	// row := Db.QueryRow(cmd, 20)
// 	// var p Person
// 	// err = row.Scan(&p.Name, &p.Age)
// 	// if err != nil {
// 	// 	if err == sql.ErrNoRows {
// 	// 		log.Println("No row")
// 	// 	} else {
// 	// 		log.Println(err)
// 	// 	}
// 	// }
// 	// fmt.Println(p.Name, p.Age)

// 	// cmd = "SELECT * FROM persons"
// 	// rows, _ := Db.Query(cmd)
// 	// defer rows.Close()
// 	// var pp []Person
// 	// for rows.Next() {
// 	// 	var p Person
// 	// 	err := rows.Scan(&p.Name, &p.Age)
// 	// 	if err != nil {
// 	// 		log.Println(err)
// 	// 	}
// 	// 	pp = append(pp, p)
// 	// }
// 	// err = rows.Err()
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }
// 	// for _, p := range pp {
// 	// 	fmt.Println(p.Name, p.Age)
// 	// }

// 	// cmd := "UPDATE persons SET age = $1 WHERE name = $2"
// 	// _, err = Db.Exec(cmd, 25, "Nancy")
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }

// 	// cmd := "DELETE FROM persons WHERE name = $1"
// 	// _, err = Db.Exec(cmd, "Nancy")
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }
// }

package main

import (
	"fmt"

	"github.com/google/uuid"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port int
	Dbname string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustInt(8080),
		Dbname: cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("SQLDriver").String(),
	}
}

func main() {	
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.Dbname, Config.Dbname)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)

	uuidObj, _ := uuid.NewUUID()
	fmt.Println("  ", uuidObj.String())

	uuidObj2, _ := uuid.NewRandom()
	fmt.Println("  ", uuidObj2.String())
}