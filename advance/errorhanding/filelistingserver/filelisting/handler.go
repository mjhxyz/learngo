package filelisting

import (
	"io"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(
		request.URL.Path, prefix) != 0 {
		return userError("必须以 " + prefix + " 开头")
	}

	// 得到 list 后面的
	path := request.URL.Path[len("/list/"):] // /list/fib.txt
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
