// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/models"
)

type FakeDomainDB struct {
	GetAllDomainsStub        func() (*models.Domains, error)
	getAllDomainsMutex       sync.RWMutex
	getAllDomainsArgsForCall []struct{}
	getAllDomainsReturns struct {
		result1 *models.Domains
		result2 error
	}
	UpsertDomainStub        func(domain string, ttl int) error
	upsertDomainMutex       sync.RWMutex
	upsertDomainArgsForCall []struct {
		domain string
		ttl    int
	}
	upsertDomainReturns struct {
		result1 error
	}
}

func (fake *FakeDomainDB) GetAllDomains() (*models.Domains, error) {
	fake.getAllDomainsMutex.Lock()
	fake.getAllDomainsArgsForCall = append(fake.getAllDomainsArgsForCall, struct{}{})
	fake.getAllDomainsMutex.Unlock()
	if fake.GetAllDomainsStub != nil {
		return fake.GetAllDomainsStub()
	} else {
		return fake.getAllDomainsReturns.result1, fake.getAllDomainsReturns.result2
	}
}

func (fake *FakeDomainDB) GetAllDomainsCallCount() int {
	fake.getAllDomainsMutex.RLock()
	defer fake.getAllDomainsMutex.RUnlock()
	return len(fake.getAllDomainsArgsForCall)
}

func (fake *FakeDomainDB) GetAllDomainsReturns(result1 *models.Domains, result2 error) {
	fake.GetAllDomainsStub = nil
	fake.getAllDomainsReturns = struct {
		result1 *models.Domains
		result2 error
	}{result1, result2}
}

func (fake *FakeDomainDB) UpsertDomain(domain string, ttl int) error {
	fake.upsertDomainMutex.Lock()
	fake.upsertDomainArgsForCall = append(fake.upsertDomainArgsForCall, struct {
		domain string
		ttl    int
	}{domain, ttl})
	fake.upsertDomainMutex.Unlock()
	if fake.UpsertDomainStub != nil {
		return fake.UpsertDomainStub(domain, ttl)
	} else {
		return fake.upsertDomainReturns.result1
	}
}

func (fake *FakeDomainDB) UpsertDomainCallCount() int {
	fake.upsertDomainMutex.RLock()
	defer fake.upsertDomainMutex.RUnlock()
	return len(fake.upsertDomainArgsForCall)
}

func (fake *FakeDomainDB) UpsertDomainArgsForCall(i int) (string, int) {
	fake.upsertDomainMutex.RLock()
	defer fake.upsertDomainMutex.RUnlock()
	return fake.upsertDomainArgsForCall[i].domain, fake.upsertDomainArgsForCall[i].ttl
}

func (fake *FakeDomainDB) UpsertDomainReturns(result1 error) {
	fake.UpsertDomainStub = nil
	fake.upsertDomainReturns = struct {
		result1 error
	}{result1}
}

var _ db.DomainDB = new(FakeDomainDB)
