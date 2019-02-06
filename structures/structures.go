package structures

//Creando estructuras

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

type Response struct {
	Status string `json:"status"`
	Data User `json:"data"`
	Message string `json:"message"`
}