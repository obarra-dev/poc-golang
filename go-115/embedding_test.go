package main_test

import (
	"fmt"
	"sync"
	"testing"
)

type person struct {
	Name string
}

type employee struct {
	person
}

func (p person) formatFullName() string {
	return fmt.Sprintf("I am %+v", p)
}

func TestEmbeddingStructIntoStruct(t *testing.T) {
	e := employee{person: person{"pablo"}}
	f := e.formatFullName()
	e.person.formatFullName()
	if e.Name != "pablo" || f != e.person.formatFullName() || f != "I am {Name:pablo}" {
		t.Error("test Error", e, f)
	}
}

type identificable interface {
	ID() string
	DI() string
}

type employeee struct {
	person
	identificable
}

type socialSecurityNumber string

func (s socialSecurityNumber) ID() string {
	return "S-ID-" + string(s)
}

func (s socialSecurityNumber) DI() string {
	return "S-DI" + string(s)
}

//It is a conflict, i solved it wrapping
func (e employeee) ID() string {
	return "E-" + e.identificable.ID()
}

func TestEmbeddingInterfaceIntoStruct(t *testing.T) {
	e := employeee{person: person{"pablo"}, identificable: socialSecurityNumber("aa1")}

	eID := e.ID()
	sID := e.identificable.ID()
	sDI := e.DI()
	if eID != "E-S-ID-aa1" || sID != "S-ID-aa1" || sDI != "S-DIaa1" || sDI != e.identificable.DI() {
		t.Error("test Error", e, eID, sID, sDI)
	}
}

type citizen interface {
	identificable
	Country() string
}

type employeeee struct {
	person
	citizen
}

type africaUnionIdentifier struct {
	id      string
	country string
}

func (a africaUnionIdentifier) ID() string {
	return "A-ID-" + a.id
}
func (a africaUnionIdentifier) DI() string {
	return "A-DI" + a.id
}
func (a africaUnionIdentifier) Country() string {
	return "A-Country" + a.country
}

// identificable iterface is embedded into citizen iterface
func TestEmbeddingInterfaceIntoInterface(t *testing.T) {
	e := employeeee{person: person{"pablo"}, citizen: africaUnionIdentifier{id: "aa1", country: "Gana"}}

	c := e.Country()
	id := e.ID()

	if c != "A-CountryGana" || id != "A-ID-aa1" {
		t.Error("test Error", e, c, id)
	}
}

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
