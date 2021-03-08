package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/henryinfanteg/heyadivinen_backgo/tree/master/tree/master/api-palabras/config"
	categoriaService "github.com/henryinfanteg/heyadivinen_backgo/tree/master/tree/master/api-palabras/services/categoria"
	palabraService "github.com/henryinfanteg/heyadivinen_backgo/tree/master/tree/master/api-palabras/services/palabra"
)

func main() {
	// Cargamos la configuracion inicial
	config.LoadConfigFile()

	var categorias, errCategoria = getCategorias()
	if errCategoria != nil {
		fmt.Println("ERROR -> ", errCategoria)
	}
	// fmt.Println("CATEGORIAS ->", categorias)
	var categoriaRepository = categoriaService.CategoriaRepository{}
	if errCategoria = categoriaRepository.CreateMany(categorias); errCategoria != nil {
		fmt.Println("ERROR -> ", errCategoria)
	} else {
		fmt.Println("CATEGORIAS cargados exitosamente!")
	}

	var palabras, errPalabra = getPalabras()
	if errPalabra != nil {
		fmt.Println("ERROR -> ", errPalabra)
	}

	var palabraRepository = palabraService.PalabraRepository{}
	if errPalabra = palabraRepository.CreateMany(palabras); errPalabra != nil {
		fmt.Println("ERROR -> ", errPalabra)
	} else {
		fmt.Println("PALABRAS cargados exitosamente!")
	}
}

func getCategorias() ([]categoriaService.Categoria, error) {
	objsJSON, err := ioutil.ReadFile("jsons/categorias.json")

	if err != nil {
		// fmt.Println("ERROR ->", err)
		return nil, err
	}

	var objs []categoriaService.Categoria
	err = json.Unmarshal(objsJSON, &objs)

	if err != nil {
		// fmt.Println("ERROR ->", err)
		return nil, err
	}

	for i := range objs {
		objs[i].ID = bson.NewObjectId()
		objs[i].FechaCreacion = time.Now()
		objs[i].UsuarioCreacion = "admin"
		objs[i].FechaModificacion = time.Time{}
	}

	return objs, nil
}

func getPalabras() ([]palabraService.Palabra, error) {
	objsJSON, err := ioutil.ReadFile("jsons/palabras.json")

	if err != nil {
		// fmt.Println("ERROR ->", err)
		return nil, err
	}

	var objs []palabraService.Palabra
	err = json.Unmarshal(objsJSON, &objs)

	if err != nil {
		// fmt.Println("ERROR ->", err)
		return nil, err
	}

	for i := range objs {
		objs[i].ID = bson.NewObjectId()
		objs[i].FechaCreacion = time.Now()
		objs[i].UsuarioCreacion = "admin"
		objs[i].FechaModificacion = time.Time{}
	}

	return objs, nil
}
