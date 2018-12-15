package todolist

import (
	"sort"
	"time"
)

type TodoList struct {
	Data       []*TodoSingle
	DataRepeat []*TodoRepeat
}

func (t *TodoList) Load(todoSingles []*TodoSingle, todoRepeats []*TodoRepeat) {
	t.Data = todoSingles
	t.DataRepeat = todoRepeats
}

func (t *TodoList) AddTodoSingle(todo *TodoSingle) {
	todo.Id = t.NextId()
	t.Data = append(t.Data, todo)
}

func (t *TodoList) AddTodoRepeat(todo *TodoRepeat) {
	todo.Id = t.NextId()
	t.DataRepeat = append(t.DataRepeat, todo)
}

func (t *TodoList) Delete(ids ...int) {
	for _, id := range ids {
		todo := t.FindById(id)
		if todo == nil {
			continue
		}
		i := -1

		if todo.Type == 0 {
			for index, todo := range t.Data {
				if todo.Id == id {
					i = index
				}
			}
			t.Data = append(t.Data[:i], t.Data[i+1:]...)
		} else {
			for index, todo := range t.DataRepeat {
				if todo.Id == id {
					i = index
				}
			}
			t.DataRepeat = append(t.DataRepeat[:i], t.DataRepeat[i+1:]...)
		}
	}
}

func (t *TodoList) CompleteTodoSingles(ids ...int) {
	for _, id := range ids {
		todo := &TodoSingle{Todo: t.FindById(id)}
		if todo == nil {
			continue
		}
		todo.Complete()
		t.Delete(id)
		t.Data = append(t.Data, todo)
	}
}

func (t *TodoList) CompleteTodoRepeats(ids ...int) {
	for _, id := range ids {
		todo := &TodoRepeat{Todo: t.FindById(id)}
		if todo == nil {
			continue
		}
		todo.Complete()
		t.Delete(id)
		t.DataRepeat = append(t.DataRepeat, todo)
	}
}

func (t *TodoList) UncompleteTodoSingles(ids ...int) {
	for _, id := range ids {
		todo := &TodoSingle{Todo: t.FindById(id)}
		if todo == nil {
			continue
		}
		todo.Uncomplete()
		t.Delete(id)
		t.Data = append(t.Data, todo)
	}
}

func (t *TodoList) UncompleteTodoRepeats(ids ...int) {
	for _, id := range ids {
		todo := &TodoRepeat{Todo: t.FindById(id)}
		if todo == nil {
			continue
		}
		todo.Uncomplete()
		t.Delete(id)
		t.DataRepeat = append(t.DataRepeat, todo)
	}
}

func (t *TodoList) Archive(ids ...int) {
	for _, id := range ids {
		// FIXME: either... or
		todo := &TodoSingle{Todo: t.FindById(id)}
		todo := &TodoRepeat{Todo: t.FindById(id)}

		if todo == nil {
			continue
		}
		todo.Archive()
		t.Delete(id)

		// FIXME: either... or
		t.Data = append(t.Data, todo)
		t.DataRepeat = append(t.DataRepeat, todo)
	}
}

func (t *TodoList) Unarchive(ids ...int) {
	for _, id := range ids {
		// FIXME: either... or
		todo := &TodoSingle{Todo: t.FindById(id)}
		todo := &TodoRepeat{Todo: t.FindById(id)}

		if todo == nil {
			continue
		}
		todo.Unarchive()
		t.Delete(id)

		// FIXME: either... or
		t.Data = append(t.Data, todo)
		t.DataRepeat = append(t.DataRepeat, todo)
	}
}

func (t *TodoList) Prioritize(ids ...int) {
	for _, id := range ids {
		// FIXME: either... or
		todo := &TodoSingle{Todo: t.FindById(id)}
		todo := &TodoRepeat{Todo: t.FindById(id)}

		if todo == nil {
			continue
		}
		todo.Prioritize()
		t.Delete(id)

		// FIXME: either... or
		t.Data = append(t.Data, todo)
		t.DataRepeat = append(t.DataRepeat, todo)
	}
}

func (t *TodoList) Unprioritize(ids ...int) {
	for _, id := range ids {
		// FIXME: either... or
		todo := &TodoSingle{Todo: t.FindById(id)}
		todo := &TodoRepeat{Todo: t.FindById(id)}

		if todo == nil {
			continue
		}
		todo.Unprioritize()
		t.Delete(id)

		// FIXME: either... or
		t.Data = append(t.Data, todo)
		t.DataRepeat = append(t.DataRepeat, todo)
	}
}

