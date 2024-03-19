package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/ccding/go-stun/stun"
)

const (
	broadcastIP = "255.255.255.255"
	broadcastPort = 5000
)

type File struct{
	Name 		string
	Size		int64
	Chunks		[][]byte
	Uploader	string
}

var files map[string]File

func main() {
	files = make(map[string]File)

	select()
}

func listenForBroadcast(){
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", broadcastIP, broadcastPort))
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn , err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error listening for UDP:", err)
		return
	}

	defer conn.Close()

	fmt.Println("Listening for broadcast messages on "+ udpAddr.String())

	buf := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP((buf))
		if err != nil {
			fmt.Println("Error reading UDP message:", err)
			continue
		}
		fmt.Println("Received Broadcast message:", string(buf[:n]))
	}
}

func broadcastPresence(){
	net.UDPAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", broadcastIP, broadcastPort))
	if err != nill {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, net.UDPAddr)
	if err != nil {
		fmt.Println("Error dialing UDP address:", err)
		return
	}
	defer conn.Close()

	fmt.Println("broadcasting presence to" + net.UDPAddr.IP.String())
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for range ticker.C{
		message := "hello from my node!!!"
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error broadcasting presence:", err)
		}
		fmt.Println("broadcasted presence:", message)
	}
}

func discoverPublicIP(){
	c := stun.NewClient()
	fmt.Println("discovering public ip using stun .....")

	for{
		result, err := c.Discover()
		if err != nil {
			fmt.Println("STUN discovery error:", err)
			return
		}

		fmt.Println("PUBLIC IP:", result.MappedAddress)
		fmt.Println("NAT TYPE:", result.NATType)

		time.Sleep(time.Second * 30)
	}
}

func uploadFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil{
		return err
	}

	fileSize := fileInfo.Size()
	fileName = fileInfo.Name()

	chunkSize := 1024 * 1024
	chunks := mke([][]byte, 0)
	buffer := make([]byte, chunkSize)
	
	for{
		n,err := file.Read(buffer)
		if err == io.EOF{
			break
		}
		chunks = append(chunks, buffer[:n])
	}

	f := File(
		Name: fileName,
		Size: fileSize,
		Cunks: chunks,
		Uploader: "asdasd"
	)
} 