package events

import (
	"github.com/emptyhopes/wildberries-l2-dev11/internal/controller"
	controllerEvents "github.com/emptyhopes/wildberries-l2-dev11/internal/controller/events"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/converter"
	converterEvents "github.com/emptyhopes/wildberries-l2-dev11/internal/converter/events"
	definition "github.com/emptyhopes/wildberries-l2-dev11/internal/provider"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/repository"
	repositoryEvents "github.com/emptyhopes/wildberries-l2-dev11/internal/repository/events"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/service"
	serviceEvents "github.com/emptyhopes/wildberries-l2-dev11/internal/service/events"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/validation"
	validationEvents "github.com/emptyhopes/wildberries-l2-dev11/internal/validation/events"
)

type ProviderEvents struct {
	controllerEvents controller.ControllerEventsInterface
	serviceEvents    service.ServiceEventsInterface
	repositoryEvents repository.RepositoryEventsInterface
	validationEvents validation.ValidationEventsInterface
	converterEvents  converter.ConverterEventsInterface
}

var _ definition.ProviderEventsInterface = (*ProviderEvents)(nil)

func NewProviderEvents() *ProviderEvents {
	return &ProviderEvents{}
}

func (p *ProviderEvents) GetControllerEvents() controller.ControllerEventsInterface {
	if p.controllerEvents == nil {
		p.controllerEvents = controllerEvents.NewControllerEvents(
			p.GetValidationEvents(),
			p.GetConverterEvents(),
		)
	}

	return p.controllerEvents
}

func (p *ProviderEvents) GetServiceEvents() service.ServiceEventsInterface {
	if p.serviceEvents == nil {
		p.serviceEvents = serviceEvents.NewServiceEvents(
			p.GetRepositoryEvents(),
		)
	}

	return p.serviceEvents
}

func (p *ProviderEvents) GetRepositoryEvents() repository.RepositoryEventsInterface {
	if p.repositoryEvents == nil {
		p.repositoryEvents = repositoryEvents.NewRepositoryEvents(
			p.GetConverterEvents(),
		)
	}

	return p.repositoryEvents
}

func (p *ProviderEvents) GetConverterEvents() converter.ConverterEventsInterface {
	if p.converterEvents == nil {
		p.converterEvents = converterEvents.NewConverterEvents()
	}

	return p.converterEvents
}

func (p *ProviderEvents) GetValidationEvents() validation.ValidationEventsInterface {
	if p.validationEvents == nil {
		p.validationEvents = validationEvents.NewValidationEvents()
	}

	return p.validationEvents
}
