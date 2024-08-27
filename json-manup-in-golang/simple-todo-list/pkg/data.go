package data

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// var todos []Todo
var Todos = make(map[int]Todo)

// Create adds a new Tod to the map
func Create(task Todo) {
	Todos[task.ID] = task
}

// Read retrives a Todo by Id from the Map
func Read(id int) (Todo, bool) {
	todo, exists := Todos[id]
	return todo, exists
}

// Update, updates a Todo in the map
func Update(data Todo) {
	Todos[data.ID] = data
}

func Delete(id int) {
	delete(Todos, id)
}
