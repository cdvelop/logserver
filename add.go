package logserver

import "log"

// log_file ej: "logs.log" default: app.log
func Add(log_file ...string) *logServer {

	var file_name = "app.log"

	for _, v := range log_file {
		file_name = v
	}

	logger := &logServer{
		log_file_name: file_name,
	}

	log.SetOutput(logger)

	return logger
}
