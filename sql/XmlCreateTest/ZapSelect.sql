SELECT zap_id
	,order_id
	,"PR_NOV"
	,"ID_PAC"
	,"VPOLIS"
	,"SPOLIS"
	,"NPOLIS"
	,"ST_OKATO"
	,"SMO"
	,"SMO_OGRN"
	,"SMO_OK"
	,"SMO_NAM"
	,"INV"
	,"MSE"
	,"NOVOR"
	,"VNOV_D"
FROM reestr_export.zap
where order_id = $1;
