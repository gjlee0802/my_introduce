package setting

import (
	"os"
)

type Database struct {
	Driver   string `toml:"driver"`
	Server   string `toml:"tcp"`
	User     string `toml:"user"`
	Pass     string `toml:"pass"`
	Database string `toml:"database"`
}

func (ds *Database) Getenv() {
	ds.Driver = os.Getenv("DB_DRIVER")
	if ds.Driver == "" {
		ds.Driver = "mysql"
	}
	ds.Server = os.Getenv("DB_SERVER")
	if ds.Server == "" {
		ds.Server = "localhost:3306"
	}
	ds.User = os.Getenv("DB_USER")
	if ds.User == "" {
		ds.User = "root"
	}
	ds.Pass = os.Getenv("DB_PASS")
	if ds.Pass == "" {
		ds.Pass = ""
	}
	ds.Database = os.Getenv("DB_DATABASE")
	if ds.Database == "" {
		ds.Database = "myDB"
	}
}

var DBsetting = &Database{}

func init() {
	DBsetting.Getenv()
}