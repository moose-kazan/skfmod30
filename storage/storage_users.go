package storage

type User struct {
	Id   int
	Name string
}

type DBPoolUserIface interface {
	UserAdd(name string) *User
	UserFindById(id int) *User
	UserNewName(id int, name string) bool
	UserDeleteById(id int) bool
	UserList() []*User
}

func (d *DBPool) UserAdd(name string) *User {
	var id int
	err := d.db.QueryRow(d.ctx, "INSERT INTO users (name) VALUES ($1) RETURNING id", name).Scan(&id)
	if err != nil {
		return nil
	}
	var u User = User{Id: id, Name: name}
	return &u
}

func (d *DBPool) UserFindById(id int) *User {
	r, err := d.db.Query(d.ctx, "SELECT id,name FROM users WHERE id = $1", id)
	if err != nil {
		return nil
	}

	if r.Next() {
		var u User
		err = r.Scan(
			&u.Id,
			&u.Name,
		)
		if err != nil {
			return nil
		}
		return &u
	}

	return nil
}

func (d *DBPool) UserNewName(id int, name string) bool {
	r, err := d.db.Query(d.ctx, "UPDATE users SET name = $2 WHERE id = $1 RETURNING name", id, name)
	if err != nil {
		return false
	}
	if r.Next() {
		return true
	}
	return false
}

func (d *DBPool) UserDeleteById(id int) bool {
	r, err := d.db.Query(d.ctx, "DELETE FROM users WHERE id = $1 RETURNING id", id)
	if err != nil {
		return false
	}
	if r.Next() {
		return true
	}
	return false
}

func (d *DBPool) UserList() []*User {
	rv := make([]*User, 0)
	r, err := d.db.Query(d.ctx, "SELECT id,name FROM users")
	if err != nil {
		return rv
	}

	for r.Next() {
		var u User
		err = r.Scan(
			&u.Id,
			&u.Name,
		)
		if err != nil {
			return rv
		}
		rv = append(rv, &u)
	}

	return rv
}
