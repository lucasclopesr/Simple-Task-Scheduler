package main

import (
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simp/api"
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simp/commands"
)

func main() {
	commands.Init(api.NewClient())
	commands.Run()
}
