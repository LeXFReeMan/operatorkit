package internal

import (
	"github.com/LeXFReeMan/operatorkit/v7/pkg/resource"
	"github.com/LeXFReeMan/operatorkit/v7/pkg/resource/crud"
)

func CRUD(r resource.Interface) (crud.Interface, bool) {
	type cruder interface {
		CRUD() crud.Interface
	}

	c, ok := r.(cruder)
	if ok {
		return c.CRUD(), true
	}

	return nil, false
}
