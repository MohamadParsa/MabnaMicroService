package trade

import (
	"database/sql"

	"./postgres"
)

// GetLastTradeOfAllInstrumentsRows returns last trade of all instruments
func GetLastTradeOfAllInstrumentsRows() (*sql.Rows, error) {
	pg, err := postgres.InitializeDatabase("postgresql://postgres:Abc1234@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}
	rows, err := pg.ExecQuery(`SELECT I.NAME INSTUMENTNAME,
		T.ID TRADEID,
		T.INSTRUMENTID,
		T.DATEEN,
		T.OPEN,
		T.HIGH,
		T.LOW,
		T.CLOSE
	FROM INSTRUMENT I
	left JOIN TRADE T ON I.ID = T.INSTRUMENTID
	WHERE NOT EXISTS
			(SELECT 1
				FROM TRADE TR
				WHERE TR.INSTRUMENTID = T.INSTRUMENTID
					AND TR.DATEEN > T.DATEEN) ;
	`)
	return rows, err
}
