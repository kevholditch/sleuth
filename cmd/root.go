package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/kevholditch/sleuth/internal/version"
	"github.com/liamg/tml"
	"github.com/spf13/cobra"
	"net"
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

type Packet struct {
	IPV4 layers.IPv4
	TCP layers.TCP
}

func buildKey (ip1, ip2 net.IP) string{
	str1 := ip1.String()
	str2 := ip2.String()
	if str1 < str2 {
		return fmt.Sprintf("%s<->%s", str1, str2)
	}
	return fmt.Sprintf("%s<->%s", str2, str1)
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
	connections := map[string][]Packet{}
	i := 0
	for packet := range packetSource.Packets() {
		for _, layer := range packet.Layers() {
			switch layer.LayerType() {
			case layers.LayerTypeIPv4:
				ip = layer.(*layers.IPv4)
				break
			case layers.LayerTypeTCP: {
					tcp := layer.(*layers.TCP)
					p := Packet{
						IPV4: *ip,
						TCP:  *tcp,
					}
				key := buildKey(ip.SrcIP, ip.DstIP)
				if val, ok := connections[key]; ok {
						val = append(val, p)
						connections[key] = val
					}else{
						connections[key] = []Packet{p}
					}
				}
			}

		}
		i++
		if i>1000{
			break
		}
	}

	for key, packets := range connections {
		fmt.Printf("%s : %d packets\n", key, len(packets))

	}

	return nil
}
