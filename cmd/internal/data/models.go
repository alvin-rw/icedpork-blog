package data

import "database/sql"

type Models struct {
	PostModel PostModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		PostModel: PostModel{DB: db},
	}
}
