package exec

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"

	"tunnel/conf"
)

func TunnelUp(command, server, port, lport, zone string) {

	conn, _ := net.Dial("tcp", ":"+lport)
	if conn != nil {
		fmt.Println("local port " + lport + " is aleady in use.")
		return
	}

	args := []string{
		"compute",
		"start-iap-tunnel",
		server,
		port,
		"--local-host-port=localhost:" + lport,
		"--zone=" + zone,
	}
	cmd := exec.Command(command, args...)
	err := cmd.Start()
	errHandler(err)

	pid := fmt.Sprint(cmd.Process.Pid)
	pidFile(server, pid)

	validate(lport)
}

func TunnelDown(server, lport string) {

	file := conf.GetPath(server + ".pid")

	pid, _ := os.ReadFile(file)

	exec.Command("kill", "-SIGKILL", string(pid)).Run()

	os.Remove(file)

	validate(lport)
}

func pidFile(server, pid string) {

	file := conf.GetPath(server + ".pid")

	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	errHandler(err)

	defer f.Close()

	_, err = f.Write([]byte(pid))
	errHandler(err)
}

func validate(port string) {

	for i := 0; i < 5; i++ {

		time.Sleep(time.Second)

		conn, _ := net.Dial("tcp", ":"+port)
		if conn != nil {
			fmt.Printf("tunnel up at local port %v\n", port)
			return
		}
	}
	fmt.Printf("tunnel down at local port %v\n", port)
}

func errHandler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
