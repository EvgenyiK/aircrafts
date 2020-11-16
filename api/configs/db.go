package configs

import "github.com/go-pg/pg"

func Connect() *pg.DB{
	opts:= &pg.Options{
		User: "db_username",
		Password: "db_password",
		Addr: "localhost:5432",
		Database: "db_name",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		
	}
}