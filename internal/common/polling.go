// file: internal/common/polling.go
// description: Common polling utilities and constants used throughout the package.
//
// Copyright 2025 Andrew Donelson. All rights reserved.
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

// Package common provides shared utilities and types for internal package use.
package common

import (
	"time"
)

// PollingMode represents the update speed for detecting input.
type PollingMode uint8

// Polling mode constants define how frequently input is checked.
const (
	// EcoMode checks input 10 times per second, suitable for standard console applications
	EcoMode PollingMode = iota
	// NormalMode checks input 30 times per second, suitable for most interactive applications
	NormalMode
	// GameMode checks input 60 times per second, suitable for games and real-time applications
	GameMode
)

// PollingConfig provides configuration options for polling-based watchers.
type PollingConfig struct {
	// Mode determines the frequency of polling
	Mode PollingMode

	// CustomInterval allows for a custom polling interval instead of using predefined modes
	// If set to a non-zero value, this overrides the Mode setting
	CustomInterval time.Duration

	// BufferSize sets the size of event channel buffers
	BufferSize int
}

// DefaultPollingConfig returns the default polling configuration.
func DefaultPollingConfig() PollingConfig {
	return PollingConfig{
		Mode:           EcoMode,
		CustomInterval: 0,
		BufferSize:     10,
	}
}

// GetInterval returns the polling interval for the specified mode or custom interval.
func (pc PollingConfig) GetInterval() time.Duration {
	// If a custom interval is specified, use it
	if pc.CustomInterval > 0 {
		return pc.CustomInterval
	}

	// Otherwise, determine interval based on mode
	switch pc.Mode {
	case EcoMode:
		return time.Second / 10 // 100ms
	case NormalMode:
		return time.Second / 30 // ~33ms
	case GameMode:
		return time.Second / 60 // ~16.6ms
	default:
		return time.Second / 10 // Default to EcoMode
	}
}

// Poller provides a standardized mechanism for polling-based operations.
type Poller struct {
	config PollingConfig
	ticker *time.Ticker
}

// NewPoller creates a new Poller with the specified configuration.
func NewPoller(config PollingConfig) *Poller {
	return &Poller{
		config: config,
	}
}

// Start begins the polling process, returning a ticker channel that signals when to poll.
// It's the caller's responsibility to stop the ticker when done.
func (p *Poller) Start() <-chan time.Time {
	interval := p.config.GetInterval()
	p.ticker = time.NewTicker(interval)
	return p.ticker.C
}

// Stop ends the polling process and cleans up resources.
func (p *Poller) Stop() {
	if p.ticker != nil {
		p.ticker.Stop()
		p.ticker = nil
	}
}

// UpdateInterval changes the polling interval.
// If the poller is currently running, it will be restarted with the new interval.
func (p *Poller) UpdateInterval(interval time.Duration) {
	wasRunning := p.ticker != nil

	// Stop if running
	if wasRunning {
		p.Stop()
	}

	// Update config
	p.config.CustomInterval = interval

	// Restart if was running
	if wasRunning {
		p.Start()
	}
}

// UpdateMode changes the polling mode.
// If the poller is currently running, it will be restarted with the new mode.
func (p *Poller) UpdateMode(mode PollingMode) {
	wasRunning := p.ticker != nil

	// Stop if running
	if wasRunning {
		p.Stop()
	}

	// Update config
	p.config.Mode = mode
	p.config.CustomInterval = 0 // Clear custom interval

	// Restart if was running
	if wasRunning {
		p.Start()
	}
}
