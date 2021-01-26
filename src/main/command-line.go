package main

import (
	"flag"
	"os"
)

func launch(job *string) {
	if *job == "" {
		print("Job id must not be null!")
	}
	print("Launching job...")
	//TODO
}

func delete(job *string) {
	if *job == "jobs" {

		//TODO delete running jobs

	} else if *job == "queue" {

		//TODO queue jobs

	} else if *job == "" {

		print("Job id must not be null!")

	}
	print("Deleting job...")
	//TODO
}

func status(job *string) {
	if *job == "jobs" {

		//TODO status jobs

	} else if *job == "queue" {

		//TODO queue jobs

	} else if *job == "" {

		print("Job id must not be null!")

	}

	print("Job status...")
	//TODO job status
}

func main() {
	command := flag.String("command", "", "Command{launch|delete|status} (Required)")
	job := flag.String("job_id", "", "Job ID")
	flag.Parse()

	switch *command {
	case "launch":
		launch(job)
	case "delete":
		delete(job)
	case "status":
		status(job)
	default:
		print("Invalid command. Acepted values: {launch|delete|status}")
		os.Exit(1)
	}
}
