package handlers

import (
	"net/http"
	"time"

	"github.com/pgmonzon/Yng_Servicios/core"
)

func SetearHeaders(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, authorization")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	core.JSONResponse(w, r, start, []byte{}, http.StatusOK)
}
