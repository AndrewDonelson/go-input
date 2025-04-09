// file: internal/common/polling_test.go
// description: Test cases for the common polling utilities and constants.
//
// Copyright 2025 Andrew Donelson. All rights reserved.
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

package common

import (
	"testing"
	"time"
)

func TestDefaultPollingConfig(t *testing.T) {
	config := DefaultPollingConfig()

	if config.Mode != EcoMode {
		t.Errorf("Expected default mode to be EcoMode, got %v", config.Mode)
	}

	if config.CustomInterval != 0 {
		t.Errorf("Expected default custom interval to be 0, got %v", config.CustomInterval)
	}

	if config.BufferSize != 10 {
		t.Errorf("Expected default buffer size to be 10, got %d", config.BufferSize)
	}
}

func TestGetInterval(t *testing.T) {
	testCases := []struct {
		name     string
		config   PollingConfig
		expected time.Duration
	}{
		{
			name: "EcoMode",
			config: PollingConfig{
				Mode:           EcoMode,
				CustomInterval: 0,
			},
			expected: time.Second / 10,
		},
		{
			name: "NormalMode",
			config: PollingConfig{
				Mode:           NormalMode,
				CustomInterval: 0,
			},
			expected: time.Second / 30,
		},
		{
			name: "GameMode",
			config: PollingConfig{
				Mode:           GameMode,
				CustomInterval: 0,
			},
			expected: time.Second / 60,
		},
		{
			name: "CustomInterval",
			config: PollingConfig{
				Mode:           GameMode, // Should be ignored
				CustomInterval: 50 * time.Millisecond,
			},
			expected: 50 * time.Millisecond,
		},
		{
			name: "InvalidMode",
			config: PollingConfig{
				Mode:           PollingMode(99), // Invalid mode
				CustomInterval: 0,
			},
			expected: time.Second / 10, // Should default to EcoMode
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			interval := tc.config.GetInterval()
			if interval != tc.expected {
				t.Errorf("Expected interval %v, got %v", tc.expected, interval)
			}
		})
	}
}

func TestPoller(t *testing.T) {
	// Test creating a new poller
	config := DefaultPollingConfig()
	poller := NewPoller(config)

	if poller == nil {
		t.Fatal("NewPoller returned nil")
	}

	if poller.config.Mode != EcoMode {
		t.Errorf("Expected poller mode to be EcoMode, got %v", poller.config.Mode)
	}

	// Test starting the poller
	tickerChan := poller.Start()
	if tickerChan == nil {
		t.Fatal("Start() returned nil channel")
	}

	// Verify that the ticker is running
	if poller.ticker == nil {
		t.Fatal("Ticker not initialized after Start()")
	}

	// Test stopping the poller
	poller.Stop()
	if poller.ticker != nil {
		t.Fatal("Ticker not nil after Stop()")
	}

	// Test updating the interval
	poller.UpdateInterval(100 * time.Millisecond)
	if poller.config.CustomInterval != 100*time.Millisecond {
		t.Errorf("Expected custom interval to be 100ms, got %v", poller.config.CustomInterval)
	}

	// Test updating the mode
	poller.UpdateMode(GameMode)
	if poller.config.Mode != GameMode {
		t.Errorf("Expected mode to be GameMode, got %v", poller.config.Mode)
	}

	if poller.config.CustomInterval != 0 {
		t.Errorf("Expected custom interval to be reset to 0, got %v", poller.config.CustomInterval)
	}

	// Test updating interval when running
	poller.Start()
	poller.UpdateInterval(200 * time.Millisecond)
	if poller.config.CustomInterval != 200*time.Millisecond {
		t.Errorf("Expected custom interval to be 200ms, got %v", poller.config.CustomInterval)
	}
	poller.Stop()

	// Test updating mode when running
	poller.Start()
	poller.UpdateMode(NormalMode)
	if poller.config.Mode != NormalMode {
		t.Errorf("Expected mode to be NormalMode, got %v", poller.config.Mode)
	}
	poller.Stop()
}
