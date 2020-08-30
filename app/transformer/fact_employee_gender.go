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

// FactEmployeeGender ...
func FactEmployeeGender() {
	fmt.Println("populate fact employee gender ...")
	conn, err := repository.InitMysql()
	if err != nil {
		logger.Error(err.Error())
	}

	// get query latest

	queryLast := fmt.Sprintf(`SELECT * FROM last_updated 
		WHERE action = 'fact_employee_gender'
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

	query := fmt.Sprintf(`
	SELECT 
		COUNT(id) AS count,
		gender_id,
		gender,
		location_id,
		location,
		join_date,
		DAY(join_date) AS day,
		MONTH(join_date) AS month,
		YEAR(join_date) AS year
	FROM source_employee
	WHERE id > %d
	GROUP BY 
		DAY(join_date),
		MONTH(join_date),
		YEAR(join_date),
		location_id,
		gender_id`, lastID)

	res, err := conn.Read(query)
	if err != nil {
		logger.Error(err.Error())
	}

	var factEmployeeGender []model.FactGenderEmployee
	if resErr := json.Unmarshal([]byte(res), &factEmployeeGender); resErr != nil {
		logger.Error(resErr.Error())
	}

	length := len(factEmployeeGender)
	if length > 0 {
		i := 0
		for _, gender := range factEmployeeGender {
			count, err := strconv.Atoi(gender.Count)
			if err != nil {
				logger.Error("[Count to INT] failed convert: ", err.Error())
			}
			day, err := strconv.Atoi(gender.Day)
			if err != nil {
				logger.Error("[Day to INT] failed convert: ", err.Error())
			}
			month, err := strconv.Atoi(gender.Month)
			if err != nil {
				logger.Error("[Day to INT] failed convert: ", err.Error())
			}
			year, err := strconv.Atoi(gender.Year)
			if err != nil {
				logger.Error("[Day to INT] failed convert: ", err.Error())
			}

			// GET TIME DIMENSION
			queryTimeDimension := fmt.Sprintf(`
				SELECT * FROM ss_dimension_time
				WHERE day = %d
					AND month = %d
					AND YEAR = %d`, day, month, year)

			resDimensionTime, err := conn.Read(queryTimeDimension)
			if err != nil {
				logger.Error("[ERROR DIMENSION TIME] ", err.Error())
			}
			var timeDim []model.DimensionTime
			if resErr := json.Unmarshal([]byte(resDimensionTime), &timeDim); resErr != nil {
				logger.Error(resErr.Error())
			}
			dimTimeID, err := strconv.Atoi(timeDim[0].ID)
			if err != nil {
				logger.Error("[ERROR DIMENSION TIME]", err.Error())
			}

			// GET LOCATION DIMENSION
			queryLocationDimension := fmt.Sprintf(`
				SELECT * FROM ss_dimension_location
				WHERE name = '%s'`, gender.Location)

			resDimensionLocation, err := conn.Read(queryLocationDimension)
			if err != nil {
				logger.Error(err.Error())
			}
			var locationDim []model.DimensionLocation
			if resErr := json.Unmarshal([]byte(resDimensionLocation), &locationDim); resErr != nil {
				logger.Error(resErr.Error())
			}
			dimLocationID, err := strconv.Atoi(locationDim[0].ID)
			if err != nil {
				logger.Error("[ERROR DIMENSION TIME]", err.Error())
			}

			// GET GENDER DIMENSION
			queryGenderDimension := fmt.Sprintf(`
				SELECT * FROM ss_dimension_gender
				WHERE name = '%s'`, gender.Location)

			resDimensionGender, err := conn.Read(queryGenderDimension)
			if err != nil {
				logger.Error(err.Error())
			}
			var genderDim []model.DimensionGender
			if resErr := json.Unmarshal([]byte(resDimensionGender), &genderDim); resErr != nil {
				logger.Error(resErr.Error())
			}
			dimGenderID, err := strconv.Atoi(locationDim[0].ID)
			if err != nil {
				logger.Error("[ERROR DIMENSION TIME]", err.Error())
			}

			queryStore := fmt.Sprintf(`
				INSERT INTO fact_employee_gender 
				(time_id, location_id, gender_id, total_employee) 
				values (%d, %d, %d, %d)`,
				dimTimeID, dimLocationID, dimGenderID, count)

			result, err := conn.Store(queryStore)
			if err != nil {
				logger.Error(err.Error())
			}
			logger.Info(result)
			i++
		}
		fmt.Println("Inserted fact employee gender: ", i)

		queryGetLastID := fmt.Sprintf(`
		SELECT id
		FROM source_employee
		WHERE id > %d
		ORDER BY id DESC
		LIMIT 1`,
			lastID)

		resGetLast, err := conn.Read(queryGetLastID)
		if err != nil {
			logger.Error("[GET LAST GENDER]", err.Error())
		}

		var lastEmployee []model.SourceEmployee
		if resErr := json.Unmarshal([]byte(resGetLast), &lastEmployee); resErr != nil {
			logger.Error("[LAST GENDER UNMARSHAL]", resErr.Error())
		}

		last, err := strconv.Atoi(lastEmployee[0].ID)
		if err != nil {
			logger.Error("[LAST GENDER UNMARSHAL] failed convert: ", err.Error())
		}

		now := time.Now()

		queryUpdated := fmt.Sprintf(`
			INSERT INTO last_updated (action, last_id, created_at)
			VALUE('%s', %d, '%s')
		`, "fact_employee_gender", last, now.Format("2006-01-02 15:04:05"))

		_, errs := conn.Store((queryUpdated))
		if errs != nil {
			logger.Error("[INSERT INTO LAST UPDATED]", errs.Error())
		}

		fmt.Println("[FACT EMPLOYEE GENDER]: Success update data")
	} else {
		fmt.Println("[FACT EMPLOYEE GENDER]: There is no new data")
	}

	conn.Close()
}
