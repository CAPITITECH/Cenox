package database

import (
	"cenox/models"
	"sync"
)

var (
	Users  []models.User
	mutex  = &sync.Mutex{}
	nextID int64 = 1
)

func GetNextID() int64 {
	mutex.Lock()
	defer mutex.Unlock()
	id := nextID
	nextID++
	return id
}