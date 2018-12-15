package todolist

type Store interface {
	Initialize()
	Load() ([]*TodoSingle, []*TodoRepeat, error)
	Save(todoSingles []*TodoSingle, todoRepeats []*TodoRepeat)
}
