package main

import (
	"testing"
	"database/sql"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
)




func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
func TestNewRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	testCases := []struct {
		name    string
		id      int
		wantErr bool
		err     error
	}{
		{
			name:    "TestFindNonExisting",
			id:      9999,
			wantErr: true,
			err:     sql.ErrNoRows,
		},
		{
			name:    "TestFindWithInvalidID",
			id:      -1,
			wantErr: true,
			err:     sql.ErrNoRows,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"id", "title", "description", "created_at"}).
				AddRow(1, "Test Task", "This is a test task", time.Now())

			mock.ExpectQuery("^SELECT (.+) FROM tasks WHERE id = ?$").
				WithArgs(tc.id).
				WillReturnRows(rows).
				WillReturnError(tc.err)

			repo := NewRepository(db)
			task, err := repo.Find(tc.id)

			if tc.wantErr {
				assert.Error(t, err)
				t.Log("Error is expected when retrieving a non-existing task or a task with an invalid ID")
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, task)
				t.Log("No error is expected when retrieving an existing task")
			}
		})
	}
}
