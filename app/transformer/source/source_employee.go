package source

import (
	"encoding/json"
	"fmt"

	"github.com/nicosrgh/straw-hat/config/repository"
	"github.com/nicosrgh/straw-hat/lib/logger"
	"github.com/nicosrgh/straw-hat/model"
)

// EmployeeSource ...
func EmployeeSource(store *repository.Store) {
	fmt.Println("Employee source transform ...")

	queryLast := "select id, source_id from source_employee order by id desc limit 1"

	res, err := store.My.Read(queryLast)
	if err != nil {
		logger.Error(err.Error())
	}

	var srcEmployees []model.SourceEmployee
	if resErr := json.Unmarshal([]byte(res), &srcEmployees); resErr != nil {
		logger.Error(resErr.Error())
	}

	if len(srcEmployees) != 0 {
		query := fmt.Sprintf("select * from employee where id > %d", srcEmployees[0].SourceID)

		res, err := store.My.Read(query)
		if err != nil {
			logger.Error(err.Error())
		}

		var employees []model.Employee
		if resErr := json.Unmarshal([]byte(res), &employees); resErr != nil {
			logger.Error(resErr.Error())
		}

		if len(employees) != 0 {
			for _, employee := range employees {
				queryInsert := fmt.Sprintf("insert into source_employee (source_id, nip, fullname, status, gender, department, title, birthdate, join_date, resign_date) values (%d, %s, %s, %s, %s, %s, %s, %s, %s, %s)",
					employee.ID, employee.NIP, employee.Fullname, employee.Status, employee.Gender, employee.Department, employee.Title, employee.Birthdate, employee.JoinDate, employee.ResignDate)

				_, err := store.My.Read(queryInsert)
				if err != nil {
					logger.Error(err.Error())
				}
			}
		}
	} else {
		fmt.Println("No new employees ...")
	}

}
