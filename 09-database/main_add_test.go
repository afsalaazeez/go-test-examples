package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// Type definition repository from package main has definition
type repository struct {
	db *sql.DB
}

// Type definition Task from package main has definition
type Task struct {
	ID          int
	Title       string
	StartDate   time.Time
	DueDate     time.Time
	Status      bool
	Priority    bool
	Description string
	CreatedAt   time.Time
}

func (r repository) Add(task Task) (int64, error) {
	stmt, err := r.db.Prepare("INSERT INTO tasks (title,start_date,due_date,status,priority,description,created_at) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(task.Title, task.StartDate, task.DueDate, task.Status, task.Priority, task.Description, task.CreatedAt)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func TestrepositoryAdd(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Create a new repository instance
	repo := repository{
		db: db,
	}

	// Define test scenarios
	testScenarios := []struct {
		name       string
		task       Task
		mock       func()
		wantErr    bool
		expectedID int64
	}{
		{
			name: "Valid input",
			task: Task{
				Title:       "Test task",
				StartDate:   time.Now(),
				DueDate:     time.Now().Add(24 * time.Hour),
				Status:      true,
				Priority:    true,
				Description: "This is a test task",
				CreatedAt:   time.Now(),
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"ID"}).AddRow(1)
				mock.ExpectQuery("^INSERT INTO tasks").WithArgs("Test task", sqlmock.AnyArg(), sqlmock.AnyArg(), true, true, "This is a test task", sqlmock.AnyArg()).WillReturnRows(rows)
			},
			wantErr:    false,
			expectedID: 1,
		},
		{
			name: "Database error",
			task: Task{
				Title:       "Test task",
				StartDate:   time.Now(),
				DueDate:     time.Now().Add(24 * time.Hour),
				Status:      true,
				Priority:    true,
				Description: "This is a test task",
				CreatedAt:   time.Now(),
			},
			mock: func() {
				mock.ExpectQuery("^INSERT INTO tasks").WithArgs("Test task", sqlmock.AnyArg(), sqlmock.AnyArg(), true, true, "This is a test task", sqlmock.AnyArg()).WillReturnError(errors.New("database error"))
			},
			wantErr: true,
		},
	}

	// Run test scenarios
	for _, tt := range testScenarios {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.mock()

			// Act
			gotID, err := repo.Add(tt.task)

			// Assert
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedID, gotID)
			}

			// Check that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
