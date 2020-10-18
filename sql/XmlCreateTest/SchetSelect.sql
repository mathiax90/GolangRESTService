SELECT schet_id
	,order_id
	,"VERSION"
	,"DATA"
	,"FILENAME"
	,"SD_Z"
	,"CODE"
	,"CODE_MO"
	,"YEAR"
	,"MONTH"
	,"NSCHET"
	,"DSCHET"
	,"PLAT"
	,"SUMMAV"
	-- ,"COMENTS"
	-- ,"SUMMAP"
	-- ,"SANK_MEK"
	-- ,"SANK_MEE"
	-- ,"SANK_EKMP"
FROM reestr_export.schet
where order_id = $1;
