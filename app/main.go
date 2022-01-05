package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

// App export
type App struct {
    Router *mux.Router
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    var response map[string]interface{}
    json.Unmarshal([]byte(`{ "hello": "world" }`), &response)
    respondWithJSON(w, http.StatusOK, response)
}

func helloWorldHandlerV2(w http.ResponseWriter, r *http.Request) {
    var response map[string]interface{}
    json.Unmarshal([]byte(`{ "hello": "world", "V": 2 }`), &response)
    respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(payload)
}

func (app *App) initialiseRoutes() {
    app.Router = mux.NewRouter()
    app.Router.HandleFunc("/", helloWorldHandler)
    app.Router.HandleFunc("/v2", helloWorldHandlerV2)
}

func (app *App) run() {
    log.Fatal(http.ListenAndServe(":8080", app.Router))
}

func main() {
    app := App{}
    app.initialiseRoutes()
    app.run()
}
