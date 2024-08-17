package zentao_models

import (
	"net/url"
	"wzDataCenter/app/zentao/zentao_common"
)

// GetAnalysisLeixing 按照用户、起止时间获取‘类型’数据
func GetAnalysisLeixing(userId uint, dateStart string, dateEnd string) (bool, *[]Leixing, int64) {
	var d1 []Leixing
	sql := `SELECT tycc cloudname,sum(esti) esti,sum(cons) cons
			  FROM k_sum
			 WHERE dat BETWEEN ? AND ?
			   AND acct = ?
			 GROUP BY tycc`
	res := zentao_common.ZENTAO_DB.Raw(sql, dateStart, dateEnd, userId)
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}

// GetAnalysisCustomer 按照用户、起止时间获取‘类型’数据
func GetAnalysisCustomer(userId uint, dateStart string, dateEnd string) (bool, *[]Customer, int64) {
	var d1 []Customer
	sql := `SELECT proj customername,sum(esti) esti,sum(cons) cons
          	  FROM k_sum
          	 WHERE dat BETWEEN ? AND ?
          	   and acct = ?
          	 GROUP BY proj`
	res := zentao_common.ZENTAO_DB.Raw(sql, dateStart, dateEnd, userId)
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}

// GetAnalysisCustomerDetail 按照用户、起止时间、类型获取详细客户数据
func GetAnalysisCustomerDetail(userId uint, type0 string, dateStart string, dateEnd string) (bool, *[]CustomerDetail, int64) {
	var d1 []CustomerDetail
	sql := `SELECT proj customername,id,namec titlename,dat workdate,esti,cons
			  FROM k_sum
			 WHERE clid = ?
			   AND acct = ?
			   AND dat BETWEEN ? AND ?`
	res := zentao_common.ZENTAO_DB.Raw(sql, type0, userId, dateStart, dateEnd)
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}

// GetAnalysisLeixingDetail 按照用户、起止时间、客户获取详细类型数据
func GetAnalysisLeixingDetail(userId uint, project string, dateStart string, dateEnd string) (bool, *[]LeixingDetail, int64) {
	var d1 []LeixingDetail
	sql := `SELECT tycc leixing,id,namec titlename,dat workdate,esti,cons
			  FROM k_sum
			 WHERE proj = ?
			   AND acct = ?
			   AND dat BETWEEN ? AND ?`
	decodeproject, _ := url.QueryUnescape(project)
	res := zentao_common.ZENTAO_DB.Raw(sql, decodeproject, userId, dateStart, dateEnd)
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}

// GetAnalysisHeatMapLeixing 获取‘问题类型’年度热力图
func GetAnalysisHeatMapLeixing(userId uint, dateYear string) (bool, *[]LeixingHeatmap, int64) {
	var d1 []LeixingHeatmap
	sql := `SELECT tycc cloudname,sum(esti) esti,month(dat) dateyear,dense_rank() over (order by tycc) rk
      		  FROM k_sum
			 WHERE acct = ?
			   AND year(dat) = ?
			 GROUP BY tycc,month(dat)
			 ORDER BY rk,tycc,month(dat)`
	res := zentao_common.ZENTAO_DB.Raw(sql, userId, dateYear)
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}

// GetAnalysisHeatMapCustomer 获取‘客户’年度热力图
func GetAnalysisHeatMapCustomer(userId uint, dateYear string) (bool, *[]CustomerHeatmap, int64) {
	var d1 []CustomerHeatmap
	sql := `SELECT concat(projid,'.',proj) customername,sum(esti) esti,month(dat) dateyear,dense_rank() over (order by concat(projid,'.',proj)) rk
          	  FROM k_sum
          	 WHERE acct = ?
			   AND year(dat) = ?
			 GROUP BY concat(projid,'.',proj),month(dat)
			 ORDER BY rk,concat(projid,'.',proj),month(dat)`
	res := zentao_common.ZENTAO_DB.Raw(sql, userId, dateYear)
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}

// GetAnalysisLineCustomization 获取‘客制返工关系图’
func GetAnalysisLineCustomization(userId uint, dateYear string) (bool, *[]CustomizationLine, int64) {
	var d1 []CustomizationLine
	sql := `SELECT clid,concat(DATE_FORMAT(dat, '%m'),'-',DATE_FORMAT(dat, '%d')) datemonthday,DAYOFYEAR(dat) dayofyear,sum(esti) esti
      		  FROM k_sum
			 WHERE acct = ?
			   AND year(dat) = ?
			   AND clid IN ('T0343','T0344')
			 GROUP BY clid,concat(DATE_FORMAT(dat, '%m'),'-',DATE_FORMAT(dat, '%d'))
			 ORDER BY clid,concat(DATE_FORMAT(dat, '%m'),'-',DATE_FORMAT(dat, '%d'))`
	res := zentao_common.ZENTAO_DB.Raw(sql, userId, dateYear)
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}
