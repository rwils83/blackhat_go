package main
import (
	"fmt"
	"net"
)
func main() {
	for i := 1; i <= 22; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			//port is closed or filtered
			fmt.Printf("Port %d closed\n", i)
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}