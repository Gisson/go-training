package api

import (
	"encoding/json"
	"fmt"
	"github.com/Gisson/simple-messenger/message"
	"github.com/Gisson/simple-messenger/server"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func AddMessageHandler(srv *server.Server) (handle httprouter.Handle) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		var result Result
		var buf []byte
		var err error
		buf, err = ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := message.Message{}
		err = json.Unmarshal(buf, &msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		srv.Manager().AddMessage(msg)
		result.Data = "OK"
		buf, _ = json.Marshal(&result)
		/*if err != nil {
			httpError(w, fmt.Errorf("error serializing results: %w", err), http.StatusInternalServerError)
		}*/

		w.WriteHeader(http.StatusOK)
		w.Write(buf)
	}
}
