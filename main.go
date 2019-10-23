	package main

	import (
		"fmt"
		"time"
		"web-room/lib/controller"
		"web-room/lib/datasource"
		wrTypes "web-room/lib/types"
	)

	func main() {
		ds := datasource.Connect()
		if ds != nil {
			// initializing room controller
			roomController := controller.RoomController{}
			var room = wrTypes.Room{
				Id: "000",
				Name: "room000001",
				CreatedAt:time.Now(),
				Owner: "miguel aliaga",
			} // creating room object
			createErr := roomController.Create(room, ds)
			if createErr != nil {
				panic(createErr.Error())
			}
		} else {
			fmt.Print("error al conectarse")
		}
	}
