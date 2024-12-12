package main

import (
	"testing"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/DATA-DOG/go-sqlmock"
	"time"
)

type Task struct {
	ID          int
	Title       string
	StartDate   time.Time
	DueDate     time.Time
	Status      string
	Priority    string
	Description string
	CreatedAt   time.Time
}

type repository struct {
	db *sql.DB
}

func (r repository) Find(id int) (Task, error) {
	var task Task

	rows, _ := r.db.Query("SELECT * FROM tasks WHERE id=?", id)
	defer rows.Close()

	if rows.Next() == false {
		return task, sql.ErrNoRows
	}

	err := rows.Scan(&task.ID, &task.Title, &task.StartDate, &task.DueDate, &task.Status, &task.Priority, &task.Description, &task.CreatedAt)
	if err != nil {
		return task, err
	}

	return task, nil
}

func TestRepositoryFind(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rep := repository{db: db}

	rows := sqlmock.NewRows([]string{"id", "title", "start_date", "due_date", "status", "priority", "description", "created_at"}).
		AddRow(1, "test_title", time.Now(), time.Now(), "status", "priority", "description", time.Now())

	mock.ExpectQuery("^SELECT (.+) FROM tasks WHERE id=?").
		WithArgs(1).
		WillReturnRows(rows)

	task, err := rep.Find(1)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if task.ID != 1 {
		t.Errorf("expected ID %v, but got %v", 1, task.ID)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
