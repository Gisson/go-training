package api

import (
	"encoding/json"
	//	"github.com/Gisson/simple-messenger/message"
	"github.com/Gisson/simple-messenger/server"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Result struct {
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func AddAllRoutes(server *server.Server) {
	server.AddServerHandler(http.MethodGet, "/healthz", Health)
	server.AddServerHandler(http.MethodPost, "/api/v0/addmessage", AddMessageHandler)
}

func Health(srv *server.Server) (handle httprouter.Handle) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		var result Result
		result.Data = "OK"
		buf, _ := json.Marshal(&result)
		/*if err != nil {
			httpError(w, fmt.Errorf("error serializing results: %w", err), http.StatusInternalServerError)
		}*/

		w.WriteHeader(http.StatusOK)
		w.Write(buf)
	}
}
