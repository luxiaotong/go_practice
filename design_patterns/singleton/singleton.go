package main

import (
	"fmt"
	"sync"
)

type single struct{}

var instance *single

var once sync.Once

func getInstance() *single {
	if instance == nil {
		once.Do(func() {
			instance = &single{}
			fmt.Println("Creating single instance now.")
		})
	} else {
		fmt.Println("Single instance already created.")
	}
	return instance
}
