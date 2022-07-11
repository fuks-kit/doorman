package rfid

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"
)

//
// Constants are defined by:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/input-event-codes.h
//

const (
	evKey    = 0x01 // EvKey indicates a normal keystroke
	keyEnter = 28   // Key code for enter
)

type InputEvent struct {
	Time  time.Time `json:"time"`
	Type  uint16    `json:"type"`
	Code  uint16    `json:"code"`
	Value int32     `json:"value"`
}

type Device struct {
	Input *os.File
}

func Reader(dev string) Device {
	inputDevice, err := os.Open(dev)
	if err != nil {
		log.Fatalln(err)
	}

	return Device{
		Input: inputDevice,
	}
}

func (device Device) Close() (_ error) {
	return device.Input.Close()
}

// ReadEvents streams parses and streams all inputs from Device.Input
func (device Device) ReadEvents() <-chan InputEvent {
	queue := make(chan InputEvent)

	go func() {
		buf := make([]byte, 16)

		for {
			_, err := device.Input.Read(buf)
			if err != nil {
				log.Fatalln(err)
			}

			//
			// For help see:
			// https://www.kernel.org/doc/Documentation/input/input.txt
			//

			sec := binary.LittleEndian.Uint32(buf[0:8])
			ts := time.Unix(int64(sec), 0)

			typ := binary.LittleEndian.Uint16(buf[8:10])
			code := binary.LittleEndian.Uint16(buf[10:12])

			var value int32
			err = binary.Read(bytes.NewReader(buf[12:]), binary.LittleEndian, &value)
			if err != nil {
				log.Fatalln(err)
			}

			event := InputEvent{
				Time:  ts,
				Type:  typ,
				Code:  code,
				Value: value,
			}

			queue <- event
		}
	}()

	return queue
}

// ReadIdentifiers assembles RFIDs from InputEvents.
func (device Device) ReadIdentifiers() <-chan uint32 {

	queue := make(chan uint32)

	go func() {
		for {
			var input string

			for event := range device.ReadEvents() {
				// 'value' is the value the event carries. Either a relative change for
				// EV_REL, absolute new value for EV_ABS (joysticks ...), or 0 for EV_KEY for
				// release, 1 for keypress and 2 for autorepeat.
				if event.Value != evKey {
					continue
				}

				if event.Code == keyEnter {
					num, err := strconv.ParseUint(input, 10, 64)
					if err != nil {
						//
						// Then too many events are triggered the input can be corrupted.
						// In this case reset the input and continue.
						//

						// log.Printf("error parsing input number: %v", err)
						input = ""
						continue
					}

					input = ""
					queue <- uint32(num)

					continue
				}

				var number string

				// See for key details:
				// https://github.com/torvalds/linux/blob/master/include/uapi/linux/input-event-codes.h
				switch event.Code {
				case 2:
					number = "1"
				case 3:
					number = "2"
				case 4:
					number = "3"
				case 5:
					number = "4"
				case 6:
					number = "5"
				case 7:
					number = "6"
				case 8:
					number = "7"
				case 9:
					number = "8"
				case 10:
					number = "9"
				case 11:
					number = "0"
				default:
					out, _ := json.MarshalIndent(event, "", "  ")
					log.Fatalf("Unknown event code %s", out)
				}

				input += number
			}
		}
	}()

	return queue
}
