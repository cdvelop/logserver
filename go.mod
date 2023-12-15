module github.com/cdvelop/logserver

go 1.20

require github.com/cdvelop/gotools v0.0.70

require github.com/cdvelop/strings v0.0.9 // indirect

require (
	github.com/cdvelop/input v0.0.67 // indirect
	github.com/cdvelop/model v0.0.95 // indirect
	github.com/cdvelop/output v0.0.16
	github.com/cdvelop/timetools v0.0.32 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/gotools => ../gotools

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/output => ../output
