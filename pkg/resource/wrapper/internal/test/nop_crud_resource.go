package test

import (
	"fmt"

	"github.com/giantswarm/micrologger/microloggertest"

	"github.com/LeXFReeMan/operatorkit/v7/pkg/resource"
	"github.com/LeXFReeMan/operatorkit/v7/pkg/resource/crud"
)

func NewNopCRUDResource() resource.Interface {
	c := crud.ResourceConfig{
		CRUD:   NewNopCRUD(),
		Logger: microloggertest.New(),
	}

	r, err := crud.NewResource(c)
	if err != nil {
		panic(fmt.Sprintf("%#v", err))
	}

	return r
}
