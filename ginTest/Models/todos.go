package Models

import (
	Config "../Config"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// get all
func GetAllTodos(todo *[]Todo) (err error) {
    if err := Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

// add todo
func NewTodo(todo *Todo) (err error) {
	if err = Config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

// get one
func GetTodo(todo *Todo, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

// update a todo
func UpdateTodo(todo *Todo, id string) (err error) {
	Config.DB.Save(todo)
	return nil
}

// del todo
func DelTodo(todo *Todo, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(todo)
	return nil
}

//del all
func DelAll(todo *[]Todo) (err error) {
    Config.DB.DropTable("todo")
    Config.DB.CreateTable(&Todo{})
    return nil
}
