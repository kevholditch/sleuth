package main

import (
	"fmt"
	"github.com/kevholditch/sleuth/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "slueth",
	Short: "Slueth will find out whats happening on your network interface",
	Long:  `Slueth is a commandline network interface listening tool - see https://github.com/kevholditch/slueth for more information`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`

''''''''''',cx0XNWWWWNKOd:,''''''''
''''''''',cONMMWXK00KNWMWXk:,''''''
'''''''',oXMMNOo:;,,;cd0WMWKxl,''''
''''''''lXMMXo,'''''''';xNMMM0:''''
''''''',xWMWx,'''TCP'''':OMMMNo''''      %s
''''''',xWMWx,'''''''''':OMMMNo''''      https://github.com/kevholditch/slueth
''''''''lXMMXo,'''''''';xNMMM0:''''
'''''''',dNMMNOo:,,,;:o0WMMKkl,''''
''''''',l0WMMMMWXK00KNWMMNk:,''''''
''''',lONMMNOx0XNWWWWNKOdc,''''''''
''',lONMMNOc,',:cllllc;,'''''''''''
::oONMMW0dc::::::::::::::::::::::::


`, version.Version)
	},
}
