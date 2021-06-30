package main

import (
	"log"
	"os"
	"path/filepath"
	"tasker/cmd"
	"tasker/db"
	"time"

	"github.com/boltdb/bolt"
	"github.com/joho/godotenv"
)

func init() {
	f := func(dbString string) error {
		var err error
		db.Db, err = bolt.Open(dbString, 0600, &bolt.Options{Timeout: 5 * time.Second})
		if err != nil {
			return err
		}
		return db.Db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists(db.TaskBucket)
			return err
		})
	}

	err := godotenv.Load("vars.env")
	if err != nil {
		log.Fatal(err)
	}
	//dbString := os.Getenv("DBSTR")

	//geting the home dir
	var homeDir string
	homeDir, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join(homeDir, "task.db")
	if err := f(dbPath); err != nil {
		log.Fatal("error is not nil", err.Error())
	}

}

func main() {
	cmd.RootCMD.Execute()
}
