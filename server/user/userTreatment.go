package user

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"log"
)

type User struct {
	login    string
	password string
	mail     string
}

func connecxion() (*bolt.DB, error) {
	db, err := bolt.Open("user.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func deconnecxion(db *bolt.DB) {
	db.Close()
}

func addUser(db *bolt.DB, u User) {
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return err
		}
		encoded, err := json.Marshal(u)
		if err != nil {
			return err
		}
		return b.Put([]byte(u.login), encoded)
	})
}

// func main() {
//
// 	u := &User{
// 		login:    "azerty",
// 		password: "azerty1",
// 		mail:     "azerty@aol.com"}
//
// 	db, _ := connecxion()
// 	addUser(db, *u)
// 	deconnecxion(db)
//
// }
