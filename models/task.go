package models

import (
	"fmt"
)

var DefaultTaskList *TaskManager

type Task struct {
	ID    int64  // Unique identifier
	Title string // Description
	Done  bool   // Is this task done?
}

// NewTask creates a new task given a title, that can't be empty.
func NewTask(title string) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("empty title")
	}
	return &Task{0, title, false}, nil
}

// TaskManager manages a list of tasks in memory.
type TaskManager struct {
	lastID int64
}

// NewTaskManager returns an empty TaskManager.
func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

// Save saves the given Task in the TaskManager.
func (m *TaskManager) Save(task *Task) error {
	if task.ID == 0 {
		m.lastID++
		task.ID = m.lastID
		if err := getDB().Insert(task); err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}

	taskS := []Task{}
	if err := getDB().FindAll(&taskS); err != nil {
		fmt.Println(err)
		return err
	}
	for _, t := range taskS {
		if t.ID == task.ID {
			if err := getDB().UpdateOne(t, task); err != nil {
				fmt.Println(err)
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("unknown task")
}

// All returns the list of all the Tasks in the TaskManager.
func (m *TaskManager) All() []Task {
	taskS := []Task{}
	if err := getDB().FindAll(&taskS); err != nil {
		fmt.Println(err)
		return nil
	}
	return taskS
}

// Find returns the Task with the given id in the TaskManager and a boolean
// indicating if the id was found.
func (m *TaskManager) Find(ID int64) (*Task, bool) {
	taskS := []Task{}
	if err := getDB().FindAll(&taskS); err != nil {
		fmt.Println(err)
		return nil, false
	}
	for _, t := range taskS {
		if t.ID == ID {
			return &t, true
		}
	}
	return nil, false
}

func init() {
	DefaultTaskList = NewTaskManager()
}
