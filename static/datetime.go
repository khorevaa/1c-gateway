package odata

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Type for Edm.DateTime
type DateTime time.Time

const OneCTimeFormat = "2006-01-02T15:04:05"
const JSTimeFormat = "2006-01-02T15:04:05Z0700"

// Maps DateTime to the graphql scalar type in the schema.
func (DateTime) ImplementsGraphQLType(name string) bool {
	return name == "DateTime"
}

// A custom graphql unmarshaler for DateTime type
func (t *DateTime) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case time.Time:
		*t = DateTime(input)
		return nil
	case string:
		val, err := time.Parse(OneCTimeFormat, input)
		if err != nil {
			val, err = time.Parse(JSTimeFormat, input)
		}
		*t = DateTime(val)
		return err
	case int:
		val := time.Unix(int64(input), 0)
		*t = DateTime(val)
		return nil
	case float64:
		val := time.Unix(int64(input), 0)
		*t = DateTime(val)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

// A custom json/graphql marshaller for DateTime type
func (t DateTime) MarshalJSON() ([]byte, error) {
	val := time.Time(t).Format(JSTimeFormat)
	return json.Marshal(val)
}

// A custom json unmarshaller for DateTime type
func (t *DateTime) UnmarshalJSON(b []byte) error {
	val, err := time.Parse(JSTimeFormat, strings.Trim(string(b), `"`))
	*t = DateTime(val)
	return err
}

// A custom marshaller to uri query format for DateTime type
func (t DateTime) AsParameter() string {
	val := time.Time(t).Format(OneCTimeFormat)
	return fmt.Sprintf("datetime'%s'", val)
}
