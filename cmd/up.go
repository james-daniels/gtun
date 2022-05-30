package cmd

import (
	"gtun/exec"
	"gtun/conf"

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "A brief description of your command",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {

		c := conf.Get()

		switch {
		case linux:
			exec.StartTunnel(c.Command, c.LinServer, c.LinPort, c.LocalLinPort, c.Zone, "linux")
		case windows:
			exec.StartTunnel(c.Command, c.WinServer, c.WinPort, c.LocalWinPort, c.Zone,"windows")
		default:
			exec.StartTunnel(c.Command, c.LinServer, c.LinPort, c.LocalLinPort, c.Zone, "linux")
			exec.StartTunnel(c.Command, c.WinServer, c.WinPort, c.LocalWinPort, c.Zone,"windows")
		}
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	upCmd.Flags().BoolVarP(&linux, "linux", "l", false, "bring up linux tunnel")
	upCmd.Flags().BoolVarP(&windows, "windows", "w", false, "bring up windows tunnel")
}
