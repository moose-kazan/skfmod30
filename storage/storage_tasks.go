package storage

type Task struct {
	Id         int
	Opened     int64
	Closed     int64
	AuthorId   int
	AssignedId int
	Title      string
	Content    string
}

type DBPoolTaskIface interface {
	TaskAdd(t Task) *Task
	TaskFindById(id int) *Task
	TaskClose(id int) bool
	TaskDeleteById(id int) bool
	TaskList() []*Task
}

func (d *DBPool) TaskAdd(t Task) *Task {
	var id int
	var opened int64
	err := d.db.QueryRow(
		d.ctx,
		`INSERT INTO tasks (
			author_id,assigned_id,title,content)
		VALUES ($1,$2,$3,$4) RETURNING id,opened`,
		t.AuthorId,
		t.AssignedId,
		t.Title,
		t.Content,
	).Scan(&id, &opened)
	if err != nil {
		return nil
	}
	t.Id = id
	t.Opened = opened
	return &t
}

func (d *DBPool) TaskFindById(id int) *Task {
	r, err := d.db.Query(d.ctx, "SELECT id,opened,closed,author_id,assigned_id,title,content FROM tasks WHERE id = $1", id)
	if err != nil {
		return nil
	}

	if r.Next() {
		var t Task
		err = r.Scan(
			&t.Id,
			&t.Opened,
			&t.Closed,
			&t.AuthorId,
			&t.AssignedId,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil
		}
		return &t
	}

	return nil
}

func (d *DBPool) TaskClose(id int) bool {
	r, err := d.db.Query(d.ctx, "UPDATE tasks SET closed = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 RETURNING id", id)
	if err != nil {
		return false
	}
	if r.Next() {
		return true
	}
	return false
}

func (d *DBPool) TaskDeleteById(id int) bool {
	r, err := d.db.Query(d.ctx, "DELETE FROM tasks WHERE id = $1 RETURNING id", id)
	if err != nil {
		return false
	}
	if r.Next() {
		return true
	}
	return false
}

func (d *DBPool) TaskList() []*Task {
	rv := make([]*Task, 0)
	r, err := d.db.Query(d.ctx, "SELECT id,opened,closed,author_id,assigned_id,title,content FROM tasks")
	if err != nil {
		return rv
	}

	for r.Next() {
		var t Task
		err = r.Scan(
			&t.Id,
			&t.Opened,
			&t.Closed,
			&t.AuthorId,
			&t.AssignedId,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return rv
		}
		rv = append(rv, &t)
	}

	return rv
}
