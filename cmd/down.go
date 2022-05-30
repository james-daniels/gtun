package cmd

import (
	"fmt"
	"gtun/conf"
	"gtun/exec"

	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "A brief description of your command",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {

		c := conf.Get()

		switch {
		case linux:
			exec.StopTunnel(c.LinServer)
		case windows:
			exec.StopTunnel(c.WinServer)
		case all:
			exec.StopTunnel(c.LinServer)
			exec.StopTunnel(c.WinServer)
		default:
			fmt.Println("an option must be specified to bring down a tunnel")
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	downCmd.Flags().BoolVarP(&linux, "linux", "l", false, "bring down linux tunnel")
	downCmd.Flags().BoolVarP(&windows, "windows", "w", false, "bring down windows tunnel")
	downCmd.Flags().BoolVarP(&all, "all", "a", false, "bring down both linux and windows tunnels")
}
