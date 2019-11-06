package basic

import (
	"github.com/jacksonyoudi/microservice-in-micro/part1/user-srv/basic/config"
	"github.com/jacksonyoudi/microservice-in-micro/part1/user-srv/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}