package main

import (
    "myapi/database"
    "myapi/routes"
)

func main() {
    database.ConnectDatabase()
    r := routes.SetupRouter()
    r.Run(":8080")
}
