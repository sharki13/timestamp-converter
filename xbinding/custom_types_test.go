package xbinding

import (
	"reflect"
	"testing"
	"time"

	"com.sharki13/timestamp.converter/timezone"
)

func TestTime(t *testing.T) {
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

func isPresetsEqual(a, b []timezone.Preset) bool {
	return reflect.DeepEqual(a, b)
}

func TestTimePreset(t *testing.T) {
	type args struct {
		value []timezone.Preset
	}
	tests := []struct {
		name    string
		tr      Presets
		args    args
		wantErr bool
	}{
		{
			name: "TestTime_Set",
			tr:   NewPresets(),
			args: args{
				value: []timezone.Preset{
					{
						Id:        0,
						Label:     "Home",
						Timezones: []int{1, 2, 3},
					},
					{
						Id:        1,
						Label:     "Work",
						Timezones: []int{1, 2, 3, 5},
					},
				},
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

			if !isPresetsEqual(got, tt.args.value) {
				t.Errorf("Time.Get() = %v, want %v", got, tt.args.value)
			}
		})
	}
}
