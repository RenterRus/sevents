package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"storing_events/internal/db"
	v1 "storing_events/internal/server/v1"
	"sync"
	"syscall"
)

func RunSevents(cmdSevents *cobra.Command, args []string) {
	l, _ := cmdSevents.Flags().GetString("listen")
	m, _ := cmdSevents.Flags().GetString("mongo")
	d, _ := cmdSevents.Flags().GetString("db")
	c, _ := cmdSevents.Flags().GetString("collection")

	serv := v1.NewServer(l, db.GetMongoClient(m, d, c))

	var wg sync.WaitGroup
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Printf("The HTTP server is raised with the parameters:\n  Server addr: %s\n  Monga addr:  %s\n  DB name:     %s\n  Coll name:   %s\n",
			l, m, d, c)
		if err := serv.Start(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	go func() {
		defer wg.Done()

		<-sign
		fmt.Println("\nServer stopping")
		serv.GraceShutdown()
		fmt.Println("DB stopping")
		serv.Mongo.Session.Close()
		fmt.Println("Gracefull shutdown complete")
	}()

	wg.Wait()
}
