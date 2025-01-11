package add

import (
	"task/data"
	"task/persistence"
)

type Service struct {
	repository persistence.ITaskDbContext
}

func InitService(repo persistence.ITaskDbContext) *Service {
	return &Service{
		repository: repo,
	}
}

func (service *Service) Add(em *data.Task) (*data.Task, *data.ErrorDetail) {
	return service.repository.Add(em)
}
