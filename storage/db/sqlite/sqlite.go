package sqlite

import (
	"database/sql"
	"fmt"
	"my_lib/models/author"
	"my_lib/models/book"

	_ "modernc.org/sqlite"
)

type SqliteStorage struct {
	db *sql.DB
}

func NewSqliteStorage(db *sql.DB) *SqliteStorage {
	return &SqliteStorage{
		db: db,
	}
}

func (s *SqliteStorage) GetBookCount() (int, error) {
	q := `SELECT COUNT(*) FROM book`
	row, err := s.db.Query(q)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var count int

	row.Next()
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *SqliteStorage) GetPublishingHouseList() ([]book.PublishingHouse, error) {
	q := `SELECT * FROM publishing_house ORDER BY name`

	rows, err := s.db.Query(q)
	if err != nil {
		return []book.PublishingHouse{}, err
	}
	defer rows.Close()

	list := []book.PublishingHouse{}
	for rows.Next() {
		ph := book.PublishingHouse{}
		err := rows.Scan(&ph.Id, &ph.Name)

		if err != nil {
			return []book.PublishingHouse{}, err
		}

		list = append(list, ph)
	}

	return list, nil
}

func (s *SqliteStorage) LinkAuthorAndLiteraryWork(authorId, bookId int) error {
	q := `INSERT INTO author_and_literary_work (author_id, literary_work_id) VALUES(:a_id, :b_id)`
	_, err := s.db.Exec(
		q,
		sql.Named("a_id", authorId),
		sql.Named("b_id", bookId),
	)

	if err != nil {
		return fmt.Errorf("can't link author and literary work: %w", err)
	}

	return nil
}

func (s *SqliteStorage) LinkBookAndLiteraryWork(lwId, bookId int) error {
	q := `INSERT INTO book_and_literary_work (literary_work_id, book_id) VALUES(:lw_id, :b_id)`
	_, err := s.db.Exec(
		q,
		sql.Named("lw_id", lwId),
		sql.Named("b_id", bookId),
	)

	if err != nil {
		return fmt.Errorf("can't link literary work and book: %w", err)
	}

	return nil
}

func (s *SqliteStorage) GetBookList(limit, offset int) ([]book.BookUnload, error) {
	q := `
	SELECT
 		b.id AS id,
 		GROUP_CONCAT(IFNULL(a.last_name || ' ' || a.name, '-'), ', ') AS author,
 		lw.name AS name,
 		ph.name AS publishing_house,
 		b.year_of_publication
	FROM book AS b
	LEFT JOIN publishing_house AS ph ON ph.id = b.publishing_house_id
	LEFT JOIN book_and_literary_work AS blw ON blw.book_id = b.id
	LEFT JOIN literary_work AS lw ON lw.id = blw.literary_work_id
	LEFT JOIN author_and_literary_work AS alw ON alw.literary_work_id = lw.id
	LEFT JOIN authors AS a ON a.id = alw.author_id
	GROUP BY b.id
	ORDER BY :sortedField
	LIMIT :limit OFFSET :offset;`

	// a.last_name, lw.name

	sortedField := "b.year_of_publication"
	rows, err := s.db.Query(q,
		sql.Named("limit", limit),
		sql.Named("offset", offset),
		sql.Named("sortedField", sortedField),
	)
	if err != nil {
		return []book.BookUnload{}, err
	}
	defer rows.Close()

	list := []book.BookUnload{}
	for rows.Next() {
		bRow := book.BookUnload{}
		err := rows.Scan(&bRow.Id, &bRow.Author, &bRow.Name, &bRow.PublishingHouse, &bRow.PublishingYear)

		if err != nil {
			return []book.BookUnload{}, err
		}

		list = append(list, bRow)
	}

	return list, nil
}

func (s *SqliteStorage) GetAuthorByName(a author.Author) ([]int, error) {
	q := `SELECT id FROM authors 
			WHERE name=:name
			AND father_name=:father_name
			AND last_name=:last_name`

	rows, err := s.db.Query(q,
		sql.Named("name", a.Name),
		sql.Named("father_name", a.FatherName),
		sql.Named("last_name", a.LastName),
	)
	list := []int{}
	if err != nil {
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return list, err
		}

		list = append(list, int(id))
	}

	return list, nil
}

func (s *SqliteStorage) GetAuthorList() ([]author.Author, error) {
	q := `SELECT * FROM authors ORDER BY last_name`

	rows, err := s.db.Query(q)
	if err != nil {
		return []author.Author{}, err
	}
	defer rows.Close()

	list := []author.Author{}
	for rows.Next() {
		authorRow := author.Author{}
		err := rows.Scan(&authorRow.Id, &authorRow.Name, &authorRow.FatherName, &authorRow.LastName)

		if err != nil {
			return []author.Author{}, err
		}

		list = append(list, authorRow)
	}

	return list, nil
}

func (s *SqliteStorage) AddLiteraryWork(lwName string) (int, error) {
	q := `INSERT INTO literary_work (name) VALUES(:name)`
	res, err := s.db.Exec(
		q,
		sql.Named("name", lwName),
	)

	if err != nil {
		return 0, fmt.Errorf("can't add new literary work: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *SqliteStorage) AddPhysicalBook(b *book.BookAdd) (int, error) {
	q := `INSERT INTO book (publishing_house_id, year_of_publication) VALUES(:ph_id, :yop)`
	res, err := s.db.Exec(
		q,
		sql.Named("ph_id", b.PublishingHouse.Id),
		sql.Named("yop", b.PublishingYear),
	)

	if err != nil {
		return 0, err
	}

	bookId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(bookId), nil
}

func (s *SqliteStorage) AddPublishingHouse(phName string) (int, error) {
	q := `INSERT INTO publishing_house (name) VALUES(:name)`
	res, err := s.db.Exec(
		q,
		sql.Named("name", phName),
	)

	if err != nil {
		return 0, fmt.Errorf("can't add new publishing house: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *SqliteStorage) AddAuthor(a author.Author) (int, error) {
	q := `INSERT INTO authors (name, father_name, last_name) VALUES(:name, :father_name, :last_name)`

	res, err := s.db.Exec(
		q,
		sql.Named("name", a.Name),
		sql.Named("father_name", a.FatherName),
		sql.Named("last_name", a.LastName),
	)

	if err != nil {
		return 0, fmt.Errorf("can't add new author: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
