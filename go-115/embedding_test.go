package main_test

import (
	"fmt"
	"sync"
	"testing"
)

type Job struct {
	Command string
	*sync.Mutex
}

func TestEmbeddingPointerExMutex(t *testing.T) {
	job := &Job{"Omar-", &sync.Mutex{}}

	job.Lock()
	job.Command += "Commander"
	job.Unlock()

	if job.Command != "Omar-Commander" {
		t.Error("test Error", job)
	}
}

type person struct {
	Name string
}

func (p person) format() string {
	return fmt.Sprintf("I am %+v", p)
}

type employee struct {
	person
}

func TestEmbeddingStructs(t *testing.T) {
	e := employee{person{"pablo"}}
	f := e.format()
	if e.Name != "pablo" || f != "I am {Name:pablo}" {
		t.Error("test Error", e, f)
	}
}

type personmanInterface interface {
	format() string
}

type employeeInterface interface {
	personmanInterface
}

func TestEmbeddingInterfce(t *testing.T) {
	var e employeeInterface = employee{person{"pablo"}}
	f := e.format()
	if f != "I am {Name:pablo}" {
		t.Error("test Error", e, f)
	}
}
