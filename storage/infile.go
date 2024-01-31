package storage

import (
	"code/app"
	"encoding/json"
	"fmt"
	"os"
)

type MyFile os.File

const DefualtPath = "./storage.txt"

func (f MyFile) CreateUser(u app.User) {
	file, err := os.OpenFile(DefualtPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Printf("error happend when open the file :%v \n", err)

		return
	}
	defer file.Close()

	data, err := json.Marshal(u)
	if err != nil {
		fmt.Printf("error happend when serializing :%v \n", err)

		return
	}

	if _, err := file.Write(data); err != nil {
		fmt.Printf("error happend when write the file :%v \n", err)

		return
	}

}

func (f MyFile) GetUser(name string) app.User {
	// work just one user
	file, err := os.ReadFile(DefualtPath)
	if err != nil {
		fmt.Printf("error happend when read the file :%v \n", err)

		return app.User{}
	}
	myUser := app.User{}
	err = json.Unmarshal(file, &myUser)
	if err != nil {
		fmt.Printf("error unmurshal %v \n", err)
	}
	fmt.Println("user : ", myUser)

	return app.User{}
}
