package controller

import (
	"database/sql"
	"time"
	"web-room/lib/types"
)

func Find(personId string, driver *sql.DB) ([]types.Person, error) {
	var query = "select id, fullname, bio, createdAt from person"
	var response  []types.Person
	queryResponse, err := driver.Query(query);
	if err != nil {
		panic(err.Error())
		return nil, err
	}
	for queryResponse.Next() {
		person := types.Person{}
		var id, fullName, bio string
		var createdAt time.Time
		err := queryResponse.Scan(&id, &fullName, &bio, &createdAt)
		if err != nil {
			panic(err.Error())
		}
		person.Id = id
		person.Bio = bio
		person.CreatedAt = createdAt
		person.FullName = fullName
		response = append(response, person)
	}
	return response, err
}
