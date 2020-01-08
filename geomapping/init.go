package geomapping

import (
	"fmt"
	"geomapping/validator"
	"log"
	"os"
	"strings"

	"github.com/ayannahindonesia/basemodel"

	"github.com/fsnotify/fsnotify"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var (
	App *Application
)

type (
	Application struct {
		Name    string      `json:"name"`
		Version string      `json:"version"`
		ENV     string      `json:"env"`
		Config  viper.Viper `json:"prog_config"`
		DB      *gorm.DB    `json:"db"`
	}
)

// Initiate instances
func init() {
	var err error
	App = &Application{}
	App.Name = os.Getenv("APPNAME")
	App.Version = os.Getenv("APPVER")
	App.loadENV()
	if err = App.LoadConfigs(); err != nil {
		log.Printf("Load config error : %v", err)
	}
	if err = App.DBinit(); err != nil {
		log.Printf("DB init error : %v", err)
	}
	// apply custom validator
	v := validator.AsiraValidator{DB: App.DB}
	v.CustomValidatorRules()
}

func (x *Application) Close() {
	var err error
	if err = x.DB.Close(); err != nil {
		log.Println(err)
	}
}

// Loads environtment setting
func (x *Application) loadENV() {
	APPENV := os.Getenv("APPENV")

	switch APPENV {
	default:
		x.ENV = "development"
		break
	case "development":
		x.ENV = "development"
		break
	case "staging":
		x.ENV = "staging"
		break
	case "production":
		x.ENV = "production"
		break
	}
}

// Loads general configs
func (x *Application) LoadConfigs() error {
	var conf *viper.Viper

	conf = viper.New()
	conf.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	conf.AutomaticEnv()
	conf.SetConfigName("config")
	conf.AddConfigPath("$GOPATH/src/" + x.Name)
	conf.SetConfigType("yaml")
	if err := conf.ReadInConfig(); err != nil {
		return err
	}
	conf.WatchConfig()
	conf.OnConfigChange(func(e fsnotify.Event) {
		log.Println("App Config file changed %s:", e.Name)
		x.LoadConfigs()
	})
	x.Config = viper.Viper(*conf)

	return nil
}

// Loads DB postgres configs
func (x *Application) DBinit() error {
	dbconf := x.Config.GetStringMap(fmt.Sprintf("%s.database", x.ENV))
	Cons := basemodel.DBConfig{
		Adapter:        basemodel.PostgresAdapter,
		Host:           dbconf["host"].(string),
		Port:           dbconf["port"].(string),
		Username:       dbconf["username"].(string),
		Password:       dbconf["password"].(string),
		Table:          dbconf["table"].(string),
		Timezone:       dbconf["timezone"].(string),
		Maxlifetime:    dbconf["maxlifetime"].(int),
		IdleConnection: dbconf["idle_conns"].(int),
		OpenConnection: dbconf["open_conns"].(int),
		SSL:            dbconf["sslmode"].(string),
		Logmode:        dbconf["logmode"].(bool),
	}
	basemodel.Start(Cons)
	x.DB = basemodel.DB

	return nil
}
