package page

import "example.com/project/config"

type Repository struct {
	AppConfig *config.AppConfig
	DBConfig  *config.DbConfig
}
var Repo *Repository

func NewRepo(ac *config.AppConfig, db *config.DbConfig) *Repository {
	return &Repository{
		AppConfig: ac,
		DBConfig: db,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}