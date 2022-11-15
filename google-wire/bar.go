package google_wire

type BarRepository struct {
}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type Bar string

type BarService struct {
	*BarRepository
	Name Bar
}

func NewBarService(barRepository *BarRepository, name Bar) *BarService {
	return &BarService{
		BarRepository: barRepository,
		Name:          name,
	}
}

type BarStruct struct {
	Name string
}
