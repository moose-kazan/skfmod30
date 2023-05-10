package storage

type TaskLabel struct {
	TaskId  int
	LabelId int
}

type DBPoolTaskLabelIface interface {
	TaskLabelAssign(task_id int, label_id int) bool
	TaskLabelAssignedForLabel(label_id int) []int
	TaskLabelAssignedForTask(task_id int) []int
	TaskLabelRemove(task_id int, label_id int) bool
}

// TODO: Disable duplicate entries!
func (d *DBPool) TaskLabelAssign(task_id int, label_id int) bool {
	err := d.db.QueryRow(d.ctx, "INSERT INTO tasks_labels (task_id, label_id) VALUES ($1, $2) RETURNING id", task_id, label_id).Scan(&task_id)
	if err != nil {
		return false
	}
	return true
}

func (d *DBPool) TaskLabelAssignedForLabel(label_id int) []int {
	rv := make([]int, 0)
	r, err := d.db.Query(d.ctx, "SELECT task_id FROM tasks_labels WHERE label_id = $1", label_id)
	if err != nil {
		return rv
	}

	for r.Next() {
		var task_id int
		err = r.Scan(&task_id)
		if err != nil {
			return rv
		}
		rv = append(rv, task_id)
	}
	return rv
}

func (d *DBPool) TaskLabelAssignedForTask(task_id int) []int {
	rv := make([]int, 0)
	r, err := d.db.Query(d.ctx, "SELECT label_id FROM tasks_labels WHERE task_id = $1", task_id)
	if err != nil {
		return rv
	}

	for r.Next() {
		var label_id int
		err = r.Scan(&label_id)
		if err != nil {
			return rv
		}
		rv = append(rv, label_id)
	}
	return rv
}

func (d *DBPool) TaskLabelRemove(task_id int, label_id int) bool {
	r, err := d.db.Query(d.ctx, "DELETE FROM tasks_labels WHERE task_id = $1 AND label_id = $2 RETURNING task_id, label_id", task_id, label_id)
	if err != nil {
		return false
	}
	if r.Next() {
		return true
	}
	return false
}
