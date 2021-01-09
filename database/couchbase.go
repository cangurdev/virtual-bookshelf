package database

import (
	"fmt"
	"github.com/couchbase/gocb/v2"
	"time"
)

func GetCollection() *gocb.Collection {
	cluster := GetCluster()
	bucket := cluster.Bucket("users")
	// We wait until the bucket is definitely connected and setup.
	err := bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		fmt.Print(err)
	}
	collection := bucket.DefaultCollection()
	return collection
}
func GetCluster() *gocb.Cluster {
	cluster, err := gocb.Connect(
		"localhost",
		gocb.ClusterOptions{
			Username: "Administrator",
			Password: "password",
		})
	if err != nil {
		fmt.Print(err)
	}
	return cluster
}
