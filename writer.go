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

	f, err := os.OpenFile(l.log_file_name+".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		e := "Archivo server.log no existe"
		fmt.Println(e, err)
		return 0, model.Error(e, err)
	}

	_, err = f.Write(p)
	if err != nil {
		e := "no se puede escribir en el archivo server.log"
		fmt.Println(e, err)
		return 0, model.Error(e, err)
	}

	return 0, nil
}
