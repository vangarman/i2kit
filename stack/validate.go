package stack

import (
	"fmt"
	"strings"
)

//TODO: add defaults

func validateStack(s *Stack) error {
	for _, service := range s.Services {
		if err := validateLinks(service, s); err != nil {
			return err
		}
		if err := validatePorts(service); err != nil {
			return err
		}
		if len(service.Containers) == 0 {
			return fmt.Errorf("No containers defined for service '%s'", service.Name)
		}
	}
	return nil
}

func validateLinks(service *Service, s *Stack) error {
	for _, link := range service.Links {
		if _, ok := s.Services[link.Service]; ok == false {
			return fmt.Errorf("Link '%s' in service '%s' does not exist", link.Alias, service.Name)
		}
	}
	return nil
}

func validatePorts(service *Service) error {
	for _, port := range service.Ports {
		if _, ok := service.Containers[port.Container]; ok == false {
			return fmt.Errorf("Container '%s' in service '%s' ports does not exist", port.Container, service.Name)
		}
		port.Protocol = strings.ToLower(port.Protocol)
		switch port.Protocol {
		case HTTP, TCP, UDP:
			continue
		default:
			return fmt.Errorf("Port protocol in service '%s' must be one of http, tcp or udp", service.Name)
		}
	}
	return nil
}
