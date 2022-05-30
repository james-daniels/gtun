package exec

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"

	"tunnel/conf"
)


func TunnelUp(comm, server, port, lport, zone, env string) {

	conn, _ := net.Dial("tcp", ":"+lport)
	if conn != nil {
		log.Fatalln(env + " tunnel at local port " + lport + " is aleady in use.")
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

func TunnelDown(server string) {
	
	file := conf.GetPath(server + ".pid")

	pid, _ := os.ReadFile(file)

	cmd := exec.Command("kill", "-SIGKILL", string(pid))
	cmd.Run()

	os.Remove(file)
}


func pidFile(server, pid string) {
	
	file := conf.GetPath(server + ".pid")

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
