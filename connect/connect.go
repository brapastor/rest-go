package connect

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/brapastor/rest/structures"
)

var connection *gorm.DB

// Variables de entorno
const engine_sql string = "mysql"
const username = "root"
const password string = "root"
const database  string = "rest_go"

func InitializeDatabase()  {
	connection = ConnectORM(CreateString())
	log.Println("LA CONEXCION CON LA BD FUE EXITOSA")
}

func CloseConnection()  {
	connection.Close()
	log.Println("LA CONEXCION CON LA BD FUE CERRADA")

}

func ConnectORM(stringConnection string) *gorm.DB {
	connection, err := gorm.Open(engine_sql, stringConnection)
	if err != nil{
		log.Fatal(err)
		return nil
	}
	return connection
}

func GerUser(id string)  structures.User {
	user := structures.User{}
	connection.Where("id = ?", id).First(&user)
	return user
}
func CreateUser(user structures.User) structures.User {
	connection.Create(&user) //Se asigar un id
	return user
}

func UpdateUser(id string, user structures.User)  structures.User{
	currectUser := structures.User{}
	connection.Where("id = ?", id).First(&currectUser)
	currectUser.Username  = user.Username
	currectUser.FirstName  = user.FirstName
	currectUser.LastName  = user.LastName
	connection.Save(&currectUser)

	return currectUser
}

func DeleteUser(id string)  {
	user := structures.User{}
	connection.Where("id = ?", id).First(&user)
	connection.Delete(&user)
}
func CreateString()  string{
	return username + ":" + password + "@/"+ database
}