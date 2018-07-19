package model

// model层操作
type Book struct {
	Tile   string
	Author string
	Price  float32
}

func AllBook() ([]*Book, error) {
	rows, err := db.Query("select * from books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.Author, &bk.Price, &bk.Tile)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}
