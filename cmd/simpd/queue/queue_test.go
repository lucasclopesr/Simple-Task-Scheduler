/* Funções testadas - Francisco:
-> Swap
-> Push
-> GetJobFromQueue
-> DeleteJobFromQueue
*/

package queue

import (
	"reflect"
	"testing"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
)

func mockQueueManager() *SimpQueueManager {
	return &SimpQueueManager{
		simpQueue: &meta.PriorityQueue{
			Queue: []*meta.Job{
				{
					Index:            0,
					ID:               "Job1",
					Priority:         2,
					ProcessName:      "ProcessOfJob1",
					ProcessParams:    nil,
					MinCPU:           10,
					MinMemory:        20,
					WorkingDirectory: "home/",
				},
				{
					Index:            1,
					ID:               "Job2",
					Priority:         4,
					ProcessName:      "ProcessOfJob2",
					ProcessParams:    nil,
					MinCPU:           20,
					MinMemory:        10,
					WorkingDirectory: "tmp/",
				},
			},
			IndexList: map[string]int{
				"Job1": 0,
				"Job2": 1,
			},
		},
	}
}

func TestSimpQueueManager_Swap(t *testing.T) {
	type args struct {
		job1Index int
		job2Index int
	}
	tests := []struct {
		name string
		pq   *SimpQueueManager
		args args
	}{
		{
			name: "Troca Job1 para Index 1 e Job2 para Index 0 da fila",
			pq:   mockQueueManager(),
			args: args{
				job1Index: 0,
				job2Index: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Swap(tt.args.job1Index, tt.args.job2Index)
			if tt.pq.simpQueue.Queue[0].ID != "Job2" || tt.pq.simpQueue.Queue[1].ID != "Job1" ||
				tt.pq.simpQueue.IndexList["Job1"] != 1 || tt.pq.simpQueue.IndexList["Job2"] != 0 {

				t.Errorf("SimpQueueManager.Swap() = %v", tt.pq.simpQueue)
			}
		})
	}
}

func TestSimpQueueManager_Push(t *testing.T) {
	type args struct {
		h interface{}
	}
	tests := []struct {
		name string
		pq   *SimpQueueManager
		args args
	}{
		{
			name: "Insere um novo job (Job3) na fila de prioridades",
			pq:   mockQueueManager(),
			args: args{
				h: &meta.Job{
					ID:               "Job3",
					Priority:         1,
					ProcessName:      "ProcessOfJob3",
					ProcessParams:    nil,
					MinCPU:           50,
					MinMemory:        50,
					WorkingDirectory: "rnd/fldr/",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Push(tt.args.h)
			if tt.pq.simpQueue.Queue[2].ID != "Job3" || tt.pq.Len() != 3 {
				t.Errorf("SimpQueueManager.Swap() = %v", tt.pq.simpQueue)
			}
		})
	}
}

func TestSimpQueueManager_GetJobFromQueue(t *testing.T) {
	type fields struct {
		simpQueue *meta.PriorityQueue
	}
	type args struct {
		jobID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantJob meta.Job
		wantErr bool
	}{
		{
			name: "Retorna job de ID 'Job1' definindo seu index como -1",
			fields: fields{
				simpQueue: mockQueueManager().simpQueue,
			},
			args: args{
				jobID: "Job1",
			},
			wantErr: false,
			wantJob: meta.Job{
				Index:            -1,
				ID:               "Job1",
				Priority:         2,
				ProcessName:      "ProcessOfJob1",
				ProcessParams:    nil,
				MinCPU:           10,
				MinMemory:        20,
				WorkingDirectory: "home/",
			},
		},
		{
			name: "Erro - Job não existe na fila de prioridades",
			fields: fields{
				simpQueue: mockQueueManager().simpQueue,
			},
			args: args{
				jobID: "Job12",
			},
			wantErr: true,
			wantJob: meta.Job{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := &SimpQueueManager{
				simpQueue: tt.fields.simpQueue,
			}
			gotJob, err := pq.GetJobFromQueue(tt.args.jobID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SimpQueueManager.GetJobFromQueue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotJob, tt.wantJob) {
				t.Errorf("SimpQueueManager.GetJobFromQueue() = %v, want %v", gotJob, tt.wantJob)
			}
		})
	}
}

func TestSimpQueueManager_DeleteJobFromQueue(t *testing.T) {
	type fields struct {
		simpQueue *meta.PriorityQueue
	}
	type args struct {
		jobID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    meta.Job
		wantErr bool
	}{
		{
			name: "Deleta job de ID 'Job1' da fila e retorna-o",
			fields: fields{
				simpQueue: mockQueueManager().simpQueue,
			},
			args: args{
				jobID: "Job1",
			},
			wantErr: false,
			want: meta.Job{
				Index:            -1,
				ID:               "Job1",
				Priority:         2,
				ProcessName:      "ProcessOfJob1",
				ProcessParams:    nil,
				MinCPU:           10,
				MinMemory:        20,
				WorkingDirectory: "home/",
			},
		},
		{
			name: "Erro - tenta deletar job inexistente na fila",
			fields: fields{
				simpQueue: mockQueueManager().simpQueue,
			},
			args: args{
				jobID: "Job12",
			},
			wantErr: true,
			want:    meta.Job{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := &SimpQueueManager{
				simpQueue: tt.fields.simpQueue,
			}
			got, err := pq.DeleteJobFromQueue(tt.args.jobID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SimpQueueManager.DeleteJobFromQueue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && (!reflect.DeepEqual(got, tt.want) || len(tt.fields.simpQueue.Queue) > 1) {
				t.Errorf("SimpQueueManager.DeleteJobFromQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
