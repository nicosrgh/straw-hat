package transformer

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/nicosrgh/straw-hat/config/repository"
	"github.com/nicosrgh/straw-hat/lib/logger"
	"github.com/nicosrgh/straw-hat/model"
)

// EmployeeTitleDatamart ...
func EmployeeTitleDatamart() {
	fmt.Println("Transform employee title datamart ...")

	conn, err := repository.InitMysql()
	if err != nil {
		logger.Error(err.Error())
	}

	query := fmt.Sprintf(`
	select 
		count(id) count, 
		title, 
		day(join_date) AS day,
		MONTH(join_date) AS month,
		year(join_date) AS year
	from employee 
	group by 
		title, 
		day(join_date),
		MONTH(join_date),
		year(join_date)`)

	res, err := conn.Read(query)
	if err != nil {
		logger.Error(err.Error())
	}

	var titles []model.SourceEmployee
	if resErr := json.Unmarshal([]byte(res), &titles); resErr != nil {
		logger.Error(resErr.Error())
	}

	length := len(titles)
	if length > 0 {
		for _, title := range titles {
			cnt, err := strconv.Atoi(title.TotalEmployee)
			if err != nil {
				logger.Error(err.Error())
			}

			day, err := strconv.Atoi(title.Day)
			if err != nil {
				logger.Error(err.Error())
			}

			month, err := strconv.Atoi(title.Month)
			if err != nil {
				logger.Error(err.Error())
			}

			year, err := strconv.Atoi(title.Year)
			if err != nil {
				logger.Error(err.Error())
			}

			queryExists := fmt.Sprintf(`select * from ss_dimension_title where name = '%s'`, title.Title)
			res, err := conn.Read(queryExists)
			if err != nil {
				logger.Error(err.Error())
			}

			var dimTitle []model.DimensionTitle
			if resErr := json.Unmarshal([]byte(res), &dimTitle); resErr != nil {
				logger.Error(resErr.Error())
			}

			queryDateExists := fmt.Sprintf(`
				select * from ss_dimension_time 
				where day = %d
					and month = %d
					and year = %d`,
				day,
				month,
				year,
			)
			fmt.Println("queryDateExists: ", queryDateExists)
			resTime, err := conn.Read(queryDateExists)
			if err != nil {
				logger.Error(err.Error())
			}

			var dimTime []model.DimensionTime
			if resErr := json.Unmarshal([]byte(resTime), &dimTime); resErr != nil {
				logger.Error(resErr.Error())
			}

			fmt.Println("dimTime: ", dimTime)

			if len(dimTitle) != 0 {
				titleID, err := strconv.Atoi(dimTitle[0].ID)
				if err != nil {
					logger.Error(err.Error())
				}

				timeID, err := strconv.Atoi(dimTime[0].ID)
				if err != nil {
					logger.Error(err.Error())
				}

				queryExists := fmt.Sprintf(`
					select * from ss_datamart_title where title_id = '%d'`,
					titleID,
				)
				resEx, err := conn.Read(queryExists)
				if err != nil {
					logger.Error(err.Error())
				}

				var datamartTitle []model.DatamartTitle
				if resErr := json.Unmarshal([]byte(resEx), &datamartTitle); resErr != nil {
					logger.Error(resErr.Error())
				}

				queryStore := fmt.Sprintf(`
					INSERT INTO ss_datamart_title 
					(
						title_id,
						time_id, 
						title_fact
					) 
					VALUES (%d, %d, %d)`,
					titleID, timeID, cnt)

				if len(datamartTitle) > 0 {
					queryStore = fmt.Sprintf(`
						UPDATE ss_datamart_title 
						SET title_fact = %d 
						WHERE title_id = %d`,
						cnt,
						titleID,
					)
				}

				_, errs := conn.Store(queryStore)
				if errs != nil {
					logger.Error(errs.Error())
				}
			}
		}
	}

	conn.Close()
}
