package persistence

import "task/data"

type ITaskDbContext interface {
	GetAll() ([]data.Task, *data.ErrorDetail)
	Add(emp *data.Task) (*data.Task, *data.ErrorDetail)
}
