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

	query := "select count(id) count, title from employee group by title"

	res, err := conn.Read(query)
	if err != nil {
		logger.Error(err.Error())
	}

	var titles []model.Employee
	if resErr := json.Unmarshal([]byte(res), &titles); resErr != nil {
		logger.Error(resErr.Error())
	}

	length := len(titles)
	if length > 0 {
		for _, title := range titles {
			cnt, err := strconv.Atoi(title.Count)
			if err != nil {
				logger.Error(err.Error())
			}

			queryExists := fmt.Sprintf("select * from ss_title_dimension where name = '%s'", title.Title)
			fmt.Println("queryExists tidim", queryExists)

			res, err := conn.Read(queryExists)
			if err != nil {
				logger.Error(err.Error())
			}

			var dimTitle []model.DimensionTitle
			if resErr := json.Unmarshal([]byte(res), &dimTitle); resErr != nil {
				logger.Error(resErr.Error())
			}
			fmt.Println("dimTItle: ", dimTitle)

			if len(dimTitle) != 0 {
				titleID, err := strconv.Atoi(dimTitle[0].ID)
				if err != nil {
					logger.Error(err.Error())
				}

				queryExists := fmt.Sprintf("select * from ss_datamart_title where title_id = '%d'", titleID)
				fmt.Println("query exist datamart title: ", queryExists)
				resEx, err := conn.Read(queryExists)
				if err != nil {
					logger.Error(err.Error())
				}

				var datamartTitle []model.DatamartTitle
				if resErr := json.Unmarshal([]byte(resEx), &datamartTitle); resErr != nil {
					logger.Error(resErr.Error())
				}
				fmt.Println("datamartTitle: ", datamartTitle)

				queryStore := fmt.Sprintf("insert into ss_datamart_title (title_id, title_fact) values (%d, %d)",
					titleID, cnt)
				if len(datamartTitle) > 0 {
					queryStore = fmt.Sprintf("update ss_datamart_title set title_fact = %d where title_id = %d",
						cnt, titleID)
				}

				fmt.Println("queryStore: ", queryStore)

				_, errs := conn.Store(queryStore)
				if errs != nil {
					logger.Error(errs.Error())
				}
			}
		}
	}

	conn.Close()
}
