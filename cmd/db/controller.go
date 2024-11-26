package db

func GetAllAlbums() []Album {
	db := NewDB()
	table, err := db.GetTable("albums")
	if err != nil {
		panic(err)
	}

	albumsSlices, err := table.GetAll()
	if err != nil {
		panic("impossible to get all albums")
	}

	albums := make([]Album, 0, len(albumsSlices))
	for _, slice := range albumsSlices {
		var album Album
		if err := sliceToStruct(slice, &album); err != nil {
			panic("error converting from slice to album")
		}
		albums = append(albums, album)
	}

	return albums
}

func CreateAlbum(album Album) error {
	newAlbum, err := structToMap(album)
	if err != nil {
		return err
	}

	db := NewDB()
	table, err := db.GetTable("albums")
	if err != nil {
		panic(err)
	}

	if err := table.InsertRow(newAlbum); err != nil {
		return err
	}

	return nil
}

func GetAlbumByID(id int32) []Album {
	db := NewDB()
	table, err := db.GetTable("albums")
	if err != nil {
		panic(err)
	}

	albumsSlices, err := table.Find("id", id)
	if err != nil {
		panic("impossible to get all albums")
	}

	albums := make([]Album, 0, len(albumsSlices))
	for _, slice := range albumsSlices {
		var album Album
		if err := sliceToStruct(slice, &album); err != nil {
			panic("error converting from slice to album")
		}
		albums = append(albums, album)
	}

	return albums
}
