package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

var db *sql.DB

func main() {
	// Capture connection properties.
	// cfg := mysql.Config{
	// 	User:   os.Getenv("DBUSER"),
	// 	Passwd: os.Getenv("DBPASS"),
	// 	Net:    "tcp",
	// 	Addr:   "127.0.0.1:3306",
	// 	DBName: "recordings",
	// }
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/recordings")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// get album by name
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found by name: %v\n", albums)

	// Get album by id
	album, err := getAlbumById(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("album found by id: %v\n", album)

	// add album

	newalbum := Album{Title: "The Score", Artist: "Ranadheer", Price: 99.99}
	id, err := addAlbum(newalbum)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New id is:", id)
	// print the new added album
	newRow, err := getAlbumById(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newRow)

}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

func getAlbumById(id int64) (Album, error) {

	row := db.QueryRow("select * from album where id = ?", id)
	var alb Album
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumn by artist %q:%v", id, err)
	}
	return alb, nil

}

func addAlbum(album Album) (int64, error) {

	row, err := db.Exec("insert into album(title,artist,price) values(?,?,?)", album.Title, album.Artist, album.Price)

	if err != nil {
		return 0, fmt.Errorf("NOT INSERTED")
	}

	id, err := row.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error")
	}
	return id, nil

}
