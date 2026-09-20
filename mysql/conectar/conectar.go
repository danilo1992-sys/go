package conectar 

import (
	"database/sql"
	_"github.com/go-sql-driver/myql"
	"github.com/joho/godotenv"
	"os"
)

var Db *sql.DB 

func Conectart()  {
	errorVariable := godotenv.Load()
	if errorVariable != nil {
		panic(errorVariable)
	}
	conection, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv(DB_PORT)+")/"os.Getenv("DB_NAME"))
	if err !=nill {
		panic(err)
	}
	Db = conection
}

func CerrarConexion() {
	Db.Close()
}