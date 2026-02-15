package utils

import (
	"database/sql"
	"fmt"
	"go-guide/db_access/types"
	"go-guide/db_access/db"
)

// GetAllRow queries all rows.
func GetAllRow() ([]types.Album, error) {
    // An albums slice to hold data from returned rows.
    var albums []types.Album

    rows, err := db.DB.Query("SELECT * FROM album")
    if err != nil {
        return nil, fmt.Errorf("getAllRow: %v", err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb types.Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("getAllRow: %v", err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("getAllRow: %v", err)
    }
    return albums, nil
}
// AlbumsByArtist queries for albums that have the specified artist name.
func AlbumsByArtist(name string) ([]types.Album, error) {
    // An albums slice to hold data from returned rows.
    var albums []types.Album

    rows, err := db.DB.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb types.Album
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


// AlbumByID queries for the album with the specified ID.
func AlbumByID(id int64) (types.Album, error) {
    // An album to hold data from the returned row.
    var alb types.Album

    row := db.DB.QueryRow("SELECT * FROM album WHERE id = ?", id)
    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
        if err == sql.ErrNoRows {
            return alb, fmt.Errorf("albumsById %d: no such album", id)
        }
        return alb, fmt.Errorf("albumsById %d: %v", id, err)
    }
    return alb, nil
}

// AddAlbum adds the specified album to the database,
// returning the album ID of the new entry
func AddAlbum(alb types.Album) (int64, error) {
    result, err := db.DB.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return id, nil
}