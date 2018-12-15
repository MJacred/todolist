package todolist

type MemoryStore struct {
	TodoSingles []*TodoSingle
	TodoRepeats []*TodoRepeat
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (m *MemoryStore) Initialize() {}

func (m *MemoryStore) Load() ([]*TodoSingle, []*TodoRepeat, error) {
	return m.TodoSingles, m.TodoRepeats, nil
}

func (m *MemoryStore) Save(todoSingles []*TodoSingle, todoRepeats []*TodoRepeat) {
	m.TodoSingles = todoSingles
	m.TodoRepeats = todoRepeats
}
