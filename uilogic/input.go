package uilogic

import (
	"fmt"
)


type InputReader interface {
	Read() (string, error)
}

// Console implementation
type FmtInput struct{}
func (r *FmtInput) Read() (string, error) {
	var response string
	_, err := fmt.Scanln(&response)
	return response, err
}

// Mock implementation
type MockInput struct{
	data []string
}

func (m *MockInput) Read() (string, error) {
	if len(m.data) == 0 {
		return "", nil
	} else {
		res := m.data[0]
		m.data = m.data[1:]
		return res, nil
	}
}

