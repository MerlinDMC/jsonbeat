package beater

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/MerlinDMC/jsonbeat/config"
)

type Jsonbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client

	listener net.Listener
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Jsonbeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Jsonbeat) Run(b *beat.Beat) error {
	logp.Info("jsonbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	bt.listener, err = net.Listen("tcp", bt.config.Address)
	if err != nil {
		logp.Err("Error listening on %s: %v", bt.config.Address, err.Error())
		return err
	}
	defer bt.listener.Close()

	logp.Info("Now listening for logs on %s", bt.config.Address)

	for {
		select {
		case <-bt.done:
			return nil
		default:
		}

		conn, err := bt.listener.Accept()
		if err != nil {
			logp.Err("Error accepting connection: %v", err.Error())
			continue
		}

		go bt.handleConnection(conn)
	}
}

func (bt *Jsonbeat) Stop() {
	bt.client.Close()
	bt.listener.Close()
	close(bt.done)
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func (bt *Jsonbeat) handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, '\n'); i >= 0 {
			return i + 1, dropCR(data[0:i]), nil
		}
		return 0, nil, nil
	})

	for scanner.Scan() {
		var data common.MapStr

		if err := json.Unmarshal(scanner.Bytes(), &data); err != nil {
			logp.Err("Error decoding event: %v", err.Error())
			logp.Debug("Data: %v", scanner.Text())

			continue
		}

		var timestamp time.Time = time.Now()

		if bt.config.TimestampField != "" {
			if ts, err := data.GetValue(bt.config.TimestampField); err == nil {
				if ts, ok := ts.(int64); ok {
					timestamp = time.Unix(ts, 0)
				}

				if ts, ok := ts.(float64); ok {
					timestamp = time.Unix(int64(ts), 0)
				}

				if ts, ok := ts.(string); ok {
					if ts, err := time.Parse(bt.config.TimestampFormat, ts); err == nil {
						timestamp = ts
					}
				}

				data.Delete(bt.config.TimestampField)
			}
		}

		if len(bt.config.Namespace) > 0 {
			for i := len(bt.config.Namespace) - 1; i >= 0; i-- {
				data = common.MapStr{
					bt.config.Namespace[i]: data,
				}
			}
		}

		event := beat.Event{
			Timestamp: timestamp,
			Fields:    data,
		}

		bt.client.Publish(event)
	}
}
