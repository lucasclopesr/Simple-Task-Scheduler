package jobhandler

import (
	"reflect"
	"testing"

	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/api/handlers"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
)

func TestNewJobHandler(t *testing.T) {
	tests := []struct {
		name string
		want handlers.JobHandler
	}{
		/*{
			name: "Teste instanciação",
			want: &jobHandler{},
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJobHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJobHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jobHandler_CreateJob(t *testing.T) {
	type args struct {
		s  string
		jr meta.JobRequest
	}
	tests := []struct {
		name    string
		j       jobHandler
		args    args
		wantErr bool
	}{
		/*{
			name: "Teste criação de Job",
			j:    jobHandler{},
			args: args{
				s: "ID1",
				jr: meta.JobRequest{
					User: "usuario-teste",
					Job:  meta.Job{},
				},
			},
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.j.CreateJob(tt.args.s, tt.args.jr); (err != nil) != tt.wantErr {
				t.Errorf("jobHandler.CreateJob() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_jobHandler_DeleteJob(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		j       jobHandler
		args    args
		wantErr bool
		init    *meta.JobRequest
	}{
		/*{
			name: "Teste deletar job",
			j:    jobHandler{},
			args: args{
				s: "ID1",
			},
			init: &meta.JobRequest{
				User: "usuario-teste",
				Job: meta.Job{
					ID: "ID1",
				},
			},
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.init != nil {
				tt.j.CreateJob(tt.args.s, *tt.init)
			}

			if err := tt.j.DeleteJob(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("jobHandler.DeleteJob() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_jobHandler_GetJob(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		j       jobHandler
		args    args
		wantErr bool
		init    *meta.JobRequest
	}{
		/*{
			name: "Teste recuperar job",
			j:    jobHandler{},
			args: args{
				s: "ID1",
			},
			init: &meta.JobRequest{
				User: "usuario-teste",
				Job: meta.Job{
					ID: "ID1",
				},
			},
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.init != nil {
				tt.j.CreateJob(tt.args.s, *tt.init)
			}

			got, err := tt.j.GetJob(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("jobHandler.GetJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.init.Job) {
				t.Errorf("jobHandler.GetJob() = %v, want %v", got, tt.init.Job)
			}
		})
	}
}

func Test_jobHandler_GetExecutingJobs(t *testing.T) {
	tests := []struct {
		name    string
		j       jobHandler
		want    []meta.Job
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetExecutingJobs()
			if (err != nil) != tt.wantErr {
				t.Errorf("jobHandler.GetExecutingJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jobHandler.GetExecutingJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jobHandler_DeleteExecutingJobs(t *testing.T) {
	tests := []struct {
		name    string
		j       *jobHandler
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.j.DeleteExecutingJobs(); (err != nil) != tt.wantErr {
				t.Errorf("jobHandler.DeleteExecutingJobs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_jobHandler_GetQueuedJobs(t *testing.T) {
	tests := []struct {
		name    string
		j       jobHandler
		want    []meta.Job
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetQueuedJobs()
			if (err != nil) != tt.wantErr {
				t.Errorf("jobHandler.GetQueuedJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jobHandler.GetQueuedJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jobHandler_DeleteQueuedJobs(t *testing.T) {
	tests := []struct {
		name    string
		j       *jobHandler
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.j.DeleteQueuedJobs(); (err != nil) != tt.wantErr {
				t.Errorf("jobHandler.DeleteQueuedJobs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
