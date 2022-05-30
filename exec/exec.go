package exec

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"

	"gtun/conf"
)

func Execute(env string) {

	c := conf.Get()

	switch {
	case env == "lin":
		startTunnel(c.Command, c.LinServer, c.LinPort, c.LocalLinPort, c.Zone)
	case env == "win":
		startTunnel(c.Command, c.WinServer, c.WinPort, c.LocalWinPort, c.Zone)
	case env == "both":
		startTunnel(c.Command, c.LinServer, c.LinPort, c.LocalLinPort, c.Zone)
		startTunnel(c.Command, c.WinServer, c.WinPort, c.LocalWinPort, c.Zone)
	default:
		fmt.Println(`enter server of "lin", "win", "both" `)
	}
}

func startTunnel(comm, server, port, lport, zone string) {

	conn, _ := net.Dial("tcp", ":"+lport)
	if conn != nil {
		log.Fatalln("Port: " + lport + " is aleady in use.")
	}

	args := []string{
		"compute",
		"start-iap-tunnel",
		server,
		port,
		"--local-host-port=localhost:" + lport,
		"--zone=" + zone,
	}
	cmd := exec.Command(comm, args...)
	err := cmd.Start()
	errHandler(err)

	pid := fmt.Sprint(cmd.Process.Pid)
	pidFile(server, pid)
}

func pidFile(server, pid string) {
	
	file := conf.SetPath(server + ".pid")

	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	errHandler(err)

	defer f.Close()

	_, err = f.Write([]byte(pid))
	errHandler(err)
}

func errHandler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
