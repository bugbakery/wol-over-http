package main

import "net"
import "regexp"
import "fmt"

const PKG_HEADER_SIZE = 6
const PKG_ADDR_TIMES = 16
const PKG_SIZE = PKG_HEADER_SIZE + PKG_ADDR_TIMES * 6

var NULL_ADDR = [6]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func ParseMac(mac string) ([6]byte, error) {
    parsedMac, err := net.ParseMAC(mac)
	if err != nil {
		return NULL_ADDR, err
	}

    if len(parsedMac) != 6 {
        return NULL_ADDR, fmt.Errorf("Invalid MAC address length")
    }

    var macBytes [6]byte
    for i := range macBytes {
        macBytes[i] = parsedMac[i]
    }

    return macBytes, nil
}

func makeWolPacket(macAddress [6]byte) [PKG_SIZE]byte {
    var wolPkg [PKG_SIZE]byte

    // 6 times 0xFF
    for idx := range PKG_HEADER_SIZE {
        wolPkg[idx] = 0xFF
	}

    // 16 times MAC address
    for rep := range PKG_ADDR_TIMES {
        for i, b := range macAddress {
            wolPkg[PKG_HEADER_SIZE + rep * 6 + i] = b
        }
	}

    return wolPkg
}

func Wake(macAddress [6]byte) error {
	wolPacket := makeWolPacket(macAddress)

	bcastAddr := fmt.Sprintf("%s:%s", "255.255.255.255", "9")
	udpAddr, err := net.ResolveUDPAddr("udp", bcastAddr)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(wolPacket[:])
	if err != nil {
		return err
	}

	return nil
}
