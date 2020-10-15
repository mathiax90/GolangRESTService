package main
import (
    "encoding/xml"
    "gopkg.in/guregu/null.v4"
    "fmt"
    //"time"
)


type NullStringXml struct {
    null.String
} 

// MarshalXML generate XML output for PrecsontructedInfo
func (nullStringXml NullStringXml) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
    if nullStringXml.Valid == false {
        return nil        
    }

    return e.EncodeElement(nullStringXml, start)
}

type NullIntXml struct {
    null.Int
} 

// MarshalXML generate XML output for PrecsontructedInfo
func (nullIntXml NullIntXml) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
    if nullIntXml.Valid == false {
        return nil
    }

    return e.EncodeElement(nullIntXml, start)
}

type NullTimeXml struct {
    null.Time
} 

// MarshalXML generate XML output for PrecsontructedInfo
func (nullTimeXml NullTimeXml) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
    if nullTimeXml.Valid == false {
        return nil
    }
    fmt.Println(nullTimeXml)
    //var tempTime time.Time
    //tempTime = nullTimeXml.Time.Time
    return e.EncodeElement(nullTimeXml.Time.Time.Format("2006-01-02"), start)
}