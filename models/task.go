// Package models
package models

type Task struct {
	ID          int    `yaml:"Id"`
	Description string `yaml:"Description"`
	State       string `yaml:"State"`
}
