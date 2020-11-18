package commons

type ClientProxy struct {
	Host     string
	Port     int
	Protocol Protocol

	ID       int
	TypeName string
}
