package db

import (
	"fmt"
	"github.com/boltdb/bolt"
	"projet/server/logger"
)

//DbFile fichier de la BD
const DbFile = "chat.db"

//MessageBucket bucket message
const MessageBucketPrefix = "Messages_"

//UserBucket bucket user
const UserBucket = "user"

//RoomBucket bucket room
const RoomBucket = "room"

//Db Objet global de gestion de la bd
var Db *DbManager

//DbManager manager de la BD
type DbManager struct {
	db *bolt.DB
}

//Init Connexion à la bd et création des buckets
func Init() {
	Db = &DbManager{}
	Db.connection()
	Db.CreateBucketsIfNotExist()
}

//connection à la BD
func (dbManager *DbManager) connection() {
	var err error
	dbManager.db, err = bolt.Open(DbFile, 0600, nil)
	if err != nil {
		logger.Fatal("Connexion BD - ", err)
	}
}

//disconnection déconnexion de la BD
func (dbManager *DbManager) disconnection() {
	err := dbManager.db.Close()
	if err != nil {
		logger.Fatal("Déconnexion BD - ", err)
	}
}

//CreateBucketsIfNotExist Créé l'ensemble des buckets pour le chat s'ils n'existent pas
func (dbManager *DbManager) CreateBucketsIfNotExist() {
	dbManager.CreateBucket(UserBucket)
	dbManager.CreateBucket(RoomBucket)
}

func (dbManager *DbManager) CreateBucket(bucketName string) {
	dbManager.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			logger.Error("Erreur lors de la création du bucket Room", err)
		}

		return nil
	})
}

//AddValue ajoute dans la BD
func (dbManager *DbManager) AddValue(bucketName string, key string, object fmt.Stringer) error {

	err := dbManager.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}

		encoded := object.String()

		return bucket.Put([]byte(key), []byte(encoded))
	})

	if err == nil {
		logger.Print("Ajout de " + key + " dans la bd")
	} else {
		logger.Error("Erreur lors de l'ajout dans la bd", err)
	}

	return nil
}

//Get Récupère la valeur associée à la clé passée en paramètre dans le bucket bucketName
func (dbManager *DbManager) Get(bucketName string, key string) []byte {
	var value []byte
	// defer dbManager.disconnection()
	dbManager.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		value = bucket.Get([]byte(key))
		if value == nil {
			value = []byte("{}")
		}

		return nil
	})

	return value
}

//GetElementsFromBucket Récupère l'ensemble des éléments du bucket passé en paramètre
func (dbManager *DbManager) GetElementsFromBucket(bucketName string) []string {
	var values []string

	dbManager.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b != nil {
			b.ForEach(func(key, value []byte) error {
				values = append(values, string(value))
				return nil
			})
		}

		return nil
	})

	return values
}
