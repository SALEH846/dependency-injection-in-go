package main

import (
	"bytes"
	"errors"
	"testing"
)

type mockDatabase struct {
	setErr error

	getErr error
	getStr string
}

func (m mockDatabase) Get(key string) (string, error) {
	return m.getStr, m.getErr
}

func (m mockDatabase) Set(key, value string) error {
	return m.setErr
}

func TestRunnerArgsError(t *testing.T) {
	r := newRunner(mockDatabase{})
	if err := r.run(&bytes.Buffer{}, []string{"./kv", "help", "123"}); err == nil {
		t.Error("expected err on empty slice for args, got nil")
	}
}

func TestRunnerSetMisingArgErr(t *testing.T) {
	r := newRunner(mockDatabase{})
	if err := r.run(&bytes.Buffer{}, []string{"./kv", "set", "bob"}); err == nil {
		t.Error("expected err on empty slice for args, got nil")
	}
}

func TestRunnerReturnsErrOnSet(t *testing.T) {
	setErr := errors.New("set err")
	r := newRunner(mockDatabase{setErr: setErr})
	err := r.run(&bytes.Buffer{}, []string{"./kv", "set", "bob", "10"})
	if err == nil {
		t.Error("expected err on empty slice for args, got nil")
	}
	if err.Error() != setErr.Error() {
		t.Errorf("expected err to be %v got %v", setErr, err)
	}
}

func TestRunnerReturnsErrOnGet(t *testing.T) {
	getErr := errors.New("get err")
	r := newRunner(mockDatabase{getErr: getErr, getStr: "10"})
	err := r.run(&bytes.Buffer{}, []string{"./kv", "get", "bob"})
	if err == nil {
		t.Error("expected err on empty slice for args, got nil")
	}
	if err.Error() != getErr.Error() {
		t.Errorf("expected err to be %v got %v", getErr, err)
	}
}

func TestRunnerExpectedOutput(t *testing.T) {
	r := newRunner(mockDatabase{getStr: "10"})
	buf := &bytes.Buffer{}
	err := r.run(buf, []string{"./kv", "get", "bob"})
	if err != nil {
		t.Error("expected err to be nil on mock db get returning strings")
	}
	if buf.String() != "10\n" {
		t.Errorf("expected buffer to be 10 got %s", buf.String())
	}
}
