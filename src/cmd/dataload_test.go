package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/sky0621/fs-mng-backend/src/models"
)

func TestDataLoad(t *testing.T) {
	db, closeDBFunc := newDB(loadEnv())
	defer func() {
		if err := closeDBFunc(); err != nil {
			log.Fatal(err)
		}
	}()
	var no int
	for {
		if no > 300 {
			break
		}
		no++
		m := &models.ViewingHistory{
			ID:      fmt.Sprintf("%d", no),
			UserID:  "b14b9a71-f74e-42bd-acda-8eeba77e25f4",
			MovieID: "09c87d1d-da43-4ba7-922d-62ca839d2a2c",
		}
		if err := m.Insert(context.Background(), db, boil.Infer()); err != nil {
			log.Fatal(err)
		}
	}
}
