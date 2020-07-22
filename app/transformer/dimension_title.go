package transformer

import (
	"encoding/json"
	"fmt"

	"github.com/nicosrgh/straw-hat/config/repository"
	"github.com/nicosrgh/straw-hat/lib/logger"
	"github.com/nicosrgh/straw-hat/model"
)

// TitleDimension ...
func TitleDimension() {
	fmt.Println("populate dimension title ...")

	query := fmt.Sprintf(`select DISTINCT title from employee`)

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
			queryExists := fmt.Sprintf(`select * from ss_dimension_title where name = '%s'`, employee.Title)

			res, err := conn.Read(queryExists)
			if err != nil {
				logger.Error(err.Error())
			}

			var dimGenders []model.DimensionGender
			if resErr := json.Unmarshal([]byte(res), &dimGenders); resErr != nil {
				logger.Error(resErr.Error())
			}

			if len(dimGenders) <= 0 {
				queryStore := fmt.Sprintf(`insert into ss_dimension_title (name) values ('%s')`, employee.Title)
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
