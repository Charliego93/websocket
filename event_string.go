// Code generated by "enumer -type Event -text -values -json -trimprefix Event -output event_string.go"; DO NOT EDIT.

package websocket

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _EventName = "NormalReadWriteDecompressUnmarshalReconnect"

var _EventIndex = [...]uint8{0, 6, 10, 15, 25, 34, 43}

const _EventLowerName = "normalreadwritedecompressunmarshalreconnect"

func (i Event) String() string {
	if i >= Event(len(_EventIndex)-1) {
		return fmt.Sprintf("Event(%d)", i)
	}
	return _EventName[_EventIndex[i]:_EventIndex[i+1]]
}

func (Event) Values() []string {
	return EventStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _EventNoOp() {
	var x [1]struct{}
	_ = x[EventNormal-(0)]
	_ = x[EventRead-(1)]
	_ = x[EventWrite-(2)]
	_ = x[EventDecompress-(3)]
	_ = x[EventUnmarshal-(4)]
	_ = x[EventReconnect-(5)]
}

var _EventValues = []Event{EventNormal, EventRead, EventWrite, EventDecompress, EventUnmarshal, EventReconnect}

var _EventNameToValueMap = map[string]Event{
	_EventName[0:6]:        EventNormal,
	_EventLowerName[0:6]:   EventNormal,
	_EventName[6:10]:       EventRead,
	_EventLowerName[6:10]:  EventRead,
	_EventName[10:15]:      EventWrite,
	_EventLowerName[10:15]: EventWrite,
	_EventName[15:25]:      EventDecompress,
	_EventLowerName[15:25]: EventDecompress,
	_EventName[25:34]:      EventUnmarshal,
	_EventLowerName[25:34]: EventUnmarshal,
	_EventName[34:43]:      EventReconnect,
	_EventLowerName[34:43]: EventReconnect,
}

var _EventNames = []string{
	_EventName[0:6],
	_EventName[6:10],
	_EventName[10:15],
	_EventName[15:25],
	_EventName[25:34],
	_EventName[34:43],
}

// EventString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func EventString(s string) (Event, error) {
	if val, ok := _EventNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _EventNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Event values", s)
}

// EventValues returns all values of the enum
func EventValues() []Event {
	return _EventValues
}

// EventStrings returns a slice of all String values of the enum
func EventStrings() []string {
	strs := make([]string, len(_EventNames))
	copy(strs, _EventNames)
	return strs
}

// IsAEvent returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Event) IsAEvent() bool {
	for _, v := range _EventValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Event
func (i Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Event
func (i *Event) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Event should be a string, got %s", data)
	}

	var err error
	*i, err = EventString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for Event
func (i Event) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Event
func (i *Event) UnmarshalText(text []byte) error {
	var err error
	*i, err = EventString(string(text))
	return err
}