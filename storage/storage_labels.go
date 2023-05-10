package storage

type Label struct {
	Id   int
	Name string
}

type DBPoolLabelIface interface {
	LabelAdd(name string) *Label
	LabelDelete(id int) bool
	LabelList() []*Label
}

func (d *DBPool) LabelAdd(name string) *Label {
	var id int
	err := d.db.QueryRow(d.ctx, "INSERT INTO labels (name) VALUES ($1) RETURNING id", name).Scan(&id)
	if err != nil {
		return nil
	}
	var tl Label = Label{Id: id, Name: name}
	return &tl
}

func (d *DBPool) LabelDelete(id int) bool {
	r, err := d.db.Query(d.ctx, "DELETE FROM labels WHERE id = $1 RETURNING id", id)
	if err != nil {
		return false
	}
	if r.Next() {
		return true
	}
	return false
}

func (d *DBPool) LabelList() []*Label {
	rv := make([]*Label, 0)
	r, err := d.db.Query(d.ctx, "SELECT id,name FROM labels")
	if err != nil {
		return rv
	}
	for r.Next() {
		var tl Label
		r.Scan(&tl.Id, &tl.Name)
		rv = append(rv, &tl)
	}
	return rv
}
