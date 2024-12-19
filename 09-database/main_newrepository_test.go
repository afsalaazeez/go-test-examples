package main

import (
	"database/sql"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {

	scenarios := []struct {
		name        string
		dbActive    bool
		dbConnected bool
		expectedErr error
	}{
		{
			name:        "Successful Repository Creation",
			dbActive:    true,
			dbConnected: true,
			expectedErr: nil,
		},
		{
			name:        "Failed Repository Creation",
			dbActive:    false,
			dbConnected: true,
			expectedErr: errors.New("database not active"),
		},
		{
			name:        "Repository Creation with Inactive Database",
			dbActive:    true,
			dbConnected: false,
			expectedErr: errors.New("database not connected"),
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			if !s.dbActive {
				mock.ExpectPing().WillReturnError(sql.ErrConnDone)
			}

			if !s.dbConnected {
				db.Close()
			}

			repo := NewRepository(db)

			if s.expectedErr == nil {
				assert.NotNil(t, repo)
			} else {
				assert.Nil(t, repo)
			}

			if s.expectedErr != nil {
				assert.EqualError(t, err, s.expectedErr.Error())
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
