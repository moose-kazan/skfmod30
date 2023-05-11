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
	DBP.UserAdd("Kolya")
	DBP.UserAdd("Tahir")

	users := DBP.UserList()
	fmt.Println("\tID\tName")
	for _, u := range users {
		fmt.Printf("\t%02d\t%s\n", u.Id, u.Name)
	}

	fmt.Println("\n\n===== Tasks =====")
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

	tasks := DBP.TaskList()
	fmt.Println("\tID\tTask")
	for _, t := range tasks {
		fmt.Printf("\t%02d\t%s\n", t.Id, t.Title)
	}

	fmt.Println("\n\n===== Labels =====")
	DBP.LabelAdd("Label One")
	linfo2 := DBP.LabelAdd("Label Two")
	linfo3 := DBP.LabelAdd("Label three")
	labels := DBP.LabelList()
	fmt.Println("\tId\tLabel")
	for _, l := range labels {
		fmt.Printf("\t%02d\t\"%s\"\n", l.Id, l.Name)
	}
	fmt.Println("Deleete first label:", DBP.LabelDelete(labels[0].Id))
	fmt.Println("Deleete first label (already deleted):", DBP.LabelDelete(labels[0].Id))
	labels = DBP.LabelList()
	fmt.Println("\tId\tLabel")
	for _, l := range labels {
		fmt.Printf("\t%02d\t\"%s\"\n", l.Id, l.Name)
	}

	fmt.Println("\n\n===== Tasks Labels =====")
	fmt.Println("Assign second label to first task:",
		DBP.TaskLabelAssign(tinfo1.Id, linfo2.Id),
	)
	fmt.Println("Assign thrid label to first task:",
		DBP.TaskLabelAssign(tinfo1.Id, linfo3.Id),
	)
	fmt.Println("Labels for first task:",
		DBP.TaskLabelAssignedForTask(tinfo1.Id),
	)
	fmt.Println("Tasks for second label:",
		DBP.TaskLabelAssignedForLabel(linfo2.Id),
	)
	fmt.Println("Remove assign for first task and secod label:",
		DBP.TaskLabelRemove(tinfo1.Id, linfo2.Id),
	)
	fmt.Println("Remove assign for first task and secod label (already removed):",
		DBP.TaskLabelRemove(tinfo1.Id, linfo2.Id),
	)

}
