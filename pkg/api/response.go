package api

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
	"sync"
)

var zwPool = sync.Pool{
	New: func() any {
		return gzip.NewWriter(nil)
	},
}

func WriteResponse(w http.ResponseWriter, code int, v interface{}) error {
	zw := zwPool.Get().(*gzip.Writer)
	defer zwPool.Put(zw)
	defer zw.Close()
	zw.Reset(w)

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Content-Encoding", "gzip")
	w.WriteHeader(code)
	return json.NewEncoder(zw).Encode(v)
}
