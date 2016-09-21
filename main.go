package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func crc8dallas(Command1 []byte) []byte {
	// CRCTable := [256]byte{0, 94, 188, 226, 97, 63, 221, 131, 194, 156, 126, 32, 163, 253, 31, 65, 157, 195, 33, 127, 252, 162, 64, 30, 95, 1, 227, 189, 62, 96, 130, 220, 35, 125, 159, 193, 66, 28, 254, 160, 225, 191, 93, 3, 128, 222, 60, 98, 190, 224, 2, 92, 223, 129, 99, 61, 124, 34, 192, 158, 29, 67, 161, 255, 70, 24, 250, 164, 39, 121, 155, 197, 132, 218, 56, 102, 229, 187, 89, 7, 219, 133, 103, 57, 186, 228, 6, 88, 25, 71, 165, 251, 120, 38, 196, 154, 101, 59, 217, 135, 4, 90, 184, 230, 167, 249, 27, 69, 198, 152, 122, 36, 248, 166, 68, 26, 153, 199, 37, 123, 58, 100, 134, 216, 91, 5, 231, 185, 140, 210, 48, 110, 237, 179, 81, 15, 78, 16, 242, 172, 47, 113, 147, 205, 17, 79, 173, 243, 112, 46, 204, 146, 211, 141, 111, 49, 178, 236, 14, 80, 175, 241, 19, 77, 206, 144, 114, 44, 109, 51, 209, 143, 12, 82, 176, 238, 50, 108, 142, 208, 83, 13, 239, 177, 240, 174, 76, 18, 145, 207, 45, 115, 202, 148, 118, 40, 171, 245, 23, 73, 8, 86, 180, 234, 105, 55, 213, 139, 87, 9, 235, 181, 54, 104, 138, 212, 149, 203, 41, 119, 244, 170, 72, 22, 233, 183, 85, 11, 136, 214, 52, 106, 43, 117, 151, 201, 74, 20, 246, 168, 116, 42, 200, 150, 21, 75, 169, 247, 182, 232, 10, 84, 215, 137, 107, 53}
	CRCTable := [256]byte{0x00, 0x5E, 0x0BC, 0x0E2, 0x61, 0x3F, 0x0DD, 0x83, 0x0C2, 0x9C, 0x7E, 0x20, 0x0A3, 0x0FD, 0x1F, 0x41, 0x9D, 0x0C3, 0x21, 0x7F, 0x0FC, 0x0A2, 0x40, 0x1E, 0x5F, 0x01, 0x0E3, 0x0BD, 0x3E, 0x60, 0x82, 0x0DC, 0x23, 0x7D, 0x9F, 0x0C1, 0x42, 0x1C, 0x0FE, 0x0A0, 0x0E1, 0x0BF, 0x5D, 0x03, 0x80, 0x0DE, 0x3C, 0x62, 0x0BE, 0x0E0, 0x02, 0x5C, 0x0DF, 0x81, 0x63, 0x3D, 0x7C, 0x22, 0x0C0, 0x9E, 0x1D, 0x43, 0x0A1, 0x0FF, 0x46, 0x18, 0x0FA, 0x0A4, 0x27, 0x79, 0x9B, 0x0C5, 0x84, 0x0DA, 0x38, 0x66, 0x0E5, 0x0BB, 0x59, 0x07, 0x0DB, 0x85, 0x67, 0x39, 0x0BA, 0x0E4, 0x06, 0x58, 0x19, 0x47, 0x0A5, 0x0FB, 0x78, 0x26, 0x0C4, 0x9A, 0x65, 0x3B, 0x0D9, 0x87, 0x04, 0x5A, 0x0B8, 0x0E6, 0x0A7, 0x0F9, 0x1B, 0x45, 0x0C6, 0x98, 0x7A, 0x24, 0x0F8, 0x0A6, 0x44, 0x1A, 0x99, 0x0C7, 0x25, 0x7B, 0x3A, 0x64, 0x86, 0x0D8, 0x5B, 0x05, 0x0E7, 0x0B9, 0x8C, 0x0D2, 0x30, 0x6E, 0x0ED, 0x0B3, 0x51, 0x0F, 0x4E, 0x10, 0x0F2, 0x0AC, 0x2F, 0x71, 0x93, 0x0CD, 0x11, 0x4F, 0x0AD, 0x0F3, 0x70, 0x2E, 0x0CC, 0x92, 0x0D3, 0x8D, 0x6F, 0x31, 0x0B2, 0x0EC, 0x0E, 0x50, 0x0AF, 0x0F1, 0x13, 0x4D, 0x0CE, 0x90, 0x72, 0x2C, 0x6D, 0x33, 0x0D1, 0x8F, 0x0C, 0x52, 0x0B0, 0x0EE, 0x32, 0x6C, 0x8E, 0x0D0, 0x53, 0x0D, 0x0EF, 0x0B1, 0x0F0, 0x0AE, 0x4C, 0x12, 0x91, 0x0CF, 0x2D, 0x73, 0x0CA, 0x94, 0x76, 0x28, 0x0AB, 0x0F5, 0x17, 0x49, 0x08, 0x56, 0x0B4, 0x0EA, 0x69, 0x37, 0x0D5, 0x8B, 0x57, 0x09, 0x0EB, 0x0B5, 0x36, 0x68, 0x8A, 0x0D4, 0x95, 0x0CB, 0x29, 0x77, 0x0F4, 0x0AA, 0x48, 0x16, 0x0E9, 0x0B7, 0x55, 0x0B, 0x88, 0x0D6, 0x34, 0x6A, 0x2B, 0x75, 0x97, 0x0C9, 0x4A, 0x14, 0x0F6, 0x0A8, 0x74, 0x2A, 0x0C8, 0x96, 0x15, 0x4B, 0x0A9, 0x0F7, 0x0B6, 0x0FC, 0x0A, 0x54, 0x0D7, 0x89, 0x6B, 0x35}
	size := len(Command1) - 1
	Command1[size] = 0
	for i := 0; i <= int(size)-1; i++ {
		Command1[size] = CRCTable[Command1[i]^Command1[size]]
	}
	return Command1
}

