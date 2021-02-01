package processes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"os/user"
	"sync"

	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/memory"
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/queue"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/simperr"
)

// ProcessManager é a interface que define os métodos do gerenciador de processos
type ProcessManager interface {
	Run(ctx context.Context, wg *sync.WaitGroup)
	GetJob(id string) (meta.Job, error)
	CreateFirstJob()
	DeleteJob(jobID string) error
}

var pm ProcessManager

// Conf define a estrura do arquivo de configuração do SIMP
type Conf struct {
	MaxMemUsage int `json:"maxMemusage"` // Memory in bytes
	MaxCPUUsage int `json:"maxCPUUsage"` // CPU cores
}

// GetProcessManager define a estrutura que consegue utilizar os métodos de envio de processos para execução na máquina
func GetProcessManager() ProcessManager {
	if pm == nil {
		// Get params from config file
		config := Conf{}

		usr, err := user.Current()
		homeFolder := fmt.Sprintf("%s/.simp/config.json", usr.HomeDir)
		println(homeFolder)
		configFile, err := ioutil.ReadFile(homeFolder)

		if err == nil {
			err = json.Unmarshal(configFile, &config)
		} else {
			panic("Config file not found!")
		}

		pm = newProcessManager(config.MaxMemUsage, config.MaxCPUUsage)
	}
	return pm
}

func newProcessManager(maxMemUsage int, maxCPUUsage int) ProcessManager {
	return &processes{
		maxMemUsage:   maxMemUsage,
		curMemUsage:   0,
		maxCPUUsage:   maxCPUUsage,
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
	processesContextMap map[string]context.CancelFunc
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
	jobCxt, jobCancel := context.WithCancel(ctx)

	p.processesContextMap[job.ID] = jobCancel

	go func() {
		cmd := exec.CommandContext(jobCxt, job.ProcessName, job.ProcessParams...)
		cmd.Dir = job.WorkingDirectory
		done := make(chan bool, 1)

		go func() {
			cmd.Run()
			done <- true
		}()

		<-done
		p.releaseJob(job)
	}()
}

func (p *processes) releaseJob(job meta.Job) {
	p.Lock()
	defer p.Unlock()
	p.curCPUUsage -= job.MinCPU
	p.curMemUsage -= job.MinMemory
	delete(p.processesContextMap, job.ID)
	memory.DeleteJob(job.ID)
}

// DeleteJob cancela o contexto de execução de um job e, assim, encerra sua execução
func (p *processes) DeleteJob(jobID string) error {
	p.Lock()
	defer p.Unlock()
	if cancel, ok := p.processesContextMap[jobID]; ok {
		cancel()
	} else {
		return simperr.NewError().DoesNotExist().Message("couldn't find job for given ID").Build()
	}

	return nil
}
