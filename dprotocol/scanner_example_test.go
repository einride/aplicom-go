package dprotocol_test

import (
	"fmt"
	"net"
	"time"

	"go.einride.tech/aplicom/dprotocol"
)

func ExampleScanner() {
	// Bind a TCP listener.
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err) // TODO: Handle error.
	}
	// Accept D protocol connections.
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err) // TODO: Handle error.
		}
		// Scan D protocol packets.
		go func() {
			sc := dprotocol.NewScanner(conn)
			for sc.ScanPacket() {
				fmt.Printf(
					"Unit ID: %d Event ID: %d GPS Time: %s\n",
					sc.Packet().Header.UnitID,
					sc.Packet().EventID,
					sc.Packet().GPSTime.Format(time.RFC3339),
				)
			}
			if sc.Err() != nil {
				panic(err) // TODO: Handle error.
			}
		}()
	}
}
