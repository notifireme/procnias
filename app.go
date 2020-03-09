package procnias

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kevinburke/twilio-go"
	"github.com/spf13/viper"
)

// App must be exported
type App struct {
	config *viper.Viper
	router *mux.Router
	logger *log.Logger
	twilio *twilio.Client
}

// NewApp msut be exported
func NewApp() (*App, error) {
	a := new(App)
	a.config = viper.New()
	a.config.SetConfigName("config")
	a.config.AddConfigPath(".")
	if err := a.config.ReadInConfig(); err != nil {
		return nil, err
	}
	a.logger = log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	a.twilio = twilio.NewClient(a.config.GetString("sid"), a.config.GetString("token"), nil)
	a.router = mux.NewRouter()
	a.router.HandleFunc("/health", a.health)
	a.router.HandleFunc("/recive", a.reciver)
	return a, nil
}

// Run dasdasds
func (a App) Run() {
	a.logger.Fatal(http.ListenAndServe(a.config.GetString("port"), a.router))
}
