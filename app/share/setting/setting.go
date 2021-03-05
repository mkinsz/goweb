package setting

import (
	"encoding/json"
	"goweb/app/share/database"
	"goweb/app/share/server"
	"goweb/app/share/session"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Configuration struct {
	Database database.Info   `json:"Database"`
	Server   server.Server   `json:"Server"`
	Session  session.Session `json:"Session"`
}

var Config = &Configuration{}

// Setup initialize the configuration instance
func init() {
	input := io.ReadCloser(os.Stdin)
	defer input.Close()
	path, e := filepath.Abs("config" + string(os.PathSeparator) + "config.json")
	if e != nil {
		log.Fatalln("Fail to find config ", e)
	}

	if input, e = os.Open(path); e != nil {
		log.Fatalln("Open file failed", e)
	}

	jbytes, e := ioutil.ReadAll(input)
	if e != nil {
		log.Fatalln(e)
	}

	if e := json.Unmarshal(jbytes, Config); e != nil {
		log.Fatalln("Could not parse ", jbytes, e)
	}
}
