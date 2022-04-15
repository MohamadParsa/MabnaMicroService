package service

// TODO: implement version management.
import (
	"net/http"

	"github.com/gorilla/mux"
)

// InitializeRouter return a configurable handler.
func InitializeRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/GetLastTradeOfAllInstruments", GetLastTradeOfAllInstrumentsHandler).Methods(http.MethodGet)
	return router
}
