package room

import (
	"github.com/boltdb/bolt"
	"log"
	"projet/server/message"
)

func connecxion() (*bolt.DB, error) {
	db, err := bolt.Open("conv.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func deconnecxion(db *bolt.DB) {
	db.Close()
}

func addConv(db *bolt.DB, m message.SendMessage) {
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("convs"))
		if err != nil {
			return err
		}
		encoded := m.ToString()
		return b.Put([]byte(m.Time), []byte(encoded))
	})
}

func getConv(db *bolt.DB, cle string) (m message.SendMessage) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("convs"))
		v := b.Get([]byte(cle))
		message.GetMessageObject(string(v[:]))
		return nil
	})
	return
}
