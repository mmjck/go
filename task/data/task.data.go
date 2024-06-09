package data

var TaskData []Task

func InitEmpData() {
	TaskData = []Task{
		{
			ID:       1,
			Name:     "Tak 1",
			Employee: 1,
		},
		{
			ID:       2,
			Name:     "Task 2",
			Employee: 2,
		},
		{
			ID:       3,
			Name:     "Task 3",
			Employee: 3,
		},
	}
}
