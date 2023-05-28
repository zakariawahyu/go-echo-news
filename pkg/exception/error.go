package exception

import (
	"errors"
	"net/http"
)

var ErrDocumentNotFound = errors.New("DocumentNotFound")

func NewErrorStatusCodeMaps() map[error]int {
	var errorStatusCodeMaps = make(map[error]int)
	errorStatusCodeMaps[ErrDocumentNotFound] = http.StatusNotFound
	return errorStatusCodeMaps
}

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}
