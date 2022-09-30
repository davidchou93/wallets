package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPgConn(t *testing.T) {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		DBName:   "postgres",
		Password: "postgres",
		Debug:    true,
	}
	t.Run("basic", func(t *testing.T) {
		tfy := assert.New(t)
		got, err := NewPgConn(&dbConfig)
		tfy.Nil(err)
		tfy.NotNil(got)
	})
}