func (t *TodoList) IndexOf(todoToFind *interface{}, data []interface{}) int {
	for i, todo := range data {
		if todo.Id == todoToFind.Id {
			return i
		}
	}
	return -1
}

type ByDate []*Todo

func (a ByDate) Len() int      { return len(a) }
func (a ByDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool {
	t1Due := a[i].CalculateDueTime()
	t2Due := a[j].CalculateDueTime()
	return t1Due.Before(t2Due)
}

func (t *TodoList) Todos() []*Todo {

	{ // create and insert copies for repeat todo functionality
		sizeForRepeats := 0

		for _, todo := range t.Data {
			if len(todo.RepeatType) > 0 && todo.RepeatType != "undefined" {
				switch todo.RepeatType {
				case "day":
					sizeForRepeats += 5
					break
				case "week", "month-day", "month-weekday":
					sizeForRepeats += 2
					break
				default:
					sizeForRepeats++
				}
			} else {
				sizeForRepeats++
			}
		}

		if sizeForRepeats > len(t.Data) {
			todoList := make([]*Todo, sizeForRepeats)

			counter := 0
			for _, todo := range t.Data {
				if len(todo.RepeatType) > 0 && todo.RepeatType != "undefined" {

					switch todo.RepeatType {
					case "day":
						for index := 0; index < 5; index++ {
							copy := &Todo{}
							*copy = *todo
							copy.Due = time.Now().AddDate(0, 0, index).Format("2006-01-02")
							todoList[counter] = copy
							counter++
						}
						break
					case "week":
						// TODO: check if this week's todo has passed
						// TODO: consider RepeatStart date to get correct day of week
						for index := 0; index < 2; index++ {
							copy := &Todo{}
							*copy = *todo
							copy.Due = time.Now().AddDate(0, 0, 7*index).Format("2006-01-02")
							todoList[counter] = copy
							counter++
						}
						break
					case "month-day":
						// TODO: check if this month-day's todo has passed
						// TODO: consider RepeatStart date to get correct day of month
						for index := 0; index < 2; index++ {
							copy := &Todo{}
							*copy = *todo
							copy.Due = time.Now().AddDate(0, index, 0).Format("2006-01-02")
							todoList[counter] = copy
							counter++
						}
						break
					case "month-weekday":
						// TODO: check if this month-weekday's todo has passed
						// TODO: consider RepeatStart date to get correct weekday
						for index := 0; index < 2; index++ {
							copy := &Todo{}
							*copy = *todo
							copy.Due = time.Now().AddDate(0, index, 0).Format("2006-01-02")
							todoList[counter] = copy
							counter++
						}
						break
					case "year":
						// TODO: check if this year's todo has passed
						copy := &Todo{}
						*copy = *todo
						copy.Due = time.Now().AddDate(1, 0, 0).Format("2006-01-02")
						todoList[counter] = copy
						counter++
						break
						// TODO: default??
					}

				} else {
					todoList[counter] = todo
					counter++
				}
			}

			//  TODO: sort also by number
			sort.Sort(ByDate(todoList))
			return todoList
		}
	}

	sort.Sort(ByDate(t.Data))
	return t.Data
}

func (t *TodoList) MaxId() int {
	maxId := 0
	for _, todo := range t.Data {
		if todo.Id > maxId {
			maxId = todo.Id
		}
	}
	for _, todo := range t.DataRepeat {
		if todo.Id > maxId {
			maxId = todo.Id
		}
	}

	return maxId
}

func (t *TodoList) NextId() int {
	var found bool
	maxID := t.MaxId()
	for i := 1; i <= maxID; i++ {
		found = false
		for _, todo := range t.Data {
			if todo.Id == i {
				found = true
				break
			}
		}
		for _, todo := range t.DataRepeat {
			if todo.Id == i {
				found = true
				break
			}
		}
		if !found {
			return i
		}
	}
	return maxID + 1
}

func (t *TodoList) FindById(id int) *Todo {
	for _, todo := range t.Data {
		if todo.Id == id {
			return todo
		}
	}
	for _, todo := range t.DataRepeat {
		if todo.Id == id {
			return todo
		}
	}
	return nil
}

func (t *TodoList) GarbageCollect() {
	var toDelete []*Todo
	for _, todo := range t.Data {
		if todo.Archived {
			toDelete = append(toDelete, todo)
		}
	}
	for _, todo := range t.DataRepeat {
		if todo.Archived {
			toDelete = append(toDelete, todo)
		}
	}
	for _, todo := range toDelete {
		t.Delete(todo.Id)
	}
}
