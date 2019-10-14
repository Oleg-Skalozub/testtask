package entity

const (
	// EventType ...
	EventType = iota + 1
	BirthType
	DeathType
)

const (
	// EventTypeName ...
	EventTypeName = "Events"
	BirthTypeName = "Births"
	DeathTypeName = "Deaths"
)

// TableName is name of table with Events
const TableName = "data_dbs"

// TaskingNameTypeMap map with event type name
var TaskingNameTypeMap = map[int]string{
	EventType: EventTypeName,
	BirthType: BirthTypeName,
	DeathType: DeathTypeName,
}

// TaskingTypeNameMap map with event type const
var TaskingTypeNameMap = map[string]int{
	EventTypeName: EventType,
	BirthTypeName: BirthType,
	DeathTypeName: DeathType,
}

// Contain ...
type Contain struct {
	Date string `json:"date"`
	Data Data   `json:"data"`
}

// Data
type Data struct {
	Events []Event `json:"Events"`
	Births []Event `json:"Births"`
	Deaths []Event `json:"Deaths"`
}

// Event
type Event struct {
	Year string `json:"year"`
	Text string `json:"text"`
}

// Link
type Link struct {
	Title string
}

// DataDB
type DataDB struct {
	Month     int
	Day       int
	Year      string
	Title     string
	EventType int
}

type DataResponse struct {
	EventType string `json:"event_type"`
	Result    int    `json:"result"`
}
