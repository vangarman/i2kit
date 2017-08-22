package stack

const (
	//HTTP protocol
	HTTP = "http"
	//TCP protocol
	TCP = "tcp"
	//UDP protocol
	UDP = "udp"
)

//Stack represents a i2kit.yml file
type Stack struct {
	Name     string
	Services map[string]*Service
}

//Service represents a service in a i2kit.yml file
type Service struct {
	Name       string
	Size       string
	Min        int
	Max        int
	Ports      []*Port
	Links      []*Link
	Containers map[string]*Container
}

//Link represents a container link
type Link struct {
	Service string
	Alias   string
}

//Port represents a container port
type Port struct {
	Container string
	Number    int
	Protocol  string
}

//Container represents a container in a i2kit.yml file
type Container struct {
	Name  string
	Image string
}

//Read returns a stack structure given a path to i2kit.yml file
func Read(name, path string) (*Stack, error) {
	sYml, err := readYml(path)
	if err != nil {
		return nil, err
	}
	s, err := createStack(name, sYml)
	if err != nil {
		return nil, err
	}
	err = validateStack(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
