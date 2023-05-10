package main

import (
	"fmt"
	"skftasks/storage"
)

var (
	dsn string = "postgres://moose@192.168.12.253:5432/skftasks"
)

func main() {
	DBP := storage.New(dsn)
	DBP.Reinit()

	uinfo1 := DBP.UserAdd("Vasya")
	uinfo2 := DBP.UserAdd("Petya")

	fmt.Println("===== Users =====")
	fmt.Println("Rename Vasya to Uasya:", DBP.UserNewName(uinfo1.Id, "Uasya"))
	fmt.Println("Delete Petya:", DBP.UserDeleteById(uinfo2.Id))
	fmt.Println("Delete Petya (already deleted):", DBP.UserDeleteById(uinfo2.Id))

	fmt.Println("New name of Vasya:", DBP.UserFindById(uinfo1.Id).Name)

	fmt.Println("===== Tasks =====")
	tinfo1 := DBP.TaskAdd(storage.Task{
		Title:      "First task",
		Content:    "Do something...",
		AuthorId:   uinfo1.Id,
		AssignedId: 0,
	})
	tinfo2 := DBP.TaskAdd(storage.Task{
		Title:      "Second task",
		Content:    "Do something else...",
		AuthorId:   uinfo1.Id,
		AssignedId: 0,
	})
	fmt.Println("Close first task:", DBP.TaskClose(tinfo1.Id))
	fmt.Println("Remove second task:", DBP.TaskDeleteById(tinfo2.Id))
	fmt.Println("Remove second task (already removed):", DBP.TaskDeleteById(tinfo2.Id))
	fmt.Println("Info about first task:", DBP.TaskFindById(tinfo1.Id))

	fmt.Println("===== Labels =====")
	DBP.LabelAdd("Label One")
	DBP.LabelAdd("Label Two")
	DBP.LabelAdd("Label three")
	labels := DBP.LabelList()
	fmt.Println("Labels:", labels)
}
