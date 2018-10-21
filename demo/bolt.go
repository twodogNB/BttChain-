package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

func main() {
	//open db
	db, err := bolt.Open("test1", 0600, nil)
	defer db.Close()
	if err != nil {
		log.Panic(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("a1"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("a1"))
			if err != nil {
				log.Panic(err)
			}
		}

		//write datas
		bucket.Put([]byte("111111"), []byte("hello"))
		bucket.Put([]byte("222222"), []byte("world"))
		return nil

	})
	//operational  databases
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("a1"))
		if err != nil {
			log.Panic(err)
		}

		//read datas
		v1 := bucket.Get([]byte("111111"))
		v2 := bucket.Get([]byte("111111"))
		fmt.Printf("v1 : %s\n", v1)
		fmt.Printf("v2 : %s\n", v2)
		return nil
	})

}
