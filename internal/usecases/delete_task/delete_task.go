package delete_task

type TaskStorage interface {
	DeleteTask(id int) error
}

func DeleteTask(storage TaskStorage, id int) error {
	return storage.DeleteTask(id)
}
