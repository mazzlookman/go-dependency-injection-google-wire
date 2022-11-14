package google_wire

import "errors"

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{
		Error: true,
	}
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(simpleRepository *SimpleRepository) (*SimpleService, error) {
	if simpleRepository.Error {
		return nil, errors.New("error create service")
	} else {
		return &SimpleService{
			SimpleRepository: simpleRepository,
		}, nil
	}
}
