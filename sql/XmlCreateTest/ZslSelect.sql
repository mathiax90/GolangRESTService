SELECT z_sl_id
	,zap_id
	,order_id
	,"IDCASE"
	,"USL_OK"
	,"VIDPOM"
	,"FOR_POM"
	,"NPR_MO"
	,"NPR_DATE"
	,"LPU"
	,"DATE_Z_1"
	,"DATE_Z_2"
	,"KD_Z"
	,"VNOV_M"
	,"RSLT"
	,"ISHOD"
	,"OS_SLUCH"
	,"VB_P"
	,"IDSP"
	,"SUMV"
	,"OPLATA"
	,"SUMP"
	,"SANK_IT"
FROM reestr_export.z_sl
where order_id = $1;
