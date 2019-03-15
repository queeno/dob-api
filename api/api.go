package api

import (
  "log"
  "context"
  "net/http"
  "github.com/gorilla/mux"
)

func putUser(w http.ResponseWriter, r *http.Request){
  log.Info("Put User function")
}

func getUser(w http.ResponseWriter, r *http.Request){
  log.Info("Get User function")
}

func Run() int {
  router := mux.NewRouter()

  // Adding routes
  router.HandleFunc("/hello/{username:[a-zA-Z0-9]+}", putUser).Methods("PUT")
  router.HandleFunc("/hello/{username:[a-zA-Z0-9]+}", getUser).Methods("PUT")

  // Create a server
  log.Info("Start server on 0.0.0.0:8000")
  server := &http.Server{
      Addr:         "0.0.0.0:8000",
      WriteTimeout: time.Second * 15,
      ReadTimeout:  time.Second * 15,
      IdleTimeout:  time.Second * 60,
      Handler: router
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

  ctx, cancel := context.WithTimeout(context.Background(), wait)
  defer cancel()

  server.Shutdown(ctx)

  log.Println("shutting down")
  return 0
}
