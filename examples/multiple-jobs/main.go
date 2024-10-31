package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jakewan/gcron"
)

func main() {
	scheduler := gcron.NewScheduler()
	// Create an unstarted job to run every even second.
	if err := scheduler.AddJob("Job 1", "*/2 * * * * *"); err != nil {
		log.Fatal(err)
	} else
	// Create an unstarted job to run every odd second.
	if err := scheduler.AddJob("Job 2", "1-59/2 * * * * *"); err != nil {
		log.Fatal(err)
	} else
	// Start the first job.
	if err := scheduler.StartJob("Job 1"); err != nil {
		log.Fatal(err)
	} else
	// Start the second job.
	if err := scheduler.StartJob("Job 2"); err != nil {
		log.Fatal(err)
	} else {
		// Continue processing until the user quits the application.
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		done := make(chan bool, 1)
		go func() {
			log.Print("Press Ctrl+C to quit")
			sig := <-sigs
			log.Print("Received signal: ", sig)
			done <- true
		}()
		<-done
		log.Print("Process completed")
	}
}
