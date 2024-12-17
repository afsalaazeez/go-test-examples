package main

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	_ "github.com/go-sql-driver/mysql"
)



func TestrepositoryFind(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "start_date", "due_date", "status", "priority", "description", "created_at"}).
		AddRow(1, "Task 1", time.Now(), time.Now(), "open", "high", "test task", time.Now())

	r := repository{db}

	t.Run("Find Task with Existing ID", func(t *testing.T) {
		mock.ExpectQuery("SELECT (.+) FROM tasks WHERE id=?").WithArgs(1).WillReturnRows(rows)

		task, err := r.Find(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, task.ID)
	})

	t.Run("Find Task with Non-Existing ID", func(t *testing.T) {
		mock.ExpectQuery("SELECT (.+) FROM tasks WHERE id=?").WithArgs(2).WillReturnError(sql.ErrNoRows)

		task, err := r.Find(2)
		assert.Equal(t, sql.ErrNoRows, err)
		assert.Equal(t, Task{}, task)
	})

	t.Run("Find Task with Database Error", func(t *testing.T) {
		mock.ExpectQuery("SELECT (.+) FROM tasks WHERE id=?").WithArgs(3).WillReturnError(errors.New("database error"))

		task, err := r.Find(3)
		assert.NotNil(t, err)
		assert.Equal(t, Task{}, task)
	})

	t.Run("Find Task with Scan Error", func(t *testing.T) {
		mock.ExpectQuery("SELECT (.+) FROM tasks WHERE id=?").WithArgs(4).WillReturnRows(sqlmock.NewRows([]string{"id", "title"}))

		task, err := r.Find(4)
		assert.NotNil(t, err)
		assert.Equal(t, Task{}, task)
	})
}
