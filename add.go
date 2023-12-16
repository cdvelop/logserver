package logserver

import (
	"log"
	"os"
)

// log_file ej: "logs.log" default: app.log
func AddLoggerAdapter(log_file ...string) *logServer {

	var file_name = "app.log"

	for _, v := range log_file {
		file_name = v
	}

	l := &logServer{
		log_file_name: file_name,
	}

	for _, arg := range os.Args {
		if arg == "dev" {
			l.dev_mode = true
		}
	}

	log.SetOutput(l)

	return l
}
