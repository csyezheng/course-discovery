package config

import (
	"github.com/jinzhu/gorm"

	//"github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/jinzhu/gorm/dialects/sqlite"
)


// Config interface implemented in context (cli) and test packages
type Config interface {
	Debug() bool
	Db() *gorm.DB

	CreateDirectories() error
	MigrateDb()

	ConfigFile() string

	AppName() string
	AppVersion() string
	AppCopyright() string

	SqlServerHost() string
	SqlServerPort() uint
	SqlServerPath() string
	SqlServerPassword() string

	HttpServerHost() string
	HttpServerPort() int
	HttpServerMode() string
	HttpServerPassword() string
	HttpTemplatesPath() string
	HttpFaviconsPath() string
	HttpPublicPath() string
	HttpPublicBuildPath() string
}