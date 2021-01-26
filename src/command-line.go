package main

import (
	"flag"
	"os"
)

func launchJob(job *string) {
	//TODO
}

func launch(job *string) {
	if *job == "" {
		print("Job id must not be null!")
	} else {
		launchJob(job)
	}
}

func deleteAllRunningJobs() {
	//TODO
}

func deleteAllJobsOnQueue() {
	//TODO

}

func deleteJob(job *string) {
	//TODO
}

func delete(job *string) {
	if *job == "jobs" {

		deleteAllRunningJobs()

	} else if *job == "queue" {

		deleteAllJobsOnQueue()

	} else if *job == "" {

		deleteJob(job)

	}
}

func statusRunningJobs() {
	//TODO
}

func statusQueueJobs() {
	//TODO
}

func jobStatus(job *string) {
	//TODO
}

func status(job *string) {
	if *job == "jobs" {

		statusRunningJobs()

	} else if *job == "queue" {

		statusQueueJobs()

	} else if *job == "" {

		print("Job id must not be null!")

	}

	jobStatus(job)

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
