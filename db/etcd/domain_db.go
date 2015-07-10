package db

import (
	"path"

	"github.com/cloudfoundry-incubator/bbs"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/coreos/go-etcd/etcd"
	"github.com/pivotal-golang/lager"
)

const DomainSchemaRoot = DataSchemaRoot + "domain"

func (db *ETCDDB) GetAllDomains(logger lager.Logger) (*models.Domains, *bbs.Error) {
	response, err := db.client.Get(DomainSchemaRoot, false, true)
	if err != nil {
		if err.(*etcd.EtcdError).ErrorCode == ETCDErrKeyNotFound {
			return &models.Domains{}, nil
		}
		logger.Error("failed-to-fetch-domains", err)
		return nil, bbs.ErrUnknownError
	}

	domains := []string{}
	for _, child := range response.Node.Nodes {
		domains = append(domains, path.Base(child.Key))
	}

	return &models.Domains{Domains: domains}, nil
}

func (db *ETCDDB) UpsertDomain(domain string, ttl int, logger lager.Logger) *bbs.Error {
	_, err := db.client.Set(DomainSchemaPath(domain), "", uint64(ttl))
	if err != nil {
		logger.Error("failed-to-upsert-domain", err)
		return bbs.ErrUnknownError
	}
	return nil
}

func DomainSchemaPath(domain string) string {
	return path.Join(DomainSchemaRoot, domain)
}
