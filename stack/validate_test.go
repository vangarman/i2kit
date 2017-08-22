package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateStackOK(t *testing.T) {
	s := &Stack{
		Name: "test",
		Services: map[string]*Service{
			"a": &Service{
				Name: "a",
				Containers: map[string]*Container{
					"a": &Container{
						Name:  "a",
						Image: "Image",
					},
				},
			},
		},
	}
	err := validateStack(s)
	require.NoError(t, err)
}

func TestValidateStackWrongLink(t *testing.T) {
	s := &Stack{
		Name: "test",
		Services: map[string]*Service{
			"a": &Service{
				Name: "a",
				Links: []*Link{
					&Link{
						Service: "b",
						Alias:   "b",
					},
				},
				Containers: map[string]*Container{
					"a": &Container{
						Name:  "a",
						Image: "Image",
					},
				},
			},
		},
	}
	err := validateStack(s)
	require.Error(t, err)
}

func TestValidateStackWrongPort(t *testing.T) {
	s := &Stack{
		Name: "test",
		Services: map[string]*Service{
			"a": &Service{
				Name: "a",
				Ports: []*Port{
					&Port{
						Container: "api",
						Number:    80,
						Protocol:  "http",
					},
				},
				Containers: map[string]*Container{
					"a": &Container{
						Name:  "a",
						Image: "Image",
					},
				},
			},
		},
	}
	err := validateStack(s)
	require.Error(t, err)
}

func TestValidateStackServiceWithoutContainers(t *testing.T) {
	s := &Stack{
		Name: "test",
		Services: map[string]*Service{
			"a": &Service{
				Name: "a",
			},
		},
	}
	err := validateStack(s)
	require.Error(t, err)
}

func TestValidateLinksOK(t *testing.T) {
	s := &Stack{
		Name: "test",
		Services: map[string]*Service{
			"a": &Service{
				Name: "a",
				Links: []*Link{
					&Link{
						Service: "b",
						Alias:   "b",
					},
					&Link{
						Service: "b",
						Alias:   "b",
					},
				},
				Containers: map[string]*Container{
					"a": &Container{
						Name:  "a",
						Image: "Image",
					},
				},
			},
			"b": &Service{
				Name: "b",
				Containers: map[string]*Container{
					"b": &Container{
						Name:  "b",
						Image: "Image",
					},
				},
			},
			"c": &Service{
				Name: "c",
				Containers: map[string]*Container{
					"c": &Container{
						Name:  "c",
						Image: "Image",
					},
				},
			},
		},
	}
	err := validateLinks(s.Services["a"], s)
	require.NoError(t, err)
}

func TestValidateLinkMissingLink(t *testing.T) {
	s := &Stack{
		Name: "test",
		Services: map[string]*Service{
			"a": &Service{
				Name: "a",
				Links: []*Link{
					&Link{
						Service: "b",
						Alias:   "b",
					},
				},
				Containers: map[string]*Container{
					"a": &Container{
						Name:  "a",
						Image: "Image",
					},
				},
			},
		},
	}
	err := validateLinks(s.Services["a"], s)
	require.Error(t, err)
}

func TestValidatePortsOK(t *testing.T) {
	service := &Service{
		Name: "test",
		Ports: []*Port{
			&Port{
				Container: "api",
				Number:    80,
				Protocol:  "http",
			},
			&Port{
				Container: "db",
				Number:    5432,
				Protocol:  "tcp",
			},
		},
		Containers: map[string]*Container{
			"api": &Container{
				Name:  "api",
				Image: "Image",
			},
			"db": &Container{
				Name:  "db",
				Image: "Image",
			},
		},
	}
	err := validatePorts(service)
	require.NoError(t, err)
}

func TestValidatePortsInvalidContainer(t *testing.T) {
	service := &Service{
		Name: "test",
		Ports: []*Port{
			&Port{
				Container: "api2",
				Number:    80,
				Protocol:  "http",
			},
		},
		Containers: map[string]*Container{
			"api": &Container{
				Name:  "api",
				Image: "Image",
			},
		},
	}
	err := validatePorts(service)
	require.Error(t, err)
}

func TestValidatePortsInvalidProtocol(t *testing.T) {
	service := &Service{
		Name: "test",
		Ports: []*Port{
			&Port{
				Container: "api",
				Number:    80,
				Protocol:  "http2",
			},
		},
		Containers: map[string]*Container{
			"api": &Container{
				Name:  "api",
				Image: "Image",
			},
		},
	}
	err := validatePorts(service)
	require.Error(t, err)
}
