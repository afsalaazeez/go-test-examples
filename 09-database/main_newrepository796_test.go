package main

import (
	"database/sql"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {

	testCases := []struct {
		name     string
		mockFunc func() *sql.DB
		wantErr  bool
	}{
		{
			name: "Successful Repository Creation with an Active and Connected Database",
			mockFunc: func() *sql.DB {
				db, _, _ := sqlmock.New()
				return db
			},
			wantErr: false,
		},
		{
			name: "Failed Repository Creation with an Inactive Database",
			mockFunc: func() *sql.DB {
				return nil
			},
			wantErr: true,
		},
		{
			name: "Failed Repository Creation with a Disconnected Database",
			mockFunc: func() *sql.DB {
				db, mock, _ := sqlmock.New()
				mock.ExpectClose()
				db.Close()
				return db
			},
			wantErr: true,
		},
		{
			name: "Handling Database Errors During Repository Creation",
			mockFunc: func() *sql.DB {
				db, mock, _ := sqlmock.New()
				mock.ExpectQuery("SELECT 1").WillReturnError(sql.ErrNoRows)
				return db
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := tc.mockFunc()
			repo := NewRepository(db)
			if tc.wantErr {
				assert.Nil(t, repo, "expected repository to be nil")
			} else {
				assert.NotNil(t, repo, "expected repository to be not nil")
			}
		})
	}
}
