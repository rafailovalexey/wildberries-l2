package provider

import (
	"github.com/emptyhopes/wildberries-l2-dev11/internal/controller"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/converter"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/repository"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/service"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/validation"
)

type ProviderEventsInterface interface {
	GetControllerEvents() controller.ControllerEventsInterface
	GetServiceEvents() service.ServiceEventsInterface
	GetRepositoryEvents() repository.RepositoryEventsInterface
	GetConverterEvents() converter.ConverterEventsInterface
	GetValidationEvents() validation.ValidationEventsInterface
}
