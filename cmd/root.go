package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
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
 ______   __       __  __   ______   _________  ___   ___     
/_____/\ /_/\     /_/\/_/\ /_____/\ /________/\/__/\ /__/\    
\::::_\/_\:\ \    \:\ \:\ \\::::_\/_\__.::.__\/\::\ \\  \ \   
 \:\/___/\\:\ \    \:\ \:\ \\:\/___/\  \::\ \   \::\/_\ .\ \  
  \_::._\:\\:\ \____\:\ \:\ \\::___\/_  \::\ \   \:: ___::\ \ 
    /____\:\\:\/___/\\:\_\:\ \\:\____/\  \::\ \   \: \ \\::\ \
    \_____\/ \_____\/ \_____\/ \_____\/   \__\/    \__\/ \::\/

 https://github.com/kevholditch/slueth
 %s

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
	var ip *layers.IPv4
	for packet := range packetSource.Packets() {


		for i, layer := range packet.Layers() {

			switch layer.LayerType() {
			case layers.LayerTypeIPv4:
				ip = layer.(*layers.IPv4)
				break
			case layers.LayerTypeTCP: {
					tcp := layer.(*layers.TCP)
					fmt.Printf("%s:%d -> %s:%d %d bytes %d layers: %d\n", ip.SrcIP, tcp.SrcPort, ip.DstIP, tcp.DstPort, len(tcp.Contents), tcp.Seq, len(packet.Layers()))
					//fmt.Printf("%s\n", packet.String())
				}
			}
			if i == 3 {
				fmt.Printf("contents: %s\n", string(layer.LayerContents()))
			}


		}
	}


	return nil
}
