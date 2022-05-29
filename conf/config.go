package conf

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"gopkg.in/ini.v1"
)

const (
	configFile = "config.ini"
	winPort    = "3389"
	linPort    = "22"
)

type config struct {
	WinServer    string
	LinServer    string
	Gpath        string
	LinPort      string
	WinPort      string
	LocalLinPort string
	LocalWinPort string
	Zone         string
}

func Get() *config {

	file := SetPath(configFile)

	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalln(file)
		}
	}

	var (
		winServer    string
		linServer    string
		gpath        string
		localLinPort string
		localWinPort string
		zone         string
	)

	cfg, err := ini.Load(file)
	if err != nil {
		log.Fatalln(err)
	}

	linServer = cfg.Section("servers").Key("linserver").String()
	winServer = cfg.Section("servers").Key("winserver").String()
	localLinPort = cfg.Section("local").Key("linport").String()
	localWinPort = cfg.Section("local").Key("winport").String()
	gpath = cfg.Section("gcloud").Key("path").String()
	zone = cfg.Section("gcloud").Key("zone").String()

	return &config{
		WinServer:    winServer,
		LinServer:    linServer,
		Gpath:        gpath,
		LinPort:      linPort,
		WinPort:      winPort,
		LocalLinPort: localLinPort,
		LocalWinPort: localWinPort,
		Zone:         zone,
	}

}

func SetPath(file string) string {

	f, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Fatalln(err)
	}

	p, err := filepath.EvalSymlinks(f)
	if err != nil {
		log.Fatalln(err)
	}

	path := path.Join(path.Dir(p), file)

	return fmt.Sprint(path)
}
