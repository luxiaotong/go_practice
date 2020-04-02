package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

//API int
type API int

//Item struct
type Item struct {
	Title string
	Body  string
}

var database []Item

//GetDB method
func (api *API) GetDB(title string, reply *[]Item) error {
	*reply = database
	return nil
}

//GetByName method
func (api *API) GetByName(title string, reply *Item) error {
	var item Item
	for _, val := range database {
		if val.Title == title {
			item = val
			break
		}
	}
	*reply = item
	return nil
}

//AddItem method
func (api *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

//EditItem method
func (api *API) EditItem(edit Item, reply *Item) error {
	var changed Item
	for idx, val := range database {
		if val.Title == edit.Title {
			database[idx] = edit
			changed = database[idx]
			break
		}
	}
	*reply = changed
	return nil
}

//DeleteItem method
func (api *API) DeleteItem(item Item, reply *Item) error {
	var del Item
	for idx, val := range database {
		if val.Title == item.Title {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}
	*reply = del
	return nil
}

func main() {
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering api:", err)
	}

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listener error:", err)
	}

	log.Printf("serving rpc on port %d\n", 4040)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving:", err)
	}

	// fmt.Println("Initial database:", database)

	// a := Item{"first", "first item"}
	// b := Item{"second", "second item"}
	// c := Item{"third", "third item"}
	// AddItem(a)
	// AddItem(b)
	// AddItem(c)
	// fmt.Println("Second database:", database)

	// DeleteItem(b)
	// fmt.Println("Third database:", database)

	// EditItem("third", Item{"fourth", "forth item"})
	// fmt.Println("Fourth database:", database)

	// x := GetByName("fourth")
	// y := GetByName("first")
	// fmt.Println(x, y)
}
