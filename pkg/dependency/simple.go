package dependency

import "errors"

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository(isErr bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isErr,
	}
}

type SimpleServices struct {
	*SimpleRepository
}

func NewSimpleServices(simpleRepo *SimpleRepository) (*SimpleServices, error) {
	if simpleRepo.Error {
		return nil, errors.New("Failed create services")
	} else {
		return &SimpleServices{
			SimpleRepository: simpleRepo,
		}, nil
	}
}
