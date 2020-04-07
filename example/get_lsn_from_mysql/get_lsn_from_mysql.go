package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func getLSN(db *gorm.DB) {
	var t string
	var name string
	var status string

	row := db.Raw("SHOW ENGINE INNODB STATUS;").Row()
	row.Scan(&t, &name, &status)
	showText := false
	lineNums := 0
	maxNums := 11
	scanner := bufio.NewScanner(strings.NewReader(status))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if scanner.Text() == "LOG" {
			fmt.Println("---")
			showText = true
		}
		if showText && lineNums <= maxNums {
			fmt.Println(scanner.Text())
			lineNums++
		}
	}

	fmt.Println()
}

// User Model
type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open("mysql", "root:@(127.0.0.1)/?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	getLSN(db)

	db.Exec("CREATE DATABASE TestLSN;")
	fmt.Println("LSN after create database:")
	getLSN(db)

	db.Exec("USE TestLSN;")
	db.CreateTable(&User{})
	fmt.Println("LSN after create table:")
	getLSN(db)

	db.Exec("ALTER TABLE users ADD gender int NULL AFTER age;")
	fmt.Println("LSN after add column:")
	getLSN(db)

	user := User{Name: "Jinzhu", Age: 18}
	db.Create(&user)
	fmt.Println("LSN after add row:")
	getLSN(db)

	user.Name = "jinzhu 2"
	user.Age = 100
	db.Save(&user)
	fmt.Println("LSN after update row:")
	getLSN(db)

	//Soft Delete
	db.Delete(&user)
	fmt.Println("LSN after delete softly:")
	getLSN(db)

	// Delete record permanently
	db.Unscoped().Delete(&user)
	fmt.Println("LSN after delete permanently:")
	getLSN(db)

	db.DropTable(&User{})
	fmt.Println("LSN after drop table:")
	getLSN(db)

	db.Exec("DROP DATABASE TestLSN;")
	fmt.Println("LSN after drop database:")
	getLSN(db)
}
