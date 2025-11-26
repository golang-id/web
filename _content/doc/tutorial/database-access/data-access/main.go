package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/lib/pq"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// albumsByArtist mengambil album-album berdasarkan nama artis.
func albumsByArtist(db *sql.DB, name string) ([]Album, error) {
	// albums adalah slice yang menyimpan data dari hasil kueri.
	var albums []Album

	rows, err := db.Query(`
		SELECT id, title, artist, price
		FROM album
		WHERE artist = $1`, name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	// Iterasi pada rows, menggunakan Scan untuk menyimpan data ke dalam
	// struct Album.
	for rows.Next() {
		var alb Album
		err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		if err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// albumByID kueri album berdasarkan ID.
func albumByID(db *sql.DB, id int64) (Album, error) {
	// Variabel yang menyimpan baris kembalian dari basis-data.
	var alb Album

	row := db.QueryRow(`
		SELECT id, title, artist, price
		FROM album WHERE id = $1", id)
	err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: album tidak ditemukan", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum menambahkan sebuah album baru ke dalam basis-data dan
// mengembalikan ID album yang baru.
func addAlbum(db *sql.DB, alb Album) (int64, error) {
	var id int64
	err := db.QueryRow(`
		INSERT INTO album (title, artist, price)
		VALUES ($1, $2, $3)
		RETURNING id`,
		alb.Title, alb.Artist, alb.Price).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func main() {
	// Contoh DATABASE_URL = "postgres://username:password@localhost:5432/database_name"
	databaseUrl := os.Getenv("DATABASE_URL")
	var connector *pq.Connector
	var err error
	connector, err = pq.NewConnector(databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	var db *sql.DB
	db = sql.OpenDB(connector)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Terhubung!")

	albums, err := albumsByArtist(db, "John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums ditemukan: %v\n", albums)

	// Tulis langsung ID 2 untuk menguji kueri.
	alb, err := albumByID(db, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album ditemukan: %v\n", alb)

	albID, err := addAlbum(db, Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID dari album baru: %v\n", albID)
}
