package storage

import "code/app"

type Memory struct {
	Users map[string]app.User
}

func (m *Memory) CreateUser(newUser app.User) {
	m.Users[newUser.Name] = newUser
}
func (m *Memory) GetUser(name string) app.User {
	return m.Users[name]
}
