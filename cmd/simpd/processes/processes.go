package processes

import (
	"context"
	"fmt"
	"os/exec"
	"sync"

	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/memory"
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/queue"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
)

type ProcessManager interface {
	Run(ctx context.Context, wg *sync.WaitGroup)
	GetJob(id string) (meta.Job, error)
	CreateFirstJob()
}

var pm ProcessManager

// GetProcessManager define a estrutura que consegue utilizar os métodos de envio de processos para execução na máquina
func GetProcessManager() ProcessManager {
	if pm == nil {
		pm = newProcessManager()
	}
	return pm
}

func newProcessManager() ProcessManager {
	return &processes{
		maxMemUsage:   100,
		curMemUsage:   0,
		maxCPUUsage:   100,
		curCPUUsage:   0,
		release:       make(chan bool),
		Mutex:         sync.Mutex{},
		queue:         queue.GetQueueManager(),
		hasJobInFront: make(chan bool),
		processQueue:  make(chan meta.Job),
	}
}

type processes struct {
	processQueue  chan meta.Job
	maxMemUsage   int
	curMemUsage   int
	maxCPUUsage   int
	curCPUUsage   int
	release       chan bool
	queue         queue.PQInterface
	hasJobInFront chan bool
	sync.Mutex
}

func (p *processes) CreateFirstJob() {
	p.hasJobInFront <- true
}

func (p *processes) Run(ctx context.Context, wg *sync.WaitGroup) {
	go func() {
		for {
			if p.queue.Len() == 0 {
				<-p.hasJobInFront
			}
			p.processQueue <- p.queue.GetFrontJob()
		}
	}()
	go func() {
		defer wg.Done()
		for {
			select {
			case newJob := <-p.processQueue:
				for p.curMemUsage+newJob.MinMemory > p.maxMemUsage || p.curCPUUsage+newJob.MinCPU > p.maxCPUUsage {
				}
				p.Lock()
				p.startJob(ctx, newJob)
				p.Unlock()
			case <-ctx.Done():
				return
			}
		}
	}()
}
func (p *processes) GetJob(jobID string) (job meta.Job, err error) {
	return memory.GetJob(jobID)
}

func (p *processes) startJob(ctx context.Context, job meta.Job) {
	p.curCPUUsage += job.MinCPU
	p.curMemUsage += job.MinMemory
	go func() {
		cmd := exec.CommandContext(ctx, job.ProcessName, job.ProcessParams...)
		cmd.Dir = job.WorkingDirectory

		_, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		p.releaseJob(job)
	}()
}

func (p *processes) releaseJob(job meta.Job) {
	p.Lock()
	defer p.Unlock()
	p.curCPUUsage -= job.MinCPU
	p.curMemUsage -= job.MinMemory
	memory.DeleteJob(job.ID)
}
