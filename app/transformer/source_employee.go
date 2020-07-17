package transformer

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/nicosrgh/straw-hat/config/repository"
	"github.com/nicosrgh/straw-hat/lib/logger"
	"github.com/nicosrgh/straw-hat/model"
)

// EmployeeSource ...
func EmployeeSource() {
	fmt.Println("Employee source transform ...")

	conn, err := repository.InitMysql()
	if err != nil {
		logger.Error(err.Error())
	}

	queryLast := fmt.Sprintf(`select id, source_id from source_employee order by id desc limit 1`)

	res, err := conn.Read(queryLast)
	if err != nil {
		logger.Error(err.Error())
	}

	var srcEmployees []model.SourceEmployee
	if resErr := json.Unmarshal([]byte(res), &srcEmployees); resErr != nil {
		logger.Error(resErr.Error())
	}

	ID := 0

	if len(srcEmployees) != 0 {
		rslt, err := strconv.Atoi(srcEmployees[0].SourceID)
		if err != nil {
			logger.Error(err.Error())
		}
		ID = rslt
	}

	query := fmt.Sprintf("select * from employee where id > %d", ID)

	resEmp, err := conn.Read(query)
	if err != nil {
		logger.Error(err.Error())
	}

	var employees []model.Employee
	if resErr := json.Unmarshal([]byte(resEmp), &employees); resErr != nil {
		logger.Error(resErr.Error())
	}

	if len(employees) != 0 {
		for _, employee := range employees {
			ID, err := strconv.Atoi(employee.ID)
			if err != nil {
				logger.Error(err.Error())
			}

			queryInsert := fmt.Sprintf(
				`insert into source_employee (
				source_id, 
				nip, 
				fullname, 
				status, 
				gender, 
				department, 
				location, 
				title, 
				birthdate, 
				join_date
			) 
			values (%d, '%s', '%s', '%s', '%s', '%s', '%s','%s', '%s', '%s')`,
				ID,
				employee.NIP,
				employee.Fullname,
				employee.Status,
				employee.Gender,
				employee.Department,
				employee.Location,
				employee.Title,
				employee.Birthdate,
				employee.JoinDate,
			)
			fmt.Println("queryInsert: ", queryInsert)

			_, errs := conn.Store(queryInsert)
			if errs != nil {
				logger.Error(errs.Error())
			}
		}
	}

	conn.Close()

}
