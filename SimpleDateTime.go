package main
import (
	"time"
	"strings"
	"encoding/json"
	// "fmt"
    // "bytes"
    "encoding/xml"
    "errors"
    "database/sql/driver"
)

type SimpleDate struct {
    Time time.Time    
}

func (simpleDate SimpleDate) MarshalJSON() ([]byte, error) {        
    return json.Marshal(simpleDate.Time.Format("2006-01-02"))
}

func (simpleDate *SimpleDate)UnmarshalJSON(in []byte) error {   
    inStr := strings.Trim(string(in), `"`)    
    var err error 
    simpleDate.Time, err = time.Parse("2006-01-02", inStr)    
    return err
}

// MarshalXML generate XML output for PrecsontructedInfo
func (simpleDate SimpleDate) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
    return e.EncodeElement(simpleDate.Time.Format("2006-01-02"), start)
}

func (simpleDate *SimpleDate) Scan(value interface{}) error {    
    if value == nil {
        return errors.New("SimpleDate can't Scan null use NullSimpleDate instead")
    } else {
        if DateAsString, err := driver.String.ConvertValue(value); err == nil {            
            if v, ok := DateAsString.(string); ok {                
                simpleDate.Time, err = time.Parse("2006-01-02 15:04:05 -0700 MST", string(v))
                if err != nil {
                    return err
                }
                return nil
            }
        }
        return errors.New("failed to scan SimpleDate")
    }
}
