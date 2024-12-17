package main

import (
	"database/sql"
	"errors"
	"log"
	"testing"
	"time"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)



func TestrepositoryAdd(t *testing.T) {

	tests := []struct {
		name    string
		task    Task
		mock    func()
		wantErr bool
	}{
		{
			name: "Successful Insertion of Task",
			task: Task{
				Title:       "Test Task",
				StartDate:   time.Now(),
				DueDate:     time.Now().AddDate(0, 0, 10),
				Status:      true,
				Priority:    true,
				Description: "Test Task Description",
				CreatedAt:   time.Now(),
			},
			mock: func() {

			},
			wantErr: false,
		},
		{
			name: "Insertion of Task with Invalid Fields",
			task: Task{
				Title:       "",
				StartDate:   time.Now(),
				DueDate:     time.Now().AddDate(0, 0, -10),
				Status:      true,
				Priority:    true,
				Description: "Test Task Description",
				CreatedAt:   time.Now(),
			},
			mock: func() {

			},
			wantErr: true,
		},
		{
			name: "Database Connection Error",
			task: Task{
				Title:       "Test Task",
				StartDate:   time.Now(),
				DueDate:     time.Now().AddDate(0, 0, 10),
				Status:      true,
				Priority:    true,
				Description: "Test Task Description",
				CreatedAt:   time.Now(),
			},
			mock: func() {

			},
			wantErr: true,
		},
		{
			name: "Error while Fetching LastInsertId",
			task: Task{
				Title:       "Test Task",
				StartDate:   time.Now(),
				DueDate:     time.Now().AddDate(0, 0, 10),
				Status:      true,
				Priority:    true,
				Description: "Test Task Description",
				CreatedAt:   time.Now(),
			},
			mock: func() {

			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			tt.mock()

			repo := repository{db: db}

			_, err = repo.Add(tt.task)

			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
