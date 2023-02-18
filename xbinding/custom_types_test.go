package xbinding

import (
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

func isPresetEqual(a, b timezone.Preset) bool {
	if a.Id != b.Id {
		return false
	}

	if a.Label != b.Label {
		return false
	}

	if len(a.Timezones) != len(b.Timezones) {
		return false
	}

	for i, v := range a.Timezones {
		if v != b.Timezones[i] {
			return false
		}
	}

	return true
}

func isPresetsEqual(a, b timezone.Presets) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !isPresetEqual(v, b[i]) {
			return false
		}
	}

	return true
}

func TestTimePreset(t *testing.T) {
	type args struct {
		value timezone.Presets
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
				value: timezone.Presets{
					timezone.Preset{
						Id: 0,
						Label: "UTC",
						Timezones: []int{1,2,3},
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
