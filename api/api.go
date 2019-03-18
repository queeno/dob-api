package api

import (
  "log"
  "context"
  "net/http"
  "os"
  "os/signal"
  "time"
  "encoding/json"
  "io/ioutil"

  "github.com/queeno/dob-api/app"
  "github.com/queeno/dob-api/db"

  "github.com/gorilla/mux"
)

type Api struct {
  router *mux.Router
  app    app.MyApp
}

type userBirthday struct {
  DateOfBirth string    `json:"dateOfBirth"`
}

type messageResponse struct {
  Message string        `json:"message"`
}


func (a Api) putUser(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  username := params["username"]

  if r.Body == nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("Please add a body to the request in the form: {\"dateOfBirth\": \"YYYY-MM-DD\" }"))
    return
  }

  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }

  var uB userBirthday

  err = json.Unmarshal(body, &uB)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }

  err = a.app.UpdateUsername(username, uB.DateOfBirth)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }
  w.WriteHeader(http.StatusNoContent)
}

func (a Api) getUser(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  username := params["username"]

  message, err := a.app.GetDateOfBirth(username)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }

  mR := &messageResponse{
    Message: message,
  }

  jsonResponse, err := json.Marshal(mR)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }

  w.WriteHeader(http.StatusOK)
  w.Write(jsonResponse)
}

func (a *Api) addRoutes() {
  a.router.HandleFunc("/hello/{username:[[:alnum:]]+}", a.putUser).Methods("PUT")
  a.router.HandleFunc("/hello/{username:[[:alnum:]]+}", a.getUser).Methods("GET")
}

func (a Api) RunServer() int {
  // Create a server
  log.Println("Starting server on 0.0.0.0:8000")
  server := &http.Server{
      Addr:         "0.0.0.0:8000",
      WriteTimeout: time.Second * 15,
      ReadTimeout:  time.Second * 15,
      IdleTimeout:  time.Second * 60,
      Handler: a.router,
  }

  idleConnsClosed := make(chan struct{})
  // Run the server
  go func() {
    // Catch the interrupt signal and shutdown
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)

    <-c
    log.Println("Gracefully shutting down...")

    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
      log.Fatal(err)
    }
    close(idleConnsClosed)
  }()

  if err := server.ListenAndServe(); err != nil {
      log.Println(err)
  }

  <-idleConnsClosed

  return 0
}

func newRouter() *mux.Router {
  return mux.NewRouter()
}

func NewApi(databaseLocation string) *Api {
  api := &Api{
    router: newRouter(),
    app: app.NewApp(&db.BoltDB{
      FilePath: databaseLocation,
    }),
  }
  api.addRoutes()
  return api
}
