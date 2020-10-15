package main

import (
	"fmt"
	//"log"
	
	//"strconv"
//"encoding/json"
		
	//"context"	
	//"os"
	//pgx "github.com/jackc/pgx/v4"	
	//"github.com/jmoiron/sqlx"
	//"github.com/joho/godotenv"
	_ "github.com/jackc/pgx/stdlib"
	//"gopkg.in/guregu/null.v4"
	//"encoding/xml"
	"database/sql"
	//"time"
	//"gopkg.in/guregu/null.v4"
)

type Reestr struct{
	Id sql.NullString  `db:"id"`
	Datein sql.NullTime  `db:"datein"`
	Dateout sql.NullTime  `db:"dateout"`
	Ds sql.NullString  `db:"ds"`
	Codeusl sql.NullString  `db:"codeusl"`
	Codemd sql.NullString  `db:"codemd"`
	KolUsl sql.NullInt64  `db:"kol_usl"`
	Idsluch sql.NullString  `db:"idsluch"`
	Prvs sql.NullString  `db:"prvs"`
	Otdel sql.NullString  `db:"otdel"`
	Lpu sql.NullString  `db:"lpu"`
	UslOk sql.NullString  `db:"usl_ok"`
	Profil sql.NullString  `db:"profil"`
	Ds1 sql.NullString  `db:"ds1"`
	Rslt sql.NullString  `db:"rslt"`
	Ishod sql.NullString  `db:"ishod"`
	Nprmo sql.NullString  `db:"nprmo"`
	PCel sql.NullString  `db:"p_cel"`
	Nprdate sql.NullTime  `db:"nprdate"`
	Idpac sql.NullString  `db:"idpac"`
	Fam sql.NullString  `db:"fam"`
	Nam sql.NullString  `db:"nam"`
	Ot sql.NullString  `db:"ot"`
	Dr sql.NullTime  `db:"dr"`
	Npolis sql.NullString  `db:"npolis"`
	Spolis sql.NullString  `db:"spolis"`
	Smo sql.NullString  `db:"smo"`
	Vpolis sql.NullString  `db:"vpolis"`
	W sql.NullString  `db:"w"`
	Iddokt sql.NullString  `db:"iddokt"`
	IdSluch sql.NullString  `db:"id_sluch"`
	Date1 sql.NullTime  `db:"date_1"`
	Date2 sql.NullTime  `db:"date_2"`
	Tarif sql.NullFloat64  `db:"tarif"`
	Unit sql.NullBool  `db:"unit"`
	NZub sql.NullInt64  `db:"n_zub"`
	Koekp sql.NullString  `db:"koekp"`
	Ksg sql.NullString  `db:"ksg"`
	KZat sql.NullString  `db:"k_zat"`
	Telephone sql.NullString  `db:"telephone"`
	TypeDiagn sql.NullString  `db:"type_diagn"`
	Det sql.NullInt64  `db:"det"`
	VidVme sql.NullString  `db:"vid_vme"`
	VidKsg sql.NullString  `db:"vid_ksg"`
	VidKz sql.NullString  `db:"vid_kz"`
	VidVmeUsl sql.NullString  `db:"vid_vme_usl"`
	Lpu1 sql.NullString  `db:"lpu1"`
}

func XmlCreateTest() {
		
	fmt.Println("start")
	var err error
	reestr := []Reestr{}
	err = db.Select(&reestr,"select id, datein, dateout, ds, codeusl, codemd, kol_usl, prvs, idsluch, otdel, lpu, usl_ok, profil, ds1, rslt, ishod, nprmo, p_cel, nprdate, idpac, fam, nam, ot, dr, npolis, spolis, smo, vpolis, w, iddokt, id_sluch, date_1, date_2, tarif, unit, n_zub, koekp, ksg, k_zat, telephone, type_diagn, det, vid_vme, vid_ksg, vid_kz, vid_vme_usl, lpu1 from reestr('2020-08-01','2020-08-31','032155',false)")
	if err != nil {
		fmt.Printf("error while select reestr from db: %v\n", err)
		return
	}	
	fmt.Println(reestr)
	// output, err := xml.MarshalIndent(root, "  ", "    ")
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }

	// os.Stdout.Write(output)
}

