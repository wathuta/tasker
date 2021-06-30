package db

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
)

var TaskBucket = []byte("tasks")

var Db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func CreateTask(task string) (int, error) {
	var id int
	err := Db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket(TaskBucket)
		id64, _ := bucket.NextSequence()
		id = int(id64)
		key := ItoB(id)
		return bucket.Put(key, []byte(task))

	})
	if err != nil {
		return -1, err
	}

	return id, nil
}
func ItoB(val int) []byte {
	byteId := make([]byte, 8)
	binary.BigEndian.PutUint64(byteId, uint64(val))
	return byteId
}
func BtoI(val []byte) int {
	return int(binary.BigEndian.Uint64(val))
}

func AllTasksFromDB() ([]Task, error) {
	var tasks []Task
	err := Db.View(func(t *bolt.Tx) error {
		b := t.Bucket(TaskBucket)
		cursor := b.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			tasks = append(tasks, Task{
				Key:   BtoI(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func DeleteTask(key int) error {
	err := Db.Update(func(t *bolt.Tx) error {
		b := t.Bucket(TaskBucket)
		return b.Delete(ItoB(key))

	})
	if err != nil {
		return err
	}
	return nil
}
