package transformer

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/nicosrgh/straw-hat/config/repository"
	"github.com/nicosrgh/straw-hat/lib/logger"
	"github.com/nicosrgh/straw-hat/model"
)

// EmployeeTotalDatamart ...
func EmployeeTotalDatamart() {
	fmt.Println("Transform employee total monthly datamart ...")

	conn, err := repository.InitMysql()
	if err != nil {
		logger.Error(err.Error())
	}

	query := fmt.Sprintf(`
	SELECT 
		SUM(1) AS total_employee,
		SUM(if(gender = 'Male', 1, 0)) AS male_employee,
		SUM(if(gender = 'Female', 1, 0)) AS female_employee,
		SUM(if(status = 'Full-time', 1, 0)) AS full_time_employee,
		SUM(if(status = 'probation', 1, 0)) AS probation_employee,
		DAY(join_date) as day,
		MONTH(join_date) AS month,
		YEAR(join_date) AS year
	FROM source_employee 
	GROUP BY 
		DAY(join_date),
		MONTH(join_date),
		YEAR(join_date)`)

	res, err := conn.Read(query)
	if err != nil {
		logger.Error(err.Error())
	}

	var employees []model.SourceEmployee
	if resErr := json.Unmarshal([]byte(res), &employees); resErr != nil {
		logger.Error(resErr.Error())
	}

	length := len(employees)
	if length > 0 {
		for _, employee := range employees {
			cnt, err := strconv.Atoi(employee.TotalEmployee)
			if err != nil {
				logger.Error(err.Error())
			}

			male, err := strconv.Atoi(employee.MaleEmployee)
			if err != nil {
				logger.Error(err.Error())
			}

			female, err := strconv.Atoi(employee.FemaleEmployee)
			if err != nil {
				logger.Error(err.Error())
			}

			fullTime, err := strconv.Atoi(employee.FullTimeEmployee)
			if err != nil {
				logger.Error(err.Error())
			}

			probation, err := strconv.Atoi(employee.ProbationEmployee)
			if err != nil {
				logger.Error(err.Error())
			}

			day, err := strconv.Atoi(employee.Day)
			if err != nil {
				logger.Error(err.Error())
			}

			month, err := strconv.Atoi(employee.Month)
			if err != nil {
				logger.Error(err.Error())
			}

			year, err := strconv.Atoi(employee.Year)
			if err != nil {
				logger.Error(err.Error())
			}

			queryDateExists := fmt.Sprintf(`
				SELECT * FROM ss_dimension_time 
				WHERE
					DAY = %d
					AND MONTH = %d
					AND year = %d`,
				day,
				month,
				year,
			)
			resTime, err := conn.Read(queryDateExists)
			if err != nil {
				logger.Error(err.Error())
			}

			var dimTime []model.DimensionTime
			if resErr := json.Unmarshal([]byte(resTime), &dimTime); resErr != nil {
				logger.Error(resErr.Error())
			}

			timeID, err := strconv.Atoi(dimTime[0].ID)
			if err != nil {
				logger.Error(err.Error())
			}

			queryExists := fmt.Sprintf(`
					select * from ss_datamart_employee_total where time_id = %d`,
				timeID,
			)
			resEx, err := conn.Read(queryExists)
			if err != nil {
				logger.Error(err.Error())
			}

			var datamartEmp []model.DatamartEmployeeTotal
			if resErr := json.Unmarshal([]byte(resEx), &datamartEmp); resErr != nil {
				logger.Error(resErr.Error())
			}

			queryStore := fmt.Sprintf(`
					INSERT INTO ss_datamart_employee_total 
					(
						time_id, 
						total_employee,
						male_employee,
						female_employee,
						probation_employee,
						full_time_employee
					) 
					VALUES (%d, %d, %d, %d, %d, %d)`,
				timeID, cnt, male, female, probation, fullTime)

			if len(datamartEmp) > 0 {
				queryStore = fmt.Sprintf(`
						UPDATE ss_datamart_employee_total 
						SET total_employee = %d,
							male_employee = %d,
							female_employee = %d, 
							full_time_employee = %d,
							probation_employee = %d,
						WHERE time_id = %d`,
					cnt,
					male,
					female,
					fullTime,
					probation,
					timeID,
				)
			}

			_, errs := conn.Store(queryStore)
			if errs != nil {
				logger.Error(errs.Error())
			}
		}
	}

	conn.Close()
}
