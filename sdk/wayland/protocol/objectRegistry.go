package protocol

import (
	"errors"
	"sync"
)

var (
	ErrObjectDoesNotExist = errors.New("object does not exist")
)

type proxy struct {
	id       uint32
	registry *ObjectRegistry
}

type dispatcher interface {
	Dispatch(header MessageHeader, data []byte) error
}

type ObjectRegistry struct {
	sync.RWMutex
	objects      map[uint32]dispatcher
	nextObjectId uint32
}

func NewRegistry() *ObjectRegistry {
	return &ObjectRegistry{
		objects: map[uint32]dispatcher{
			1: &WlDisplayDispatcher{},
		},
		nextObjectId: 2,
	}
}

func (r *ObjectRegistry) getProxy(objectID uint32) (proxy, error) {
	if _, ok := r.objects[objectID]; !ok {
		return proxy{}, ErrObjectDoesNotExist
	}

	return proxy{
		id:       objectID,
		registry: r,
	}, nil
}
