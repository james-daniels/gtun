package cmd

import (
	"tunnel/conf"
	"tunnel/exec"

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Establish a tunnel session",
	Long:  "Establish the linux or windows tunnels or all tunnels simultaneously",
	Run: func(cmd *cobra.Command, args []string) {

		c := conf.Get()

		switch {
		case linux:
			exec.TunnelUp(c.Command, c.LinServer, c.LinPort, c.LocalLinPort, c.Zone)
		case windows:
			exec.TunnelUp(c.Command, c.WinServer, c.WinPort, c.LocalWinPort, c.Zone)
		default:
			exec.TunnelUp(c.Command, c.LinServer, c.LinPort, c.LocalLinPort, c.Zone)
			exec.TunnelUp(c.Command, c.WinServer, c.WinPort, c.LocalWinPort, c.Zone)
		}
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	upCmd.Flags().BoolVarP(&linux, "linux", "l", false, "bring up linux tunnel")
	upCmd.Flags().BoolVarP(&windows, "windows", "w", false, "bring up windows tunnel")
}
