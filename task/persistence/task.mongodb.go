package persistence

import "task/data"

type TaskMongoDb struct {
	connctionString string
	dbname          string
}

func InitMongoDb(connctionString, dbname string) TaskMongoDb {
	data.InitEmpData()

	return TaskMongoDb{
		connctionString: connctionString,
		dbname:          dbname,
	}
}

func (dbContext TaskMongoDb) GetAll() ([]data.Task, *data.ErrorDetail) {
	return data.TaskData, nil
}

func (dbContext TaskMongoDb) Add(task *data.Task) (*data.Task, *data.ErrorDetail) {
	task.ID = len(data.TaskData) + 1
	result := data.TaskData

	result = append(result, *task)
	data.TaskData = result
	return task, nil
}
