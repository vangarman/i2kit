package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadYmlOK(t *testing.T) {
	_, err := readYml("./fixtures/ok.yml")
	require.NoError(t, err)
}

func TestReadYmlNonExistingFile(t *testing.T) {
	_, err := readYml("./fixtures/doesnotexist.yml")
	require.Error(t, err)
}

func TestReadYmlInvalidFile(t *testing.T) {
	_, err := readYml("./fixtures/invalid")
	require.Error(t, err)
}

func TestCreateStackOK(t *testing.T) {
	sYml, err := readYml("./fixtures/ok.yml")
	require.NoError(t, err)
	s, err := createStack("test", sYml)
	require.NoError(t, err)
	if s.Name != "test" {
		t.Fatalf("Wrong stack name: %s", s.Name)
	}
	if len(s.Services) != 2 {
		t.Fatalf("Wrong number of services: %d", len(s.Services))
	}
	if s.Services["api"].Name != "api" {
		t.Fatalf("Wrong api name: %s", s.Services["api"].Name)
	}
	if s.Services["api"].Size != "t2.small" {
		t.Fatalf("Wrong api size: %s", s.Services["api"].Size)
	}
	if s.Services["api"].Min != 2 {
		t.Fatalf("Wrong api min: %d", s.Services["api"].Min)
	}
	if s.Services["api"].Max != 5 {
		t.Fatalf("Wrong api max: %d", s.Services["api"].Max)
	}
	if len(s.Services["api"].Links) != 1 {
		t.Fatalf("Wrong api number of links: %d", len(s.Services["api"].Links))
	}
	if s.Services["api"].Links[0].Alias != "db" {
		t.Fatalf("Wrong api link alias: %s", s.Services["api"].Links[0].Alias)
	}
	if s.Services["api"].Links[0].Service != "db" {
		t.Fatalf("Wrong api link service: %s", s.Services["api"].Links[0].Service)
	}
	if len(s.Services["api"].Ports) != 1 {
		t.Fatalf("Wrong api number of ports: %d", len(s.Services["api"].Ports))
	}
	if s.Services["api"].Ports[0].Container != "api" {
		t.Fatalf("Wrong api port container: %s", s.Services["api"].Ports[0].Container)
	}
	if s.Services["api"].Ports[0].Number != 90 {
		t.Fatalf("Wrong api port number: %d", s.Services["api"].Ports[0].Number)
	}
	if s.Services["api"].Ports[0].Protocol != "http" {
		t.Fatalf("Wrong api port protocol: %s", s.Services["api"].Ports[0].Protocol)
	}
	if len(s.Services["api"].Containers) != 1 {
		t.Fatalf("Wrong api number of containers: %d", len(s.Services["api"].Containers))
	}
	if s.Services["api"].Containers["api"].Name != "api" {
		t.Fatalf("Wrong api container api name: %s", s.Services["api"].Containers["api"].Name)
	}
	if s.Services["api"].Containers["api"].Image != "test" {
		t.Fatalf("Wrong api container api image: %s", s.Services["api"].Containers["api"].Image)
	}
	if s.Services["db"].Name != "db" {
		t.Fatalf("Wrong db name: %s", s.Services["db"].Name)
	}
	if s.Services["db"].Size != "t2.small" {
		t.Fatalf("Wrong db size: %s", s.Services["db"].Size)
	}
	if s.Services["db"].Min != 1 {
		t.Fatalf("Wrong db min: %d", s.Services["db"].Min)
	}
	if s.Services["db"].Max != 1 {
		t.Fatalf("Wrong db max: %d", s.Services["db"].Max)
	}
	if len(s.Services["db"].Links) != 0 {
		t.Fatalf("Wrong db number of links: %d", len(s.Services["db"].Links))
	}
	if len(s.Services["db"].Ports) != 1 {
		t.Fatalf("Wrong db number of ports: %d", len(s.Services["db"].Ports))
	}
	if s.Services["db"].Ports[0].Container != "db" {
		t.Fatalf("Wrong db port container: %s", s.Services["db"].Ports[0].Container)
	}
	if s.Services["db"].Ports[0].Number != 5432 {
		t.Fatalf("Wrong db port number: %d", s.Services["db"].Ports[0].Number)
	}
	if s.Services["db"].Ports[0].Protocol != "tcp" {
		t.Fatalf("Wrong db port protocol: %s", s.Services["db"].Ports[0].Protocol)
	}
	if len(s.Services["db"].Containers) != 1 {
		t.Fatalf("Wrong db number of containers: %d", len(s.Services["db"].Containers))
	}
	if s.Services["db"].Containers["db"].Name != "db" {
		t.Fatalf("Wrong db container db name: %s", s.Services["db"].Containers["db"].Name)
	}
	if s.Services["db"].Containers["db"].Image != "test" {
		t.Fatalf("Wrong db container db image: %s", s.Services["db"].Containers["db"].Image)
	}
}

func TestCreateStackBadLink(t *testing.T) {
	sYml, err := readYml("./fixtures/bad-link.yml")
	require.NoError(t, err)
	_, err = createStack("test", sYml)
	require.Error(t, err)
}

func TestCreateStackBadPort(t *testing.T) {
	sYml, err := readYml("./fixtures/bad-port.yml")
	require.NoError(t, err)
	_, err = createStack("test", sYml)
	require.Error(t, err)
}
