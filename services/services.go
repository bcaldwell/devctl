package services

import (
	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/utilities"
)

type Service interface {
	Setup()
	PreInstall(*parser.ConfigurationStruct)
	Install(*parser.ConfigurationStruct)
	Scripts(c *parser.ConfigurationStruct) map[string]utilities.RunCommand
	IsProjectType(*parser.ConfigurationStruct) bool
}

var ServiceList = []Service{Node{}}

func ServicesUsed(c *parser.ConfigurationStruct) []Service {
	servicesUsed := []Service{}

	for _, Service := range ServiceList {
		if Service.IsProjectType(c) {
			servicesUsed = append(servicesUsed, Service)
		}
	}

	return servicesUsed
}
