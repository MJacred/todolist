package todolist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

type FileStore struct {
	FileLocation string
	Loaded       bool
}

func NewFileStore() *FileStore {
	return &FileStore{FileLocation: "", Loaded: false}
}

func (f *FileStore) Initialize() {
	if f.FileLocation == "" {
		f.FileLocation = ".todos.json"
	}

	_, err := ioutil.ReadFile(f.FileLocation)
	if err == nil {
		fmt.Println("It looks like a .todos.json file already exists!  Doing nothing.")
		os.Exit(0)
	}
	if err := ioutil.WriteFile(f.FileLocation, []byte("[]"), 0644); err != nil {
		fmt.Println("Error writing json file", err)
		os.Exit(1)
	}
	fmt.Println("Todo repo initialized.")
}

func (f *FileStore) Load() ([]*TodoSingle, []*TodoRepeat, error) {
	if f.FileLocation == "" {
		f.FileLocation = getLocation()
	}

	data, err := ioutil.ReadFile(f.FileLocation)
	if err != nil {
		fmt.Println("No todo file found!")
		fmt.Println("Initialize a new todo repo by running 'todolist init'")
		os.Exit(0)
		return nil, nil, err
	}

	var todoSingles []*TodoSingle
	jerr := json.Unmarshal(data, &todoSingles)
	if jerr != nil {
		fmt.Println("Error reading json data", jerr)
		os.Exit(1)
		return nil, nil, jerr
	}

	var todoRepeats []*TodoRepeat
	jerr = json.Unmarshal(data, &todoRepeats)
	if jerr != nil {
		fmt.Println("Error reading json data", jerr)
		os.Exit(1)
		return nil, nil, jerr
	}

	f.Loaded = true

	return todoSingles, todoRepeats, nil
}

func (f *FileStore) Save(todoSingles []*TodoSingle, todoRepeats []*TodoRepeat) {
	data, _ := json.Marshal(todoSingles)
	data2, _ := json.Marshal(todoRepeats)
	data = append(data, data2...)
	if err := ioutil.WriteFile(f.FileLocation, []byte(data), 0644); err != nil {
		fmt.Println("Error writing json file", err)
	}
}

func getLocation() string {
	localrepo := ".todos.json"
	usr, _ := user.Current()
	homerepo := fmt.Sprintf("%s/.todos.json", usr.HomeDir)
	_, ferr := os.Stat(localrepo)

	if ferr == nil {
		return localrepo
	} else {
		return homerepo
	}
}
