package mysql

import (
	"database/sql"
	"errors"

	"github.com/Gabriel0110/Golang-Practice/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	statement := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(statement, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Use the LastInsertId() method (not available in all drivers) on the result object to get
	// the ID of our newly inserted record in the snippets table
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, so we convert it to an int
	// type before returning
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	statement := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(statement, id)

	// init a pointer to a new zeroed snippet struct
	s := &models.Snippet{}

	// Use row.Scan() to copy the values from each field in sql.Row to the
	// corresponding field in the Snippet struct. Notice that the arguments
	// to row.Scan are *pointers* to the place you want to copy the data into,
	// and the number of arguments must be exactly the same as the number of
	// columns returned by your statement.
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// If the query returns no rows, then row.Scan() will return a
		// sql.ErrNoRows error. We use the errors.Is() function check for that
		// error specifically, and return our own models.ErrNoRecord error
		// instead.

		// The reason we return models.ErrNoRecord error instead of sql.ErrNoRows directly
		// is to help encapsulate the model completely, so that our application isn’t
		// concerned with the underlying datastore or reliant on datastore-specific errors
		// for its behavior.
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	// If everything went ok, then return the snippet object
	return s, nil

}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
