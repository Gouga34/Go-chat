package user

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"log"
)

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
		return b.Put([]byte(u.Login), encoded)
	})
}

func getUser(db *bolt.DB, cle string) (u User) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("users"))
		v := b.Get([]byte(cle))
		json.Unmarshal(v, &u)
		return nil
	})
	return
}

func existUser(db *bolt.DB, cle string) bool {

	var u User
	var v []byte

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("users"))
		v = b.Get([]byte(cle))
		json.Unmarshal(v, &u)
		return nil
	})
	if v != nil {
		return true
	} else {
		return false
	}
}

// func main() {
//
// 	u := &User{
// 		Login:    "azerty",
// 		Password: "azerty1",
// 		Mail:     "azerty@aol.com"}
//
// 	db, _ := connecxion()
// 	addUser(db, *u)
// 	res := getUser(db, "azerty")
//
// 	fmt.Println(res)
//
// 	deconnecxion(db)
//
// }
