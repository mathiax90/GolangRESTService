package main

import (
	"fmt"
	//"log"
	
	//"strconv"
//"encoding/json"
		
	//"context"	
	"os"
	//pgx "github.com/jackc/pgx/v4"	
	//"github.com/jmoiron/sqlx"
	//"github.com/joho/godotenv"
	_ "github.com/jackc/pgx/stdlib"
	//"gopkg.in/guregu/null.v4"
	//"encoding/xml"
	"database/sql"
	"time"
	"io/ioutil"
	//"gopkg.in/guregu/null.v4"
	"encoding/xml"
)
type Root struct {
	Schet Schet
	ZapCollection ZapCollection
}

type Schet struct {
	XMLName   xml.Name `xml:"SCHET"`
	SchetId string  `db:"schet_id" xml:"-"`
	OrderId sql.NullString  `db:"order_id" xml:"-"`
	VERSION string  `db:"VERSION"`
	DATA SimpleDate  `db:"DATA"`
	FILENAME string  `db:"FILENAME"`
	SD_Z int64  `db:"SD_Z"`
	CODE int64  `db:"CODE"`
	CODE_MO string  `db:"CODE_MO"`
	YEAR int64  `db:"YEAR"`
	MONTH int64  `db:"MONTH"`
	NSCHET string  `db:"NSCHET"`
	DSCHET time.Time  `db:"DSCHET"`
	PLAT sql.NullString  `db:"PLAT"`
	SUMMAV float64  `db:"SUMMAV"`
	// COMENTS sql.NullString  `db:"COMENTS"`
	// SUMMAP sql.NullInt64  `db:"SUMMAP"`
	// SANK_MEK sql.NullInt64  `db:"SANK_MEK"`
	// SANK_MEE sql.NullInt64  `db:"SANK_MEE"`
	// SANK_EKMP sql.NullInt64  `db:"SANK_EKMP"`
}

type ZapCollection struct {
	Zaps []Zap
}

type Zap struct {
	ZapId string  `db:"zap_id"`
	OrderId sql.NullString  `db:"order_id"`
	N_ZAP int64  `db:"N_ZAP"`
	PR_NOV int64  `db:"PR_NOV"`
	ID_PAC string  `db:"ID_PAC"`
	VPOLIS int64  `db:"VPOLIS"`
	SPOLIS sql.NullString  `db:"SPOLIS"`
	NPOLIS string  `db:"NPOLIS"`
	ST_OKATO sql.NullString  `db:"ST_OKATO"`
	SMO sql.NullString  `db:"SMO"`
	SMO_OGRN sql.NullString  `db:"SMO_OGRN"`
	SMO_OK sql.NullString  `db:"SMO_OK"`
	SMO_NAM sql.NullString  `db:"SMO_NAM"`
	INV sql.NullInt64  `db:"INV"`
	MSE sql.NullInt64  `db:"MSE"`
	NOVOR string  `db:"NOVOR"`
	VNOV_D sql.NullInt64  `db:"VNOV_D"`
}


func XmlCreateTest() {
		
	fmt.Println("start")
	var err error	
	order_id := "6424858d-1d3e-4e23-ae1b-78e65c1f3e19"

	//schet
	Schet := Schet{}
	SchetSelectSqlFile, err := os.Open("./sql/XmlCreateTest/SchetSelect.sql")	
	if err != nil {
		fmt.Println(err)
	}
	defer SchetSelectSqlFile.Close()
	byteSql, _ := ioutil.ReadAll(SchetSelectSqlFile)
	SchetSelectSql := string(byteSql)

	err = db.Get(&Schet,SchetSelectSql,order_id)
	if err != nil {
		fmt.Printf("error while select reestr from db: %v\n", err)
		return
	}

	//zaps

	ZapCollection:= ZapCollection{}
	Zaps := []Zap{}
	ZapCollection.Zaps = Zaps

	ZapSelectSqlFile, err := os.Open("./sql/XmlCreateTest/ZapSelect.sql")	
	if err != nil {
		fmt.Println(err)
	}
	defer ZapSelectSqlFile.Close()
	byteSql, _ = ioutil.ReadAll(ZapSelectSqlFile)
	ZapSelectSql := string(byteSql)

	err = db.Select(&Zaps,ZapSelectSql,order_id)
	if err != nil {
		fmt.Printf("error while select reestr from db: %v\n", err)
		return
	}

	fmt.Println(Zaps)
	// for _, v := range ReestrCollection.Reestrs {
	// 	fmt.Println(v)
	// 	if !ZapCollection.Contains(v.ID_PAC){
	// 		fmt.Println("")
	// 	}
		
	// 	// if zaps

	Root := Root{}
	Root.Schet = Schet
	Root.ZapCollection = ZapCollection

	// 	break
	// }
	output, err := xml.MarshalIndent(Root, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
}



// MarshalXML generate XML output for PrecsontructedInfo
// func (reestr Reestr) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
// 	zap := []Zap{}
// 	for i, v := range reestr.Row {
		
// 	}

//     return e.EncodeElement(reestr, start)
// }
