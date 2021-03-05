package model

import (
	db "goweb/app/share/database"
)

func Setup() {
	db.Connect()
}
