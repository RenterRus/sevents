package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"storing_events/internal/db"
	v1 "storing_events/internal/server/v1"
	"sync"
	"syscall"
)

func RunSevents(cmdSevents *cobra.Command, args []string){
	l, _ := cmdSevents.Flags().GetString("listen")
	serv := v1.NewServer(l)

	var wg sync.WaitGroup
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(3)

	go func(){
		defer wg.Done()

		m, _ := cmdSevents.Flags().GetString("mongo")
		db.MongoParam.Addr = m
		d, _ := cmdSevents.Flags().GetString("db")
		db.MongoParam.DBName = d
		c, _ := cmdSevents.Flags().GetString("collection")
		db.MongoParam.CollName = c

		db.MongoParam.GetClient()
		err := db.MongoParam.Session.Ping()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("The connection to the database is established in: %s for %s -> %s", m, d, c)
		}
	}()

	go func(){
		defer wg.Done()
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
		db.MongoParam.Session.Close()
		fmt.Println("Gracefull shutdown complete")
	}()

	wg.Wait()
}