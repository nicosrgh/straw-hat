package transformer

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/nicosrgh/straw-hat/config/repository"
	"github.com/nicosrgh/straw-hat/lib/logger"
	"github.com/nicosrgh/straw-hat/model"
)

// FactEmployeeLocation ...
func FactEmployeeLocation() {
	fmt.Println("Transform fact employee location ...")

	conn, err := repository.InitMysql()
	if err != nil {
		logger.Error(err.Error())
	}

	// get last id
	queryLast := fmt.Sprintf(`SELECT * FROM last_updated 
		WHERE action = 'ss_fact_employee_location'
		ORDER BY created_at DESC
		LIMIT 1`)
	resLast, err := conn.Read(queryLast)
	if err != nil {
		logger.Error(err.Error())
	}
	var lastID = 0
	var lastUpdate []model.LastUpdate
	if resErr := json.Unmarshal([]byte(resLast), &lastUpdate); resErr != nil {
		logger.Error(resErr.Error())
	}

	if len(lastUpdate) > 0 {
		rs, err := strconv.Atoi(lastUpdate[0].LastID)
		if err != nil {
			logger.Error(err.Error())
		}
		lastID = rs
	}

	// get employee
	query := fmt.Sprintf(`
		SELECT location, 
			DAY(join_date) AS day, 
			MONTH(join_date) AS month, 
			YEAR(join_date) AS year, 
			count(location) AS total_employee
		FROM source_employee
		WHERE id > %d
		GROUP BY location, 
			DAY(join_date), 
			MONTH(join_date), 
			YEAR(join_date)
	`, lastID)

	res, err := conn.Read(query)
	if err != nil {
		logger.Error("[GET EMPLOYEE]", err.Error())
	}

	var employees []model.SourceEmployee
	if resErr := json.Unmarshal([]byte(res), &employees); resErr != nil {
		logger.Error("[GET EMPLOYEE UNMARSHAL]", resErr.Error())
	}

	// insert to fact table
	if len(employees) > 0 {
		for _, employee := range employees {
			// get time dimension
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

			queryDimensionTime := fmt.Sprintf(`
				SELECT * 
				FROM ss_dimension_time
				WHERE day = %d
					AND month = %d
					AND year = %d`,
				day, month, year)

			resDimTime, err := conn.Read(queryDimensionTime)
			if err != nil {
				logger.Error("[DIMENSION TIME]: ", err.Error())
			}

			var dimTime []model.DimensionTime
			if resErr := json.Unmarshal([]byte(resDimTime), &dimTime); resErr != nil {
				logger.Error("[DIMENSION TIME UNMARSHAL]", resErr.Error())
			}

			// get dimension location
			queryDimensionLocation := fmt.Sprintf(`
				SELECT * FROM ss_dimension_location
				WHERE name = '%s'
			`, employee.Location)

			resDimLoc, err := conn.Read(queryDimensionLocation)
			if err != nil {
				logger.Error("[GET DIMENSION LOCATION]", err.Error())
			}

			var dimLocation []model.DimensionLocation
			if resErr := json.Unmarshal([]byte(resDimLoc), &dimLocation); resErr != nil {
				logger.Error("[DIMENSION LOCATION UNMARSHAL]", resErr.Error())
			}

			// insert to fact table

			timeID, err := strconv.Atoi(dimTime[0].ID)
			if err != nil {
				logger.Error("[TIME ID] failed convert: ", err.Error())
			}

			locationID, err := strconv.Atoi(dimLocation[0].ID)
			if err != nil {
				logger.Error("[LOCATION ID] failed convert: ", err.Error())
			}

			totalEmployee, err := strconv.Atoi(employee.TotalEmployee)
			if err != nil {
				logger.Error("[TOTAL EMPLOYEE] failed convert: ", err.Error())
			}

			queryFactTable := fmt.Sprintf(`
				INSERT INTO ss_fact_employee_location 
				(time_id, location_id, total_employee)
				VALUES (%d, %d, %d)
			`, timeID, locationID, totalEmployee)

			_, errs := conn.Store((queryFactTable))
			if errs != nil {
				logger.Error("[INSERT INTO FACT TABLE]", errs.Error())
			}
		}

		queryGetLastID := fmt.Sprintf(`
		SELECT id
		FROM source_employee
		WHERE id > %d
		ORDER BY id DESC
		LIMIT 1`,
			lastID)

		resGetLast, err := conn.Read(queryGetLastID)
		if err != nil {
			logger.Error("[GET LAST EMPLOYEE]", err.Error())
		}

		var lastEmployee []model.SourceEmployee
		if resErr := json.Unmarshal([]byte(resGetLast), &lastEmployee); resErr != nil {
			logger.Error("[LAST EMPLOYEE UNMARSHAL]", resErr.Error())
		}

		last, err := strconv.Atoi(lastEmployee[0].ID)
		if err != nil {
			logger.Error("[LAST EMPLOYEE UNMARSHAL] failed convert: ", err.Error())
		}

		now := time.Now()

		queryUpdated := fmt.Sprintf(`
			INSERT INTO last_updated (action, last_id, created_at)
			VALUE('%s', %d, '%s')
		`, "ss_fact_employee_location", last, now.Format("2006-01-02 15:04:05"))

		_, errs := conn.Store((queryUpdated))
		if errs != nil {
			logger.Error("[INSERT INTO LAST UPDATED]", errs.Error())
		}

		fmt.Println("[FACT TABLE EMPLOYEE LOCATION]: Success update data")
	} else {
		fmt.Println("[FACT TABLE EMPLOYEE LOCATION]: There is no new data")
	}

	conn.Close()
}
