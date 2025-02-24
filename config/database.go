package config

import (
	pdb "app/app/provider/database"
	"log"
	"strconv"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

func Database() {
	// Connect to database
	srv, _ := strconv.ParseBool(confString("DB_SRV", "false")) // อ่าน DB_SRV, default เป็น false
	pdb.Register(
		&db,
		&pdb.DBOption{
			Host:     confString("DB_HOST", "127.0.0.1"),
			Port:     confInt64("DB_PORT", int64(27017)),
			Database: confString("DB_DATABASE", "Database"),
			Username: confString("DB_USER", ""),
			Password: confString("DB_PASSWORD", ""),
			SRV:      srv, // เพิ่ม SRV เข้าไปใน DBOption
		},
	)
	log.Println("database connected success")
}

var (
	db     *pdb.MongoDB
	dbMap  = make(map[string]*pdb.MongoDB)
	dbLock sync.RWMutex
)

func GetDB() *mongo.Database {
	return db.DB
}

func DB(name ...string) *pdb.MongoDB {
	dbLock.RLock()
	defer dbLock.RUnlock()
	if dbMap == nil {
		panic("database not initialized")
	}

	if len(name) == 0 {
		return dbMap["default"]
	}

	db, ok := dbMap[name[0]]
	if !ok {
		panic("database not initialized")
	}
	return db
}
