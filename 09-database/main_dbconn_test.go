package main

import (
	"database/sql"
	"testing"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	_ "github.com/go-sql-driver/mysql"
)

func TestdbConn(t *testing.T) {
	t.Parallel()

	t.Run("Successful Database Connection", func(t *testing.T) {

		dbMock, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer dbMock.Close()

		mock.ExpectPing().WillReturnError(nil)

		db := dbConn()

		assert.NotNil(t, db)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("Database Connection with Invalid Credentials", func(t *testing.T) {

		dbMock, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer dbMock.Close()

		mock.ExpectPing().WillReturnError(sql.ErrConnDone)

		db := dbConn()

		assert.Nil(t, db)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("Database Connection Timeout", func(t *testing.T) {

		dbMock, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer dbMock.Close()

		mock.ExpectPing().WillReturnError(sql.ErrConnDone)

		db := dbConn()

		assert.Nil(t, db)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

