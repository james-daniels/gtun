package cmd

import (
	"fmt"
	"gtun/conf"
	"gtun/exec"

	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Terminate the tunnel sessions",
	Long: "Terminate the linux or windows tunnels or all tunnels simultaneously",
	Run: func(cmd *cobra.Command, args []string) {

		c := conf.Get()

		switch {
		case linux:
			exec.TunnelDown(c.LinServer)
		case windows:
			exec.TunnelDown(c.WinServer)
		case all:
			exec.TunnelDown(c.LinServer)
			exec.TunnelDown(c.WinServer)
		default:
			fmt.Println("Error: an option must be specified to bring down a tunnel")
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	downCmd.Flags().BoolVarP(&linux, "linux", "l", false, "bring down linux tunnel")
	downCmd.Flags().BoolVarP(&windows, "windows", "w", false, "bring down windows tunnel")
	downCmd.Flags().BoolVarP(&all, "all", "a", false, "bring down all tunnels")
}
