package controller

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	wr "web-room/lib/types"
)
type RoomController struct {}

func (r RoomController) Create(room wr.Room, driver *sql.DB) error{
	query := "insert into room(id, name, owner, createdAt) values (uuid(),?,?,?)"
	created, err := driver.Prepare(query)
	if err != nil {
		panic(err.Error())
		return  err
	}
	created.Exec(room.Name, room.Owner, time.Now())
	fmt.Fprintf(os.Stdout, "{message: %s}", 12)
	return nil
}

func (r RoomController) Find (room string, driver *sql.DB) ([]wr.Room, error) {
	var query = "select id, name, owner, createdAt from room"
	var response  []wr.Room
	queryResponse, err := driver.Query(query);
	if err != nil {
		panic(err.Error())
		return nil, err
	}
	for queryResponse.Next() {
		room := wr.Room{}
		var id, name, owner string
		var createdAt time.Time
		err := queryResponse.Scan(&id, &name, &owner, &createdAt)
		if err != nil {
			panic(err.Error())
		}
		room.Id = id
		room.Name = name
		room.CreatedAt = createdAt
		room.Owner = owner
		response = append(response, room)
	}
	return response, err
}

func (r RoomController) Delete(personId string) {

}

func (r RoomController) Update(personId string, p wr.Room) {

}


