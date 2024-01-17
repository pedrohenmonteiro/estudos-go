package events

import "time"

// Evento (Carrega dados)
type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

// Operaçoes que serão executadas quando um evento é chamado
type EventHandlerInterface interface {
	Handle(event EventInterface)
}

// Gerenciador dos nossos eventos/operaçoes
// Registrar os eventos e suas operaçoes
// Despachar / Fire no evento para que sua operaçoes sejam executadas
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear()
}
