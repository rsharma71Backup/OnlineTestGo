package daoimpl

import (
	"database/sql"
	"log"
)

type TypeImpl struct{}

func (dao TypeImpl) GetIdfromType(s string) int64 {
	var id int64
	db := connection()
	defer db.Close()
	err := db.QueryRow("select id from type where value=?", s).Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No type with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("Type is %v\n", id)
	}
	return id
}
