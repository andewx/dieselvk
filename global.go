package dieselvk

import (
	"log"

	"github.com/andewx/dieselvk/json"
)

//Global Variables
var (
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	WarnLog    *log.Logger
	TraceLog   *log.Logger
	Dictionary json.Dictionary
	Vlk        json.Vlk
)