func main() {
	count := flag.Int("count", 10, "an int")
	addr := flag.Int("addr", 127, "an int")
	relay := flag.Int("relay", 1, "an int")
	on := flag.Int("on", 1, "an int")
	cmd := flag.Int("cmd", 21, "an int")
	strHex := flag.String("hex", "00", "an string")
	flag.Parse()
	var countUint8 uint8 = uint8(*count)
	var addrUint8 uint8 = uint8(*addr)
	var cmdUint8 uint8 = uint8(*cmd)
	var relayUint8 uint8 = uint8(*relay)
	var onUint8 uint8 = uint8(*on)

	strHexByte, err := hex.DecodeString(*strHex)
	CheckError(err)

	fmt.Printf("\nAddr:%d  Relay:%d Cmd:%d 0x%x  Count:%d\n", addrUint8, relayUint8, cmdUint8, cmdUint8, countUint8)

	// ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:40008")
	// ServerAddr, err := net.ResolveUDPAddr("udp", "192.168.127.2:40008")
	ServerAddr, err := net.ResolveUDPAddr("udp", "192.168.127.254:40000")
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "192.168.127.2:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()
	//after this command,  c2ke-fth start fast send status
	// buf := []byte{0x10, 0x01, 0xff, 0xff, 0xff, 0x7f, 0x06, 0x00, 0x17, 0x00, 0x00, 0x10, 0xff}

	// [ADDR][06h][00h][15h][01h][01h][CRC8] -Включить реле №1
	//
	// [ADDR][06h][00h][15h][01h][00h][CRC8] -Выключить реле №1
	//
	// [ADDR][06h][00h][15h][01h][03h][CRC8] -Включить реле №1 на время
	//
	// [ADDR][06h][00h][15h][02h][01h][CRC8] -Включить реле №2
	//
	// [ADDR][06h][00h][15h][03h][01h][CRC8] -Включить реле №3

	Command1 := []byte{5, 0x06, 0x00, 0x15, 0x01, 0x00, 0xFF}
	Command1[0] = addrUint8
	Command1[3] = cmdUint8
	Command1[4] = relayUint8
	Command1[5] = onUint8

	buf := []byte{0x10, 0x07, 0x01, 0x15, 0x10}
	buf[3] = countUint8
	bufCommand1 := crc8dallas(Command1)
	bufOut := append(buf, bufCommand1...)

	if len(strHexByte) > 1 {
		buf1 := strHexByte[:5]
		buf2 := crc8dallas(strHexByte[5:])
		bufOut = append(buf1, buf2...)
	}
	// buf := []byte{0xed, 0x9c, 0x22, 0x01, 0x14, 0x4d, 0x01, 0xbc, 0x18, 0x00}
	_, err = Conn.Write(bufOut)
	if err != nil {
		fmt.Println(bufOut, err)
	} else {
		fmt.Printf("\n%x\n", bufOut)
	}

	time.Sleep(time.Second * 1)

	bufOut[3] = bufOut[3] + 1

	_, err = Conn.Write(bufOut)
	if err != nil {
		fmt.Println(bufOut, err)
	} else {
		fmt.Printf("\n%x\n", bufOut)
	}

	var icount byte = 1
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text2, _ := reader.ReadString('\n')
		fmt.Println(text2)

		text2 = strings.Replace(text2, " ", "", -1)
		text2 = strings.Trim(text2, "\n")

		strHexByte, err = hex.DecodeString(text2)
		CheckError(err)

		if len(strHexByte) > 1 {
			buf1 := strHexByte[:5]
			buf2 := crc8dallas(strHexByte[5:])
			bufOut = append(buf1, buf2...)
		}

		bufOut[3] = icount

		_, err = Conn.Write(bufOut)
		if err != nil {
			fmt.Println(bufOut, err)
		} else {
			fmt.Printf("\n%x\n", bufOut)
		}
		icount++

	}

	//i := 0
	// for {
	// 	msg := strconv.Itoa(i)
	// 	i++
	// 	buf := []byte(msg)
	// 	_, err := Conn.Write(buf)
	// 	if err != nil {
	// 		fmt.Println(msg, err)
	// 	}
	// 	time.Sleep(time.Second * 1)
	// }
}
