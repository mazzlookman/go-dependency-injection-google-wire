package google_wire

type Configuration struct {
	Name string
}

type Application struct {
	*Configuration
}

func NewApplication() *Application {
	return &Application{
		&Configuration{Name: "iamaqibmoh"},
	}
}
