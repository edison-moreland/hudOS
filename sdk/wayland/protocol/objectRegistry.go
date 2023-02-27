package protocol

import (
	"errors"
	"github.com/edison-moreland/nreal-hud/sdk/log"
	"sync"
)

var (
	ErrObjectDoesNotExist = errors.New("object does not exist")
)

type Dispatcher interface {
	Dispatch(header MessageHeader, data []byte) error
	AttachListener(listener interface{})
}

type Proxy struct {
	id       uint32
	registry *ObjectRegistry
}

func (p *Proxy) Dispatcher() Dispatcher {
	d, err := p.registry.GetDispatch(p.id)
	if err != nil {
		log.Panic().
			Err(err).
			Msg("Proxy used after object destroyed")
	}

	return d
}

func (p *Proxy) GetId() uint32 {
	return p.id
}

type ObjectRegistry struct {
	UnimplementedWlDisplayListener
	sync.RWMutex
	objects       map[uint32]Dispatcher
	_nextObjectId uint32
}

func NewRegistry() *ObjectRegistry {
	reg := &ObjectRegistry{
		_nextObjectId: 2,
	}

	displayDispatcher := NewWlDisplayDispatcher()
	displayDispatcher.AttachListener(reg)

	reg.objects = map[uint32]Dispatcher{
		1: displayDispatcher,
	}

	return reg
}

func (r *ObjectRegistry) Error(e WlDisplayErrorEvent) {
	log.Error().
		Uint32("object_id", e.ObjectId).
		Uint32("code", e.Code).
		Str("source", "compositor").
		Msg(e.Message)
}

func (r *ObjectRegistry) DeleteId(e WlDisplayDeleteIdEvent) {
	log.Info().
		Uint32("object_id", e.Id).
		Msg("object deleted")

	r.Lock()
	defer r.Unlock()

	delete(r.objects, e.Id)
}

func (r *ObjectRegistry) GetProxy(objectID uint32) (Proxy, error) {
	r.RLock()
	defer r.RUnlock()

	if _, ok := r.objects[objectID]; !ok {
		return Proxy{}, ErrObjectDoesNotExist
	}

	return Proxy{
		id:       objectID,
		registry: r,
	}, nil
}

func (r *ObjectRegistry) GetDispatch(objectID uint32) (Dispatcher, error) {
	r.RLock()
	defer r.RUnlock()

	d, ok := r.objects[objectID]
	if !ok {
		return nil, ErrObjectDoesNotExist
	}

	return d, nil
}

func (r *ObjectRegistry) nextObjectId() uint32 {
	n := r._nextObjectId
	r._nextObjectId++
	return n
}

func (r *ObjectRegistry) NewWlRegistry() WlRegistry {
	r.Lock()
	defer r.Unlock()

	id := r.nextObjectId()
	r.objects[id] = NewWlRegistryDispatcher()

	proxy := Proxy{
		id:       id,
		registry: r,
	}

	return NewWlRegistry(proxy)
}
