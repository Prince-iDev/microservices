package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Prince-iDev/microservices/handlers"
)

func main() {

	// server := gin.Default()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hp := handlers.NewPackage(l)
	gb := handlers.NewBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hp)
	sm.Handle("/goodbye", gb)

	st := &http.Server{
		Addr:         ":5050",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := st.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	sig := <-signalChan
	l.Println("Graceful shutdown: Termination recieved!", sig)

	// ctx :=
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	st.Shutdown(ctx)

}
