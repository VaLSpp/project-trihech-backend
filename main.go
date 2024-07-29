package main

import (
    "project-trihech-backend/config"
    "project-trihech-backend/routes"
    "project-trihech-backend/database"
)

func main() {
    r := routes.SetupRouter()
    config.ConnectDatabase()
    database.Migrate()
    r.Run(":8080")
}
