package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
	DB     *mongo.Database
}

type DBOption struct {
	Host     string
	Port     int64
	Database string
	Username string
	Password string
	SRV      bool // เพิ่ม field นี้เพื่อรองรับ SRV
}

func Register(conn **MongoDB, conf *DBOption) {
	var uri string
	if conf.SRV {
		// ใช้ SRV connection string
		uri = fmt.Sprintf("mongodb+srv://%s:%s@%s/%s",
			conf.Username, conf.Password, conf.Host, conf.Database)
	} else {
		// ใช้ standard connection string
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
			conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	*conn = &MongoDB{
		Client: client,
		DB:     client.Database(conf.Database),
	}
}
