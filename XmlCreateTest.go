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
	SL []SL `xml:"SL"`
}

type SL struct {
	XMLName xml.Name `xml:"SL"`
	SlId string  `db:"sl_id" xml:"-"`
	ZapId string  `db:"zap_id" xml:"-"` 
	OrderId string  `db:"order_id" xml:"-"`
	LPU_1 NullStringXml  `db:"LPU_1"`
	PODR NullStringXml  `db:"PODR"`
	PROFIL int64  `db:"PROFIL"`
	PROFIL_K NullIntXml  `db:"PROFIL_K"`
	DET int64  `db:"DET"`
	P_CEL NullStringXml  `db:"P_CEL"`
	NHISTORY string  `db:"NHISTORY"`
	P_PER NullIntXml  `db:"P_PER"`
	DATE_1 SimpleDate  `db:"DATE_1"`
	DATE_2 SimpleDate  `db:"DATE_2"`
	KD NullIntXml  `db:"KD"`
	DS0 NullStringXml  `db:"DS0"`
	DS1 string  `db:"DS1"`
	DS2 NullStringXml  `db:"DS2"`
	DS3 NullStringXml  `db:"DS3"`
	C_ZAB NullIntXml  `db:"C_ZAB"`
	DN NullIntXml  `db:"DN"`
	CODE_MES1 NullStringXml  `db:"CODE_MES1"`
	CODE_MES2 NullStringXml  `db:"CODE_MES2"`
	REAB NullIntXml  `db:"REAB"`
	PRVS int64  `db:"PRVS"`
	VERS_SPEC string  `db:"VERS_SPEC"`
	IDDOKT string  `db:"IDDOKT"`
	ED_COL NullFloatXml  `db:"ED_COL"`
	N_Z NullIntXml  `db:"N_Z"`
	TARIF NullFloatXml  `db:"TARIF"`
	SUM_M float64  `db:"SUM_M"`
	COMENTSL NullStringXml  `db:"COMENTSL"`
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

	SlSelectSqlFile, err := os.Open("./sql/XmlCreateTest/SlSelect.sql")	
	if err != nil {
		fmt.Println(err)
	}
	defer SlSelectSqlFile.Close()
	byteSql, _ = ioutil.ReadAll(SlSelectSqlFile)
	SlSelectSql := string(byteSql)

	err = db.Select(&SLs,SlSelectSql,order_id)
	if err != nil {
		fmt.Printf("error while select reestr from db: %v\n", err)
		return
	}

	SLs_temp := SLs

	//assign sl to z_sl
		for i:= 0;i<len(Z_SLs);i++{			
			for sl_index, sl := range SLs_temp {
				if Z_SLs[i].ZapId == sl.ZapId{
					//zsl.SLs = append(zsl.SLs, SLs[0])		
					Z_SLs[i].SL = append(Z_SLs[i].SL,sl)			
					SLs_temp[sl_index] = SLs_temp[len(SLs_temp)-1] // Copy last element to index i.				
					SLs_temp = SLs_temp[:len(SLs_temp)-1]   // Truncate slice.									
					break;				
				}
			}	
		}
	
	Z_SLs_temp := Z_SLs
	//assign z_sl to zap
	for i:= 0;i<len(ZAPs);i++{	
		for zsl_index, zsl := range Z_SLs_temp {
			if ZAPs[i].ZapId == zsl.ZapId{
				ZAPs[i].Z_SL = zsl				
				Z_SLs_temp[zsl_index] = Z_SLs_temp[len(Z_SLs_temp)-1] // Copy last element to index i.				
				Z_SLs_temp= Z_SLs_temp[:len(Z_SLs_temp)-1]   // Truncate slice.				
				break;				
			}
		}
	}
			


	fmt.Println(Z_SLs[0].SL)
	fmt.Println(ZAPs[0].Z_SL.SL)
	ZL_LIST := ZL_LIST{}
	ZL_LIST.SCHET = SCHET
	ZL_LIST.ZAPs = ZAPs


	output, err := xml.MarshalIndent(ZL_LIST, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	_ = output
	os.Stdout.Write(output)
}



// MarshalXML generate XML output for PrecsontructedInfo
// func (reestr Reestr) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
// 	zap := []Zap{}
// 	for i, v := range reestr.Row {
		
// 	}

//     return e.EncodeElement(reestr, start)
// }
