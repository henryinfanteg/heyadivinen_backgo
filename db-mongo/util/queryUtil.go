package util

import (
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CreateQuery crea un query
func CreateQuery(collection *mgo.Collection, params map[string]interface{}, sort string, pageNumber int, pageSize int) *mgo.Query {

	// fmt.Println("params", params)
	query := collection.Find(params)

	// Ordenamos la consulta
	AddSortAQuery(query, sort)

	// Paginamos la consulta
	AddPaginationAQuery(query, pageNumber, pageSize)

	return query
}

// CreateQuery crea un query
func CreateQueryRandom(collection *mgo.Collection, params map[string]interface{}) *mgo.Pipe {
	pipe := collection.Pipe([]bson.M{{"$match": bson.M{"categoriaId":params["categoriaId"]}}, {"$sample": bson.M{"size":10}} })
	resp := []bson.M{}
	err := pipe.All(&resp)
	if err != nil {
	//handle error
	}
	return pipe
}

// CreateQuerySearch crea un query de busqueda
func CreateQuerySearch(collection *mgo.Collection, querySearch bson.M, sort string, pageNumber int, pageSize int) *mgo.Query {

	// fmt.Println("params", params)
	query := collection.Find(querySearch)

	// Ordenamos la consulta
	AddSortAQuery(query, sort)

	// Paginamos la consulta
	AddPaginationAQuery(query, pageNumber, pageSize)

	return query
}

// CreateQuerySearchWithSelect crea un query de busqueda con select
func CreateQuerySearchWithSelect(collection *mgo.Collection, querySearch bson.M, sort string, pageNumber int, pageSize int, bsonSelect bson.M) *mgo.Query {

	// fmt.Println("params", params)
	query := collection.Find(querySearch).Select(bsonSelect)

	// Ordenamos la consulta
	AddSortAQuery(query, sort)

	// Paginamos la consulta
	AddPaginationAQuery(query, pageNumber, pageSize)

	return query
}

// CreateQueryWithSelect crea un query con select
func CreateQueryWithSelect(collection *mgo.Collection, params map[string]interface{}, sort string, pageNumber int, pageSize int, bsonSelect bson.M) *mgo.Query {

	// fmt.Println("params", params)
	query := collection.Find(params).Select(bsonSelect)

	// Ordenamos la consulta
	AddSortAQuery(query, sort)

	// Paginamos la consulta
	AddPaginationAQuery(query, pageNumber, pageSize)

	return query
}

// AddSortAQuery Ordena la consulta
func AddSortAQuery(query *mgo.Query, sort string) {
	if len(sort) > 0 {
		sortResult := strings.Split(sort, ",")
		for _, obj := range sortResult {
			query = query.Sort(strings.TrimSpace(obj))
		}
	}
}

// AddPaginationAQuery agrega la paginacion a la consulta
func AddPaginationAQuery(query *mgo.Query, pageNumber int, pageSize int) {
	if pageNumber > 0 && pageSize > 0 {
		query = query.Limit(pageSize).Skip((pageNumber - 1) * pageSize)
	}
}
