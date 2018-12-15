package todolist

import "time"

// Timestamp format to include date, time with timezone support. Easy to parse
const ISO8601_TIMESTAMP_FORMAT = "2006-01-02T15:04:05Z07:00"

type Todo struct {
	Id         int      `json:"id"`
	Type       byte     `json:"id"`
	Subject    string   `json:"subject"`
	Projects   []string `json:"projects"`
	Contexts   []string `json:"contexts"`
	Due        string   `json:"due"`
	Archived   bool     `json:"archived"`
	IsPriority bool     `json:"isPriority"`
	Notes      []string `json:"notes"`
}

type TodoSingle struct {
	Todo
	Completed     bool   `json:"completed"`
	CompletedDate string `json:"completedDate"`
}

type TodoRepeat struct {
	Todo
	RepeatStart    string   `json:"repeat-start"`
	RepeatType     string   `json:"repeat-type"`
	CompletedTodos []bool   `json:"completedTodos"`
	CompletedDates []string `json:"completedDates"`
}

func NewTodoSingle() *TodoSingle {
	return &TodoSingle{Type: 0, Completed: false, Archived: false, IsPriority: false}
}

func NewTodoRepeat() *TodoRepeat {
	return &TodoRepeat{Type: 1, CompletedTodos: {false, false, false, false, false, false, false}, Archived: false, IsPriority: false}
}

func (t Todo) Valid() bool {
	return (t.Subject != "")
}

func (t Todo) CalculateDueTime() time.Time {
	if t.Due != "" {
		parsedTime, _ := time.Parse("2006-01-02", t.Due)
		return parsedTime
	} else {
		parsedTime, _ := time.Parse("2006-01-02", "1900-01-01")
		return parsedTime
	}
}

func (t *TodoSingle) Complete() {
	t.Completed = true
	t.CompletedDate = timestamp(time.Now()).Format(ISO8601_TIMESTAMP_FORMAT)
}

func (t *TodoSingle) Uncomplete() {
	t.Completed = false
	t.CompletedDate = ""
}

func (t *Todo) Archive() {
	t.Archived = true
}

func (t *TodoSingle) Unarchive() {
	t.Archived = false
}

func (t *TodoSingle) Prioritize() {
	t.IsPriority = true
}

func (t *TodoSingle) Unprioritize() {
	t.IsPriority = false
}

func (t Todo) CompletedDateToDate() string {
	parsedTime, _ := time.Parse(ISO8601_TIMESTAMP_FORMAT, t.CompletedDate)
	return parsedTime.Format("2006-01-02")
}

func (t *TodoRepeat) Complete() {
}

func (t *TodoRepeat) Uncomplete() {

}

func (t *TodoRepeat) Archive() {

}

func (t *TodoRepeat) Unarchive() {

}

func (t *TodoRepeat) Prioritize() {

}

func (t *TodoRepeat) Unprioritize() {
	t.IsPriority = false
}

// func (t TodoRepeat) CompletedDateToDate() string {
// 	return ""
// }
