package handlers 

import (
	"mysql/conectar"
	"fmt"
)

func Listar()  {
	conectar.Conectar()
	sql := "select id,correo,telefono from clientes order by id desc;"
	data, err := conectar.Db.Query(sql)
	if err != nil {
		fmt.println(err)
	}
}