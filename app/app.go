package app

type UserStorage interface {
	CreateUser(u User)
	GetUser(name string) User
}
type App struct {
	Name string

	// insteed of these we use interface
	// storagePath string
	// memory      storage.Memory

	Storage UserStorage
}

type User struct {
	Name string
	Age  int
}

// func (a *App) GetStoragePath() string {
// 	return a.storagePath
// }
// func (a *App) SetStoragePath(newPath string) {
// 	a.storagePath = newPath
// }

func (a App) CreateUser(newUser User) {
	// file, err := os.OpenFile(a.storagePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	// if err != nil {
	// 	fmt.Printf("this error happend when open file : %v \n", err)

	// 	return
	// }

	// defer file.Close()

	// data, err := json.Marshal(newUser)
	// if err != nil {
	// 	fmt.Printf("this error happend when serializing : %v \n", err)

	// 	return
	// }
	// if _, err := file.Write(data); err != nil {
	// 	fmt.Printf("this error happend when writing file : %v \n", err)

	// 	return
	// }

	a.Storage.CreateUser(newUser)
}

func (a App) GetUser(name string) User {
	return a.Storage.GetUser(name)
}
