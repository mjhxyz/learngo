package main

import (
	"go.uber.org/zap"
	"learngo/advance/errorhanding/filelistingserver/filelisting"
	"net/http"
	_ "net/http/pprof"
	"os"
)

var logger, _ = zap.NewProduction()

type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 定义能给用户看的 Error
type userError interface {
	error
	Message() string
}

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				logger.Sugar().Warn("Panic: ", r)
				http.Error(
					writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if userErr, ok := err.(userError); ok {
			http.Error(writer, userErr.Message(), http.StatusBadRequest)
			return
		}
		if err != nil {
			logger.Warn("Error handling request: " + err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(
				writer,
				http.StatusText(code),
				code)
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
