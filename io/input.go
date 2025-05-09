package io

import (
	"bufio"
	"os"
)


type InputReader interface {
	Read() (string, error)
}

// Console implementation
type FmtInput struct{}
func (r *FmtInput) Read() (string, error) {
	reader := bufio.NewReader(os.Stdin)
    response, err := reader.ReadString('\n') // Read until newline
    if err != nil {
        return "", err
    }

    // Remove the trailing newline character
    response = response[:len(response)-1]
    return response, nil
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

