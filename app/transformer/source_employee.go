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

// EmployeeSource ...
func EmployeeSource() {
	fmt.Println("Employee source transform ...")

	conn, err := repository.InitMysql()
	if err != nil {
		logger.Error(err.Error())
	}

	// QUERY GET LATEST
	queryLast := fmt.Sprintf(`SELECT * FROM last_updated 
		WHERE action = 'source_employee'
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
		employee.id,
		employee.fullname,
		employee.nip,
		employee.status_id,
		status.name AS status,
		employee.department_id,
		department.name AS department,
		employee.location_id,
		location.name AS location,
		employee.title_id,
		title.name AS title,
		employee.gender_id,
		gender.name AS gender,
		employee.birthdate,
		employee.join_date
	FROM employee
	JOIN status ON status.id = employee.status_id
	JOIN gender ON gender.id = employee.gender_id
	JOIN department ON department.id = employee.department_id
	JOIN title ON title.id = employee.title_id
	JOIN location ON location.id = employee.location_id
	WHERE employee.id > %d;`, lastID)

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

			titleID, err := strconv.Atoi(employee.TitleID)
			if err != nil {
				logger.Error(err.Error())
			}
			genderID, err := strconv.Atoi(employee.GenderID)
			if err != nil {
				logger.Error(err.Error())
			}
			locationID, err := strconv.Atoi(employee.LocationID)
			if err != nil {
				logger.Error(err.Error())
			}
			statusID, err := strconv.Atoi(employee.StatusID)
			if err != nil {
				logger.Error(err.Error())
			}
			departmentID, err := strconv.Atoi(employee.DepartmentID)
			if err != nil {
				logger.Error(err.Error())
			}

			queryInsert := fmt.Sprintf(
				`INSERT INTO source_employee (
					source_id,
					nip,
					fullname, 
					status,
					status_id,
					gender,
					gender_id, 
					department,
					department_id, 
					location,
					location_id,
					title,
					title_id,
					birthdate, 
					join_date
				) 
				values (%d, '%s', '%s', '%s', %d, '%s', %d, '%s', %d, '%s', %d, '%s', %d, '%s', '%s')`,
				ID,
				employee.NIP,
				employee.Fullname,
				employee.Status,
				statusID,
				employee.Gender,
				genderID,
				employee.Department,
				departmentID,
				employee.Location,
				locationID,
				employee.Title,
				titleID,
				employee.Birthdate,
				employee.JoinDate,
			)

			_, errs := conn.Store(queryInsert)
			if errs != nil {
				logger.Error(errs.Error())
			}
		}

		queryGetLastID := fmt.Sprintf(`
		SELECT id
		FROM employee
		WHERE id > %d
		ORDER BY id DESC
		LIMIT 1`,
			lastID)

		resGetLast, err := conn.Read(queryGetLastID)
		if err != nil {
			logger.Error("[GET LAST SOURCE EMPLOYEE]", err.Error())
		}

		var lastTitle []model.Title
		if resErr := json.Unmarshal([]byte(resGetLast), &lastTitle); resErr != nil {
			logger.Error("[LAST SOURCE EMPLOYEE UNMARSHAL]", resErr.Error())
		}

		last, err := strconv.Atoi(lastTitle[0].ID)
		if err != nil {
			logger.Error("[LAST SOURCE EMPLOYEE UNMARSHAL] failed convert: ", err.Error())
		}

		now := time.Now()

		queryUpdated := fmt.Sprintf(`
			INSERT INTO last_updated (action, last_id, created_at)
			VALUE('%s', %d, '%s')
		`, "source_employee", last, now.Format("2006-01-02 15:04:05"))

		_, errs := conn.Store((queryUpdated))
		if errs != nil {
			logger.Error("[INSERT INTO LAST UPDATED]", errs.Error())
		}

		fmt.Println("[SOURCE EMPOYEE]: Success update data")
	} else {
		fmt.Println("[SOURCE EMPLOYEE]: There is no new data")
	}

	conn.Close()

}
