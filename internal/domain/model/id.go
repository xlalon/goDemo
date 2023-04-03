package model

import (
	"github.com/xlalon/golee/pkg/database/mysql"
	"github.com/xlalon/golee/pkg/ecode"
)

type IdentifiedDomainObject struct {
	ID int64
}

func (id *IdentifiedDomainObject) SetId(newId int64) error {
	if newId <= 0 {
		return ecode.DomainIdInvalid
	}
	if id.Id() > 0 {
		return ecode.DomainIdChange
	}
	id.ID = newId

	return nil
}

func (id *IdentifiedDomainObject) Id() int64 {
	return id.ID
}

func (id *IdentifiedDomainObject) NextId() int64 {
	return mysql.NextID()
}
