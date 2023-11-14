package logserver

import (
	"fmt"
	"log"
	"os"

	"github.com/cdvelop/model"
)

func (logServer) Log(message ...interface{}) interface{} {
	log.Println(message...)
	return nil
}

func (l logServer) Write(p []byte) (n int, err error) {

	f, err := os.OpenFile(l.log_file_name, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		e := "Archivo " + l.log_file_name + " no existe"
		fmt.Println(e, err)
		return 0, model.Error(e, err)
	}

	if len(p) != 0 {
		_, err = f.Write(p)
		if err != nil {
			e := "no se puede escribir en el archivo " + l.log_file_name
			fmt.Println(e, err)
			return 0, model.Error(e, err)
		}
	}

	return 0, nil
}
