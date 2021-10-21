package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/kevholditch/sleuth/internal/version"
	"github.com/liamg/tml"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "slueth",
	Short: "Slueth will find out whats happening on your network interface",
	Long:  `Slueth is a commandline network interface listening tool - see https://github.com/kevholditch/slueth for more information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`

''''''''''',cx0XNWWWWNKOd:,''''''''
''''''''',cONMMWXK00KNWMWXk:,''''''
'''''''',oXMMNOo:;,,;cd0WMWKxl,''''
''''''''lXMMXo,'''''''';xNMMM0:''''
-------,xWMWx,'''TCP'''':OMMMNo''''      %s
''''''',xWMWx,'''''''''':OMMMNo''''      https://github.com/kevholditch/slueth
''''''''lXMMXo,'''''''';xNMMM0:''''
'''''''',dNMMNOo:,,,;:o0WMMKkl,''''
''''''',l0WMMMMWXK00KNWMMNk:,''''''
''''',lONMMNOx0XNWWWWNKOdc,''''''''
''',lONMMNOc,',:cllllc;,'''''''''''
::oONMMW0dc::::::::::::::::::::::::


`, version.Version)

		err := foo()
		if err != nil {
			_ = tml.Printf("<bold><red>Error:</red></bold>%s\n", err)
			os.Exit(1)
		}


	},
}

func foo() error  {
	timeout     := 30 * time.Second
	snapshotLen := int32(1024)
	promiscuous := true
	handle, err := pcap.OpenLive("en0", snapshotLen, promiscuous, timeout)
	if err != nil {
		return err
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {

		fmt.Printf("%s\n", packet.String())
	}


	return nil
}
