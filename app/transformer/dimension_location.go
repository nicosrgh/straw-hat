package transformer

import (
	"encoding/json"
	"fmt"

	"github.com/nicosrgh/straw-hat/config/repository"
	"github.com/nicosrgh/straw-hat/lib/logger"
	"github.com/nicosrgh/straw-hat/model"
)

// LocationDimension ...
func LocationDimension() {
	fmt.Println("populate dimension title ...")

	query := fmt.Sprintf(`select DISTINCT location from employee`)

	conn, err := repository.InitMysql()
	if err != nil {
		logger.Error(err.Error())
	}

	res, err := conn.Read(query)
	if err != nil {
		logger.Error(err.Error())
	}

	var employees []model.Employee
	if resErr := json.Unmarshal([]byte(res), &employees); resErr != nil {
		logger.Error(resErr.Error())
	}

	length := len(employees)
	if length > 0 {
		i := 0
		for _, employee := range employees {
			queryExists := fmt.Sprintf(`select * from ss_dimension_location where name = '%s'`, employee.Location)

			res, err := conn.Read(queryExists)
			if err != nil {
				logger.Error(err.Error())
			}

			var dimLocation []model.DimensionLocation
			if resErr := json.Unmarshal([]byte(res), &dimLocation); resErr != nil {
				logger.Error(resErr.Error())
			}

			if len(dimLocation) <= 0 {
				queryStore := fmt.Sprintf(`insert into ss_dimension_location (name) values ('%s')`, employee.Location)
				_, err := conn.Store(queryStore)
				if err != nil {
					logger.Error(err.Error())
				}

				i++
			}
		}
		fmt.Println("Inserted dimension title: ", i)
	}

	conn.Close()
}
