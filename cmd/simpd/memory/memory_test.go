package memory

import (
	"reflect"
	"testing"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
)

func Test_memory_CreateJob(t *testing.T) {
	type args struct {
		id  string
		job meta.Job
	}
	tests := []struct {
		name    string
		mem     memory
		args    args
		wantErr bool
		want    memory
	}{
		{
			name: "creating job on empty memory",
			mem:  memory{},
			args: args{
				id: "123",
				job: meta.Job{
					ID:          "123",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
			},
			wantErr: false,
			want: memory{
				"123": meta.Job{
					ID:          "123",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
			},
		},
		{
			name: "creating conflicting job",
			mem: memory{
				"123": meta.Job{
					ID:          "123",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
			},
			args: args{
				id: "123",
				job: meta.Job{
					ID:          "123",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
			},
			wantErr: true,
			want: memory{
				"123": meta.Job{
					ID:          "123",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
			},
		},
		{
			name: "creating job on big memory",
			mem: memory{
				"123": meta.Job{
					ID:          "123",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
			},
			args: args{
				id: "1234",
				job: meta.Job{
					ID:          "1234",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
			},
			wantErr: false,
			want: memory{
				"123": meta.Job{
					ID:          "123",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
				"1234": meta.Job{
					ID:          "1234",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.mem.CreateJob(tt.args.id, tt.args.job); (err != nil) != tt.wantErr {
				t.Errorf("memory.CreateJob() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.mem, tt.want) {
				t.Errorf("memory.CreateJob() = %v, want %v", tt.mem, tt.want)
			}
		})
	}
}

func TestGetMemory(t *testing.T) {
	tests := []struct {
		name          string
		mem           memory
		wantDifferent bool
	}{
		{
			name: "already instantiated",
			mem: memory{
				"teste": meta.Job{},
			},
			wantDifferent: false,
		},
		{
			name:          "not instantiated",
			mem:           nil,
			wantDifferent: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m = tt.mem
			if got := GetMemory(); (!(m == nil) && !reflect.DeepEqual(m, got)) || ((m == nil) && !reflect.DeepEqual(got, memory{})) {
				t.Errorf("GetMemory() = %v, want %v", got, tt.mem)
			}
		})
	}
}

func Test_memory_DeleteJob(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		mem     memory
		args    args
		wantErr bool
		want    memory
	}{
		{
			name: "delete existing job",
			mem: memory{
				"123": meta.Job{
					ID:          "123",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
			},
			args: args{
				id: "123",
			},
			wantErr: false,
			want:    memory{},
		},
		{
			name: "delete non existing job",
			mem: memory{
				"1234": meta.Job{
					ID:          "1234",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
			},
			args: args{
				id: "123",
			},
			wantErr: true,
			want: memory{
				"1234": meta.Job{
					ID:          "1234",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.mem.DeleteJob(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("memory.DeleteJob() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.mem, tt.want) {
				t.Errorf("memory.DeleteJob() = %v, want %v", tt.mem, tt.want)
			}
		})
	}
}

func Test_memory_GetJob(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		mem     memory
		args    args
		wantMem memory
		want    meta.Job
		wantErr bool
	}{
		{
			name: "get existing job",
			mem: memory{
				"1234": meta.Job{
					ID:          "1234",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
			},
			args: args{
				id: "1234",
			},
			want: meta.Job{
				ID:          "1234",
				Priority:    1,
				ProcessName: "processName1",
				ProcessParams: []string{
					"param1", "param2",
				},
				MinCPU:           10,
				MinMemory:        10,
				WorkingDirectory: "home",
				Index:            1,
			},
			wantMem: memory{
				"1234": meta.Job{
					ID:          "1234",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
			},
			wantErr: false,
		},
		{
			name: "get existing job",
			mem: memory{
				"1234": meta.Job{
					ID:          "1234",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
			},
			args: args{
				id: "123",
			},
			want: meta.Job{},
			wantMem: memory{
				"1234": meta.Job{
					ID:          "1234",
					Priority:    1,
					ProcessName: "processName1",
					ProcessParams: []string{
						"param1", "param2",
					},
					MinCPU:           10,
					MinMemory:        10,
					WorkingDirectory: "home",
					Index:            1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.mem.GetJob(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("memory.GetJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memory.GetJob() = %v, want %v", got, tt.want)
			}
		})
	}
}
