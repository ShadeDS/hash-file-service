package util

import "github.com/heroiclabs/nakama-common/runtime"

type MockLogger struct{}

func (m *MockLogger) Debug(msg string, keyvals ...interface{}) {}
func (m *MockLogger) Info(msg string, keyvals ...interface{})  {}
func (m *MockLogger) Warn(msg string, keyvals ...interface{})  {}
func (m *MockLogger) Error(msg string, keyvals ...interface{}) {}

func (m *MockLogger) WithField(key string, v interface{}) runtime.Logger {
	return m
}

func (m *MockLogger) WithFields(fields map[string]interface{}) runtime.Logger {
	return m
}

func (m *MockLogger) Fields() map[string]interface{} {
	return nil
}
