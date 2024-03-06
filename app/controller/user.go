package controller

import (
	"log"
	"net/http"
	"time"

	"cheki-back/model"
	"github.com/jmoiron/sqlx"
)

type SqlxUser struct {
	DB *sqlx.DB
}

func (s *SqlxUser) GetUser(w http.ResponseWriter, r *http.Request) {
	userMaps := []map[string]interface{}{
		{"name": "Admin", "password": "admin", "role": "admin", "created": time.Now(), "modified": time.Now()},
		{"name": "Kaito", "password": "kaito", "role": "kaito", "created": time.Now(), "modified": time.Now()},
		{"name": "Tester", "password": "tester", "role": "tester", "created": time.Now(), "modified": time.Now()},
	}

	_, err := s.DB.NamedExec(`INSERT INTO user (name, password, role) VALUES (:name, :password, :role)`, userMaps)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}

	user := model.User{Name: "Jane", Password: "password", Role: "student", Created: time.Now(), Modified: time.Now()}
	_, err = s.DB.NamedExec("INSERT INTO user (name, password, role) VALUES (:name, :password, :role)", user)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}

	_, _ = w.Write([]byte(`{"status": "complete"}`))
}
