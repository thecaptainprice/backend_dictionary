package utils

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
)

func Logger(w http.ResponseWriter, err error, message string) {
	_, file, line, _ := runtime.Caller(1)
	http.Error(w, message, http.StatusInternalServerError)
	fmt.Println(err.Error() + " `File `:" + file + " `line number` : " + strconv.Itoa(line))
}
