package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/henryinfanteg/heyadivinen_backgo/api-palabras/config"
	categoryService "github.com/henryinfanteg/heyadivinen_backgo/api-words/services/category"
)

func main() {
	// Cargamos la configuracion inicial
	config.LoadConfigFile()
	var categories, errCategory = getCategories()
	if errCategory != nil {
		fmt.Println("ERROR -> ", errCategory)
	}

	var categoryRepository = categoryService.CategoryRepository{}
	if errCategory = categoryRepository.CreateMany(categories); errCategory != nil {
		fmt.Println("ERROR -> ", errCategory)
	} else {
		fmt.Println("CATEGORIAS cargados exitosamente!")
	}

}

func getCategories() ([]categoryService.Category, error) {
	objsJSON, err := ioutil.ReadFile("jsons/categories.json")

	if err != nil {
		// fmt.Println("ERROR ->", err)
		return nil, err
	}

	var objs []categoryService.Category
	err = json.Unmarshal(objsJSON, &objs)

	if err != nil {
		// fmt.Println("ERROR ->", err)
		return nil, err
	}

	for i := range objs {
		objs[i].ID = bson.NewObjectId()
		objs[i].CreationDate = time.Now()
		objs[i].CreationUser = "admin"
		objs[i].DateModify = time.Time{}
	}

	return objs, nil
}
