package room

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"log"
)

func connecxion() (*bolt.DB, error) {
	db, err := bolt.Open("room.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func deconnecxion(db *bolt.DB) {
	db.Close()
}

func addRoom(db *bolt.DB, roomName string) {
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("rooms"))
		if err != nil {
			return err
		}
		encoded, err := json.Marshal(roomName)
		if err != nil {
			return err
		}
		return b.Put([]byte(roomName), encoded)
	})
}

func getRoom(db *bolt.DB, roomName string) (r Room) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("users"))
		v := b.Get([]byte(roomName))
		json.Unmarshal(v, &r)
		return nil
	})
	return
}
