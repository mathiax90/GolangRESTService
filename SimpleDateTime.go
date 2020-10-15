package main
import (
	"time"
	"strings"
	//"encoding/json"
	//"fmt"
    "bytes"
    "encoding/xml"
)

type SimpleDate struct {
    Time time.Time    
}

func (simpleDate SimpleDate) MarshalJSON() ([]byte, error) {
    var buf bytes.Buffer    
    buf.WriteString(simpleDate.Time.Format("2006-01-02"))
    return buf.Bytes(), nil
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


// func (simpleDate *SimpleDate) Scan(value interface{}) error {

// }
