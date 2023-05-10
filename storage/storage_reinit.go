package storage

import (
	"log"
)

var reinitQueries []string = []string{
	`DROP TABLE IF EXISTS tasks_labels, tasks, labels, users;`,
	`CREATE TABLE users ( 
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	);`,
	`CREATE TABLE labels (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	);`,
	`CREATE TABLE tasks (
		id SERIAL PRIMARY KEY,
		opened BIGINT NOT NULL DEFAULT extract(epoch from now()),
		closed BIGINT DEFAULT 0,
		author_id INTEGER REFERENCES users(id) DEFAULT 0,
		assigned_id INTEGER REFERENCES users(id) DEFAULT 0,
		title TEXT,
		content TEXT
	);`,
	`CREATE TABLE tasks_labels (
		task_id INTEGER REFERENCES tasks(id),
		label_id INTEGER REFERENCES labels(id)
	);`,
	`INSERT INTO users (id, name) VALUES (0, 'default');`,
}

type DBPoolReinitIface interface {
	Reinit()
}

func (d *DBPool) Reinit() {
	tx, err := d.db.Begin(d.ctx)
	defer tx.Rollback(d.ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, q := range reinitQueries {
		//fmt.Println(q)
		_, err = tx.Exec(d.ctx, q)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit(d.ctx)
	if err != nil {
		log.Fatal(err)
	}
}
