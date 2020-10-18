SELECT sl_id
	,zap_id
	,order_id
	,"LPU_1"
	,"PODR"
	,"PROFIL"
	,"PROFIL_K"
	,"DET"
	,"P_CEL"
	,"NHISTORY"
	,"P_PER"
	,"DATE_1"
	,"DATE_2"
	,"KD"
	,"DS0"
	,"DS1"
	,"DS2"
	,"DS3"
	,"C_ZAB"
	,"DN"
	,"CODE_MES1"
	,"CODE_MES2"
	,"REAB"
	,"PRVS"
	,"VERS_SPEC"
	,"IDDOKT"
	,"ED_COL"
	,"N_Z"
	,"TARIF"
	,"SUM_M"
	,"COMENTSL"
FROM reestr_export.sl
where order_id = $1
order by zap_id, "DATE_1" ,"DATE_2";
