package database

import (
	"github.com/couchbase/gocb/v2"
	"time"
)

func GetCollection() *gocb.Collection {
	cluster, err := gocb.Connect(
		"localhost",
		gocb.ClusterOptions{
			Username: "Administrator",
			Password: "password",
		})
	if err != nil {
		panic(err)
	}
	bucket := cluster.Bucket("users")
	// We wait until the bucket is definitely connected and setup.
	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		panic(err)
	}
	collection := bucket.DefaultCollection()
	return collection
}
