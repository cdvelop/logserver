package logserver

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/cdvelop/model"
	"github.com/cdvelop/output"
)

func (l *logServer) AddHandlerToRegisterLogsInDB(h *model.MainHandler) {
	l.Log("AddHandlerToRegisterLogsInDB no implementado en logserver")
}

func (l logServer) Log(message ...interface{}) interface{} {

	if l.dev_mode {
		var other bool
		var out_string string
		var space string

		for _, msg := range message {
			switch msg := msg.(type) {
			case string:
				out_string += space + msg
			case int:
				out_string += space + strconv.Itoa(msg)

			case int64:
				out_string += space + strconv.FormatInt(msg, 10)

			case bool:
				out_string += space + strconv.FormatBool(msg)

			default:
				other = true
			}
			space = " "
		}

		if other {
			fmt.Println(message...)
		} else {
			output.PrintInfo(out_string)
		}
	} else {
		log.Println(message...)
	}

	return nil
}

func (l logServer) Write(p []byte) (n int, err error) {

	f, err := os.OpenFile(l.log_file_name, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		e := "Archivo " + l.log_file_name + " no existe"
		fmt.Println(e, err)
		return 0, fmt.Errorf(e, err)
	}

	if len(p) != 0 {
		_, err = f.Write(p)
		if err != nil {
			e := "no se puede escribir en el archivo " + l.log_file_name
			fmt.Println(e, err)
			return 0, fmt.Errorf(e, err)
		}
	}

	return 0, nil
}
