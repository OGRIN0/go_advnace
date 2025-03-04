package main

import (
	"errors"
	"testing"
	"time"
)

type ReceiverCommandConfig struct {
	Name           string        `json:"name"`
	Command        string        `json:"command"`
	Args           []string      `json:"args,omitempty"`
	Timeout        time.Duration `json:"timeout"`
	RetryCount     int           `json:"retry_count"`
	RetryDelay     time.Duration `json:"retry_delay"`
	WorkingDir     string        `json:"working_dir,omitempty"`
	RequireSuccess bool          `json:"require_success"`
}

func (c *ReceiverCommandConfig) Validate() error {
	if c.Name == "" {
		return errors.New("name cannot be empty")
	}
	
	if c.Command == "" {
		return errors.New("command cannot be empty")
	}
	
	if c.Timeout <= 0 {
		return errors.New("timeout must be greater than 0")
	}
	
	if c.RetryCount < 0 {
		return errors.New("retry count cannot be negative")
	}
	
	if c.RetryCount > 0 && c.RetryDelay <= 0 {
		return errors.New("retry delay must be greater than 0 when retries are enabled")
	}
	
	return nil
}

func TestReceiverCommandConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  ReceiverCommandConfig
		wantErr bool
	}{
		{
			name: "valid configuration",
			config: ReceiverCommandConfig{
				Name:           "test-command",
				Command:        "echo",
				Args:           []string{"hello", "world"},
				Timeout:        5 * time.Second,
				RetryCount:     3,
				RetryDelay:     1 * time.Second,
				RequireSuccess: true,
			},
			wantErr: false,
		},
		{
			name: "missing name",
			config: ReceiverCommandConfig{
				Command:        "echo",
				Timeout:        5 * time.Second,
				RetryCount:     0,
				RequireSuccess: true,
			},
			wantErr: true,
		},
		{
			name: "missing command",
			config: ReceiverCommandConfig{
				Name:           "test-command",
				Timeout:        5 * time.Second,
				RetryCount:     0,
				RequireSuccess: true,
			},
			wantErr: true,
		},
		{
			name: "invalid timeout",
			config: ReceiverCommandConfig{
				Name:           "test-command",
				Command:        "echo",
				Timeout:        0,
				RetryCount:     0,
				RequireSuccess: true,
			},
			wantErr: true,
		},
		{
			name: "negative retry count",
			config: ReceiverCommandConfig{
				Name:           "test-command",
				Command:        "echo",
				Timeout:        5 * time.Second,
				RetryCount:     -1,
				RequireSuccess: true,
			},
			wantErr: true,
		},
		{
			name: "retry without delay",
			config: ReceiverCommandConfig{
				Name:           "test-command",
				Command:        "echo",
				Timeout:        5 * time.Second,
				RetryCount:     3,
				RetryDelay:     0,
				RequireSuccess: true,
			},
			wantErr: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReceiverCommandConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
