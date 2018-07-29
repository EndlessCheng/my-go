package main

import (
	"bufio"
	"io"
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readAll(r io.Reader) (data string, err error) {
	buf := make([]byte, 4096)
	for {
		n, err := r.Read(buf)
		switch err {
		case nil:
			data += string(buf[:n])
			if strings.HasSuffix(data, "\r\n") {
				// for POST form file
				n, err = r.Read(buf)
				if err == nil {
					data += string(buf[:n])
				}
				return data, nil
			}
		case io.EOF:
			return data, nil
		default:
			return "", err
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3334")
	checkError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		checkError(err)

		go func() {
			defer conn.Close()

			defer func() {
				err := recover()
				if err != nil {
					log.Warn(err)
				}
			}()

			// read
			r := bufio.NewReader(conn)
			data, err := readAll(r)
			checkError(err)

			log.Info(data)

			// write
			w := bufio.NewWriter(conn)
			_, err = w.Write([]byte(`HTTP/1.1 200 OK

{"msg":"DONE"}`))
			checkError(err)
			err = w.Flush()
			checkError(err)
		}()
	}
}
