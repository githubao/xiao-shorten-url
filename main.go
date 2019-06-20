// run code
// author: baoqiang
// time: 2019-06-20 11:22
package main

import (
	"flag"
	"github.com/boltdb/bolt"
	goji "goji.io"
	"goji.io/pat"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8000", "Port for server")
	dbfile := flag.String("dbfile", "my.db", "full path of db file")
	flag.Parse()

	db, err := bolt.Open(*dbfile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_ = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("shortURL"))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte("viewCount"))
		if err != nil {
			return err
		}
		return nil
	})

	// new instance
	h := DBHandler{DB: db}

	// mut
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/api"), h.api)
	mux.HandleFunc(pat.Post("/api"), h.api)
	mux.HandleFunc(pat.Get("/:surl"), h.getShortUrl)

	// run server
	log.Printf("run server at: http://localhost:%s\n", *port)
	_ = http.ListenAndServe("localhost:"+*port, mux)
}
