package icapclient

import (
	"context"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// transport represents the transport layer data
type transport struct {
	network      string
	addr         string
	timeout      time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
	sckt         net.Conn
}

// dial fires up a tcp socket
func (t *transport) dial() error {
	sckt, err := net.DialTimeout(t.network, t.addr, t.timeout)

	if err != nil {
		return err
	}

	if err := sckt.SetReadDeadline(time.Now().UTC().Add(t.readTimeout)); err != nil {
		return err
	}

	if err := sckt.SetWriteDeadline(time.Now().UTC().Add(t.writeTimeout)); err != nil {
		return err
	}

	t.sckt = sckt

	return nil
}

// dialWithContext fires up a tcp socket
func (t *transport) dialWithContext(ctx context.Context) error {
	sckt, err := (&net.Dialer{
		Timeout: t.timeout,
	}).DialContext(ctx, t.network, t.addr)

	if err != nil {
		return err
	}

	if err := sckt.SetReadDeadline(time.Now().UTC().Add(t.readTimeout)); err != nil {
		return err
	}

	if err := sckt.SetWriteDeadline(time.Now().UTC().Add(t.writeTimeout)); err != nil {
		return err
	}

	t.sckt = sckt

	return nil
}

// Write writes data to the server
func (t *transport) write(data []byte) (int, error) {
	return t.sckt.Write(data)
}

// Read reads data from server
func (t *transport) read() (string, error) {

	data := make([]byte, 0)

	if DEBUG {
		log.Println("Dumping messages received from the socket...")
	}

	for {
		tmp := make([]byte, 1096)

		n, err := t.sckt.Read(tmp)

		if err != nil {
			if err == io.EOF {
				if DEBUG {
					log.Println("End of file detected from EOF error")
				}
				break
			}
			return "", err
		}

		if n == 0 {
			if DEBUG {
				log.Println("End of file detected by 0 bytes")
			}
			break
		}

		data = append(data, tmp[:n]...)
		if string(data) == icap100ContinueMsg { // explicitly breaking because the Read blocks for 100 continue message // TODO: find out why
			break
		}

		if strings.HasSuffix(string(data), "0\r\n\r\n") {
			if DEBUG {
				log.Println("End of the file detected by 0 Double CRLF indicator")
			}
			break
		}

		if strings.Contains(string(data), icap204NoModsMsg){
			if DEBUG {
				log.Println("End of file detected by 204 no modifications and Double CRLF at the end")
			}
			break
		}

		if DEBUG {
			spew.Dump(tmp)
		}

	}

	return string(data), nil
}

// close closes the tcp connection
func (t *transport) close() error {
	return t.sckt.Close()
}
