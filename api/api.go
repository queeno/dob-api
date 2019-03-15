package api

import (
  "log"
  "context"
  "net/http"
  "os"
  "os/signal"
  "time"

  "github.com/gorilla/mux"
)

func putUser(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
}

func getUser(w http.ResponseWriter, r *http.Request){
  log.Println("Get User function")
}

func Run() int {
  router := mux.NewRouter()

  // Adding routes
  router.HandleFunc("/hello/{username:[a-zA-Z0-9]+}", putUser).Methods("PUT")
  router.HandleFunc("/hello/{username:[a-zA-Z0-9]+}", getUser).Methods("GET")

  // Create a server
  log.Println("Starting server on 0.0.0.0:8000")
  server := &http.Server{
      Addr:         "0.0.0.0:8000",
      WriteTimeout: time.Second * 15,
      ReadTimeout:  time.Second * 15,
      IdleTimeout:  time.Second * 60,
      Handler: router,
  }

  // Run the server
  go func() {
      if err := server.ListenAndServe(); err != nil {
          log.Fatal(err)
      }
  }()

  // Catch the interrupt signal and shutdown
  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt)

  <-c

  ctx, cancel := context.WithTimeout(context.Background(), 15)
  defer cancel()

  server.Shutdown(ctx)

  log.Println("Gracefully shutting down...")
  return 0
}
