package main
import (
    "encoding/xml"
    "gopkg.in/guregu/null.v4"
    //"fmt"
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
    //fmt.Println(nullStringXml)
    return e.EncodeElement(nullStringXml.String, start)
}

type NullIntXml struct {
    null.Int
} 

// MarshalXML generate XML output for PrecsontructedInfo
func (nullIntXml NullIntXml) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
    if nullIntXml.Valid == false {
        return nil
    }

    return e.EncodeElement(nullIntXml.Int64, start)
}

type NullTimeXml struct {
    null.Time
} 

// MarshalXML generate XML output for PrecsontructedInfo
func (nullTimeXml NullTimeXml) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
    if nullTimeXml.Valid == false {
        return nil
    }
//    fmt.Println(nullTimeXml)
    return e.EncodeElement(nullTimeXml.Time.Time.Format("2006-01-02"), start)
}

type NullFloatXml struct {
    null.Float
} 

// MarshalXML generate XML output for PrecsontructedInfo
func (nullFloatXml NullFloatXml) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
    if nullFloatXml.Valid == false {
        return nil
    }

    return e.EncodeElement(nullFloatXml.Float64, start)
}