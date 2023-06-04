package zentao_models

import (
	"wzDataCenter/app/zentao/zentao_common"
)

// GetAnalysisLeixing 按照用户、起止时间获取‘类型’数据
func GetAnalysisLeixing(userId uint, dateStart string, dateEnd string) (bool, *[]Leixing, int64) {
	var d1 []Leixing
	sql := `SELECT cloudname,sum(tt.esti) esti,sum(tt.cons) cons
			  FROM (SELECT t3.id id,t1.name namec,t1.project proj,t3.account acct,t3.date dat,
			               t1.estimate esti,t3.consumed cons,t1.type tycc
			          FROM zt_effort t3
			          LEFT JOIN zt_task t1 ON t1.id = t3.objectID
			         WHERE t1.deleted <> 2
			           AND t3.deleted <> 2
			           AND t1.status IN ( 'closed', 'done' )
			           AND t1.type = 'T0343'
			           AND NOT EXISTS (SELECT 1 FROM zt_task t2 WHERE t2.parent = t1.id AND t2.deleted <> 2 )
			           AND EXISTS (SELECT 1 FROM(SELECT id,ROW_NUMBER() over ( PARTITION BY objectID ORDER BY date,id) rk
			                                       FROM zt_effort
			                                      WHERE deleted <> 2 ) t6
			                               WHERE rk = 1
			                                 AND t6.id = t3.id)
			         UNION ALL
			        SELECT t3.id id,t1.name namec,t1.project proj,t3.account acct,t3.date dat,
			               0 esti,t3.consumed cons,t1.type tycc
			          FROM zt_effort t3
			          LEFT JOIN zt_task t1 ON t1.id = t3.objectID
			         WHERE t1.deleted <> 2
			           AND t3.deleted <> 2
			           AND t1.status IN ( 'closed', 'done' )
			           AND t1.type = 'T0343'
			           AND NOT EXISTS (SELECT 1 FROM zt_task t2 WHERE t2.parent = t1.id AND t2.deleted <> 2 )
			           AND EXISTS (SELECT 1 FROM(SELECT id,ROW_NUMBER() over ( PARTITION BY objectID ORDER BY date,id) rk
			                                       FROM zt_effort
			                                      WHERE deleted <> 2 ) t6
			                               WHERE rk > 1
			                                 AND t6.id = t3.id)
			         UNION ALL
			        SELECT t3.id id,t1.name namec,t1.project proj,t3.account acct,t3.date dat,
			               t3.consumed esti,t3.consumed cons,t1.type tycc
			          FROM zt_effort t3
			          LEFT JOIN zt_task t1 ON t1.id = t3.objectID
			         WHERE t1.deleted <> 2
			           AND t3.deleted <> 2
			           AND t1.status IN ( 'closed', 'done' )
			           AND t1.type <> 'T0343'
			           AND NOT EXISTS (SELECT 1 FROM zt_task t2 WHERE t2.parent = t1.id AND t2.deleted <> 2 )
			       ) tt
			  left join kt_cloud on cloudid=tt.tycc
			 WHERE tt.dat BETWEEN ? AND ?
			   and tt.acct = ?
			 GROUP BY tt.acct,tt.tycc
			 ORDER BY tt.acct,tt.tycc`
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
	sql := `SELECT t0.name customername,sum(tt.esti) esti,sum(tt.cons) cons
          	  FROM (SELECT t3.id id,t1.name namec,t1.project proj,t3.account acct,t3.date dat,
          	               t1.estimate esti,t3.consumed cons,t1.type tycc
          	          FROM zt_effort t3
          	          LEFT JOIN zt_task t1 ON t1.id = t3.objectID
          	         WHERE t1.deleted <> 2
          	           AND t3.deleted <> 2
          	           AND t1.status IN ( 'closed', 'done' )
          	           AND t1.type = 'T0343'
          	           AND NOT EXISTS (SELECT 1 FROM zt_task t2 WHERE t2.parent = t1.id AND t2.deleted <> 2 )
          	           AND EXISTS (SELECT 1 FROM (SELECT id,ROW_NUMBER() over ( PARTITION BY objectID ORDER BY date,id) rk
          	                                        FROM zt_effort
          	                                       WHERE deleted <> 2 ) t6
          	                               WHERE rk = 1
          	                                 AND t6.id = t3.id)
          	         UNION ALL
          	        SELECT t3.id id,t1.name namec,t1.project proj,t3.account acct,t3.date dat,
          	               0 esti,t3.consumed cons,t1.type tycc
          	          FROM zt_effort t3
          	          LEFT JOIN zt_task t1 ON t1.id = t3.objectID
          	         WHERE t1.deleted <> 2
          	           AND t3.deleted <> 2
          	           AND t1.status IN ( 'closed', 'done' )
          	           AND t1.type = 'T0343'
          	           AND NOT EXISTS (SELECT 1 FROM zt_task t2 WHERE t2.parent = t1.id AND t2.deleted <> 2 )
          	           AND EXISTS (SELECT 1 FROM (SELECT id,ROW_NUMBER() over ( PARTITION BY objectID ORDER BY date,id) rk
          	                                        FROM zt_effort
          	                                       WHERE deleted <> 2 ) t6
          	                               WHERE rk > 1
          	                                 AND t6.id = t3.id)
          	         UNION ALL
          	        SELECT t3.id id,t1.name namec,t1.project proj,t3.account acct,t3.date dat,
          	               t3.consumed esti,t3.consumed cons,t1.type tycc
          	          FROM zt_effort t3
          	          LEFT JOIN zt_task t1 ON t1.id = t3.objectID
          	         WHERE t1.deleted <> 2
          	           AND t3.deleted <> 2
          	           AND t1.status IN ( 'closed', 'done' )
          	           AND t1.type <> 'T0343'
          	           AND NOT EXISTS (SELECT 1 FROM zt_task t2 WHERE t2.parent = t1.id AND t2.deleted <> 2 )) tt
          	  LEFT JOIN zt_project t0 ON t0.id = tt.proj
          	 WHERE tt.dat BETWEEN ? AND ?
          	   and tt.acct = ?
          	 GROUP BY tt.acct,tt.proj
          	 ORDER BY tt.acct,tt.proj`
	res := zentao_common.ZENTAO_DB.Raw(sql, dateStart, dateEnd, userId)
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}
