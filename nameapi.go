package nameapi

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func GetName(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["NAME"]
    fmt.Fprintln(w, "Hi " + name)
}
