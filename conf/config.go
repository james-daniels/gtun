package conf

import (
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
	Command      string
	LinPort      string
	WinPort      string
	LocalLinPort string
	LocalWinPort string
	Zone         string
}

func Get() *config {

	file := GetPath(configFile)

	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("file %v does not exist\n", file)
		}
	}
	
	cfg, err := ini.Load(file)
	if err != nil {
		log.Fatalln(err)
	}

	var (
		winServer    string
		linServer    string
		command      string
		localLinPort string
		localWinPort string
		zone         string
	)

	linServer = cfg.Section("servers").Key("linserver").String()
	winServer = cfg.Section("servers").Key("winserver").String()
	localLinPort = cfg.Section("local").Key("linport").String()
	localWinPort = cfg.Section("local").Key("winport").String()
	command = cfg.Section("gcloud").Key("command").String()
	zone = cfg.Section("gcloud").Key("zone").String()

	return &config{
		WinServer:    winServer,
		LinServer:    linServer,
		Command:      command,
		LinPort:      linPort,
		WinPort:      winPort,
		LocalLinPort: localLinPort,
		LocalWinPort: localWinPort,
		Zone:         zone,
	}
}

func GetPath(file string) string {

	f, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Fatalln(err)
	}

	p, err := filepath.EvalSymlinks(f)
	if err != nil {
		log.Fatalln(err)
	}

	return path.Join(path.Dir(p), file)
}
