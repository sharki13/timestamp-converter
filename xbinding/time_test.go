package xbinding

import (
	"testing"
	"time"
)

func TestTime_Set(t *testing.T) {
	type args struct {
		value time.Time
	}
	tests := []struct {
		name    string
		tr      Time
		args    args
		wantErr bool
	}{
		{
			name: "TestTime_Set",
			tr:   NewTime(),
			args: args{
				value: time.Now(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.Set(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Time.Set() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := tt.tr.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Time.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.args.value {
				t.Errorf("Time.Get() = %v, want %v", got, tt.args.value)
			}
		})
	}
}
