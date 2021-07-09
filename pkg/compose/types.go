package compose

import "encoding/json"

type StackResource struct {
	LogicalID string
	Type      string
	ARN       string
	Status    string
}

type LoadBalancer struct {
	URL           string
	TargetPort    int
	PublishedPort int
	Protocol      string
}

type ServiceStatus struct {
	ID            string
	Name          string
	Replicas      int
	Desired       int
	Ports         []string
	LoadBalancers []LoadBalancer
}

const (
	StackCreate = iota
	StackUpdate
	StackDelete
)

type LogConsumer interface {
	Log(service, container, message string)
}

type Secret struct {
	ID          string            `json:"ID"`
	Name        string            `json:"Name"`
	Labels      map[string]string `json:"Labels"`
	Description string            `json:"Description"`
	username    string
	password    string
}

func NewSecret(name, username, password, description string) Secret {
	return Secret{
		Name:        name,
		username:    username,
		password:    password,
		Description: description,
	}
}

func (s Secret) ToJSON() (string, error) {
	b, err := json.MarshalIndent(&s, "", "\t")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (s Secret) GetCredString() (string, error) {
	creds := map[string]string{
		"username": s.username,
		"password": s.password,
	}
	b, err := json.Marshal(&creds)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
