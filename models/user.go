package models

type User struct {
	Login      string
	Email      string
	Password   string
	Role       int
	Active     bool
	Created_at string
	Updated_at string
}

// Проверка пароля
func (user *User) CheckPasswod(password, encPassword string) bool {

	return true
}

// Авторизация
func (user *User) Authorize(login, password string) bool {

	return true
}

// Возвращает зашифрованный пароль
func (user *User) GetPassword(password string) string {

	return password
}
