package domain

// Representación de un usuario, muy básico, sin atributos de auditoría, estado de activación, ni mail, etc.
type User struct {
	Id        string
	FirstName string
	LastName  string
	Username  string
	Password  string
}
