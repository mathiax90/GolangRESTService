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
	//"database/sql"
	//"time"
	"io/ioutil"
	//"gopkg.in/guregu/null.v4"
	"encoding/xml"
)
type ZL_LIST struct {
	XMLName xml.Name `xml:"ZL_LIST"`
	SCHET SCHET
	ZAPs []ZAP	
}

type SCHET struct {
	XMLName xml.Name `xml:"SCHET"`
	SchetId string  `db:"schet_id" xml:"-"`
	OrderId NullStringXml  `db:"order_id" xml:"-"`
	VERSION string  `db:"VERSION"`
	DATA SimpleDate  `db:"DATA"`
	FILENAME string  `db:"FILENAME"`
	SD_Z int64  `db:"SD_Z"`
	CODE int64  `db:"CODE"`
	CODE_MO string  `db:"CODE_MO"`
	YEAR int64  `db:"YEAR"`
	MONTH int64  `db:"MONTH"`
	NSCHET string  `db:"NSCHET"`
	DSCHET SimpleDate  `db:"DSCHET"`
	PLAT NullStringXml  `db:"PLAT"`
	SUMMAV float64  `db:"SUMMAV"`
	// COMENTS sql.NullString  `db:"COMENTS"`
	// SUMMAP sql.NullInt64  `db:"SUMMAP"`
	// SANK_MEK sql.NullInt64  `db:"SANK_MEK"`
	// SANK_MEE sql.NullInt64  `db:"SANK_MEE"`
	// SANK_EKMP sql.NullInt64  `db:"SANK_EKMP"`
}

type ZAP struct {
	XMLName xml.Name `xml:"ZAP"`
	ZapId string  `db:"zap_id"  xml:"-"`
	OrderId NullStringXml  `db:"order_id"  xml:"-"`
	N_ZAP int64  `db:"N_ZAP"`
	PR_NOV int64  `db:"PR_NOV"`
	ID_PAC string  `db:"ID_PAC"`
	VPOLIS int64  `db:"VPOLIS"`
	SPOLIS NullStringXml  `db:"SPOLIS"`
	NPOLIS string  `db:"NPOLIS"`
	ST_OKATO NullStringXml  `db:"ST_OKATO"`
	SMO NullStringXml  `db:"SMO"`
	SMO_OGRN NullStringXml  `db:"SMO_OGRN"`
	SMO_OK NullStringXml  `db:"SMO_OK"`
	SMO_NAM NullStringXml  `db:"SMO_NAM"`
	INV NullIntXml  `db:"INV"`
	MSE NullIntXml  `db:"MSE"`
	NOVOR string  `db:"NOVOR"`
	VNOV_D NullIntXml  `db:"VNOV_D"`
	Z_SL Z_SL
}

type Z_SL struct {
	XMLName xml.Name `xml:"Z_SL"`
	ZslId string  `db:"z_sl_id" xml:"-"`
	ZapId string  `db:"zap_id" xml:"-"`
	OrderId NullStringXml  `db:"order_id" xml:"-"`
	IDCASE int64  `db:"IDCASE"`
	USL_OK int64  `db:"USL_OK"`
	VIDPOM int64  `db:"VIDPOM"`
	FOR_POM int64  `db:"FOR_POM"`
	NPR_MO NullStringXml  `db:"NPR_MO"`
	NPR_DATE NullTimeXml  `db:"NPR_DATE"`
	LPU string  `db:"LPU"`
	DATE_Z_1 SimpleDate  `db:"DATE_Z_1"`
	DATE_Z_2 SimpleDate  `db:"DATE_Z_2"`
	KD_Z NullIntXml  `db:"KD_Z"`
	VNOV_M NullIntXml  `db:"VNOV_M"`
	RSLT int64  `db:"RSLT"`
	ISHOD int64  `db:"ISHOD"`
	OS_SLUCH NullIntXml  `db:"OS_SLUCH"`
	VB_P NullIntXml  `db:"VB_P"`
	IDSP int64  `db:"IDSP"`
	SUMV NullFloatXml  `db:"SUMV"`
	OPLATA NullIntXml  `db:"OPLATA"`
	SUMP NullFloatXml  `db:"SUMP"`
	SANK_IT NullFloatXml  `db:"SANK_IT"`
}


func XmlCreateTest() {
		
	fmt.Println("start")
	var err error	
	order_id := "6424858d-1d3e-4e23-ae1b-78e65c1f3e19"

	//schet
	SCHET := SCHET{}
	SchetSelectSqlFile, err := os.Open("./sql/XmlCreateTest/SchetSelect.sql")	
	if err != nil {
		fmt.Println(err)
	}
	defer SchetSelectSqlFile.Close()
	byteSql, _ := ioutil.ReadAll(SchetSelectSqlFile)
	SchetSelectSql := string(byteSql)

	err = db.Get(&SCHET,SchetSelectSql,order_id)
	if err != nil {
		fmt.Printf("error while select reestr from db: %v\n", err)
		return
	}

	//zaps

	
	ZAPs := []ZAP{}	
	Z_SLs := []Z_SL{}
	SLs := []SL{}
	ZapSelectSqlFile, err := os.Open("./sql/XmlCreateTest/ZapSelect.sql")	
	if err != nil {
		fmt.Println(err)
	}
	defer ZapSelectSqlFile.Close()
	byteSql, _ = ioutil.ReadAll(ZapSelectSqlFile)
	ZapSelectSql := string(byteSql)

	err = db.Select(&ZAPs,ZapSelectSql,order_id)
	if err != nil {
		fmt.Printf("error while select reestr from db: %v\n", err)
		return
	}

	ZslSelectSqlFile, err := os.Open("./sql/XmlCreateTest/ZslSelect.sql")	
	if err != nil {
		fmt.Println(err)
	}
	defer ZslSelectSqlFile.Close()
	byteSql, _ = ioutil.ReadAll(ZslSelectSqlFile)
	ZslSelectSql := string(byteSql)

	err = db.Select(&Z_SLs,ZslSelectSql,order_id)
	if err != nil {
		fmt.Printf("error while select reestr from db: %v\n", err)
		return
	}

	//assign z_sl to zap
	for _, zsl := range Z_SLs {
		for _, zap := range ZAPs {
			if zsl.ZapId == zap.ZapId{
				zap.Z_SL = zsl
			}
		}
	}


	ZL_LIST := ZL_LIST{}
	ZL_LIST.SCHET = SCHET
	ZL_LIST.ZAPs = ZAPs


	output, err := xml.MarshalIndent(ZL_LIST, "  ", "    ")
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
