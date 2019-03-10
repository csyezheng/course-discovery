package context

import (
	// "log"
	// "os"
	// "time"

	"github.com/jinzhu/gorm"
	//"github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kylelemons/go-gypsy/yaml"
	"github.com/csyezheng/course-discovery/course_discovery/apps/models"
	// "github.com/urfave/cli"
)


var DB *gorm.DB


const (
	DbTiDB  = "internal"
	DbMySQL = "mysql"
)

// Config provides a struct in which application configuration is stored.
// Application code must use functions to get config values, for two reasons:
//
// 1. Some values are computed and we don't want to leak implementation details (aims at reducing refactoring overhead).
//
// 2. Paths might actually be dynamic later (if we build a multi-user version).
//
// See https://github.com/photoprism/photoprism/issues/50#issuecomment-433856358
type Config struct {
	appName            string
	appVersion         string
	appCopyright       string
	debug              bool
	configFile         string
	assetsPath         string
	cachePath          string
	originalsPath      string
	importPath         string
	exportPath         string
	sqlServerHost      string
	sqlServerPort      uint
	sqlServerPath      string
	sqlServerPassword  string
	httpServerHost     string
	httpServerPort     int
	httpServerMode     string
	httpServerPassword string
	darktableCli       string
	databaseDriver     string
	databaseDsn        string
	db                 *gorm.DB
}

// NewConfig() creates a new configuration entity by using two methods:
//
// 1. SetValuesFromFile: This will initialize values from a yaml config file.
//
// 2. SetValuesFromCliContext: Which comes after SetValuesFromFile and overrides
//    any previous values giving an option two override file configs through the CLI.
func NewConfig() *Config {
	c := &Config{}
	c.appName = ""
	c.appCopyright = ""
	c.appVersion = ""
	c.SetValuesFromFile("../../configs/settings.yml")
	//c.SetValuesFromCliContext(ctx)

	return c
}

// SetValuesFromFile uses a yaml config file to initiate the configuration entity.
func (c *Config) SetValuesFromFile(fileName string) error {
	yamlConfig, err := yaml.ReadFile(fileName)

	if err != nil {
		return err
	}

	c.configFile = fileName
	if debug, err := yamlConfig.GetBool("debug"); err == nil {
		c.debug = debug
	}

	if sqlServerHost, err := yamlConfig.Get("sql-host"); err == nil {
		c.sqlServerHost = sqlServerHost
	}

	if sqlServerPort, err := yamlConfig.GetInt("sql-port"); err == nil {
		c.sqlServerPort = uint(sqlServerPort)
	}

	if sqlServerPassword, err := yamlConfig.Get("sql-password"); err == nil {
		c.sqlServerPassword = sqlServerPassword
	}

	if sqlServerPath, err := yamlConfig.Get("sql-path"); err == nil {
		c.sqlServerPath = sqlServerPath
	}

	if httpServerHost, err := yamlConfig.Get("http-host"); err == nil {
		c.httpServerHost = httpServerHost
	}

	if httpServerPort, err := yamlConfig.GetInt("http-port"); err == nil {
		c.httpServerPort = int(httpServerPort)
	}

	if httpServerMode, err := yamlConfig.Get("http-mode"); err == nil {
		c.httpServerMode = httpServerMode
	}

	if httpServerPassword, err := yamlConfig.Get("http-password"); err == nil {
		c.httpServerPassword = httpServerPassword
	}

	if databaseDriver, err := yamlConfig.Get("database-driver"); err == nil {
		c.databaseDriver = databaseDriver
	}

	if databaseDsn, err := yamlConfig.Get("database-dsn"); err == nil {
		c.databaseDsn = databaseDsn
	}

	return nil
}


// connectToDatabase establishes a database connection.
// When used with the internal driver, it may create a new database server instance.
// It tries to do this 12 times with a 5 second sleep interval in between.
func (c *Config) connectToDatabase() error {

	dbDsn := "zhengye:csye3631@tcp(127.0.0.1:3307)/discovery?parseTime=true"
	db, err := gorm.Open("mysql", dbDsn)

	//if err != nil || db == nil {
	//
	//	for i := 1; i <= 12; i++ {
	//		time.Sleep(5 * time.Second)
	//
	//		db, err = gorm.Open("mysql", dbDsn)
	//
	//		if db != nil && err == nil {
	//			break
	//		}
	//	}
	//
	//	if err != nil || db == nil {
	//		log.Fatal(err)
	//	}
	//}

	c.db = db

	return err
}

// AppName returns the application name.
func (c *Config) AppName() string {
	return c.appName
}

// AppVersion returns the application version.
func (c *Config) AppVersion() string {
	return c.appVersion
}

// AppCopyright returns the application copyright.
func (c *Config) AppCopyright() string {
	return c.appCopyright
}

// Debug returns true if debug mode is on.
func (c *Config) Debug() bool {
	return c.debug
}

// ConfigFile returns the config file name.
func (c *Config) ConfigFile() string {
	return c.configFile
}

// SqlServerHost returns the built-in SQL server host name or IP address (empty for all interfaces).
func (c *Config) SqlServerHost() string {
	return c.sqlServerHost
}

// SqlServerPort returns the built-in SQL server port.
func (c *Config) SqlServerPort() uint {
	return c.sqlServerPort
}

// SqlServerPassword returns the password for the built-in database server.
func (c *Config) SqlServerPassword() string {
	return c.sqlServerPassword
}

// HttpServerHost returns the built-in HTTP server host name or IP address (empty for all interfaces).
func (c *Config) HttpServerHost() string {
	return c.httpServerHost
}

// HttpServerPort returns the built-in HTTP server port.
func (c *Config) HttpServerPort() int {
	return c.httpServerPort
}

// HttpServerMode returns the server mode.
func (c *Config) HttpServerMode() string {
	return c.httpServerMode
}

// HttpServerPassword returns the password for the user interface (optional).
func (c *Config) HttpServerPassword() string {
	return c.httpServerPassword
}

// CachePath returns the path to the cache.
func (c *Config) CachePath() string {
	return c.cachePath
}

// ServerPath returns the server assets path (public files, favicons, templates,...).
//func (c *Config) ServerPath() string {
//	return c.AssetsPath() + "/server"
//}

// HttpTemplatesPath returns the server templates path.
//func (c *Config) HttpTemplatesPath() string {
//	return c.ServerPath() + "/templates"
//}

// HttpFaviconsPath returns the favicons path.
//func (c *Config) HttpFaviconsPath() string {
//	return c.HttpPublicPath() + "/favicons"
//}

// HttpPublicPath returns the public server path (//server/assets/*).
//func (c *Config) HttpPublicPath() string {
//	return c.ServerPath() + "/public"
//}

// HttpPublicBuildPath returns the public build path (//server/assets/build/*).
//func (c *Config) HttpPublicBuildPath() string {
//	return c.HttpPublicPath() + "/build"
//}

// Db returns the db connection.
func (c *Config) Db() *gorm.DB {
	if c.db == nil {
		c.connectToDatabase()
	}

	return c.db
}

// MigrateDb will start a migration process.
func (c *Config) MigrateDb() {
	db := c.Db()
	DB = c.Db()

	db.AutoMigrate(&models.Course{})
}

func Setup() error {
	conf := NewConfig()

	conf.MigrateDb()

	return nil
}