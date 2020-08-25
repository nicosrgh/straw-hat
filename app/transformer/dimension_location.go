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

// LocationDimension ...
func LocationDimension() {
	fmt.Println("populate dimension location ...")
	conn, err := repository.InitMysql()
	if err != nil {
		logger.Error(err.Error())
	}

	// get query latest

	queryLast := fmt.Sprintf(`SELECT * FROM last_updated 
		WHERE action = 'ss_dimension_location'
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

	query := fmt.Sprintf(`select * from location WHERE id > %d`, lastID)

	res, err := conn.Read(query)
	if err != nil {
		logger.Error(err.Error())
	}

	var locations []model.Location
	if resErr := json.Unmarshal([]byte(res), &locations); resErr != nil {
		logger.Error(resErr.Error())
	}

	length := len(locations)
	if length > 0 {
		i := 0
		for _, location := range locations {
			queryStore := fmt.Sprintf(`insert into ss_dimension_location (name) values ('%s')`, location.Name)
			_, err := conn.Store(queryStore)
			if err != nil {
				logger.Error(err.Error())
			}

			i++
		}
		fmt.Println("Inserted dimension location: ", i)

		queryGetLastID := fmt.Sprintf(`
		SELECT id
		FROM location
		WHERE id > %d
		ORDER BY id DESC
		LIMIT 1`,
			lastID)

		resGetLast, err := conn.Read(queryGetLastID)
		if err != nil {
			logger.Error("[GET LAST LOCATION]", err.Error())
		}

		var lastLocation []model.Location
		if resErr := json.Unmarshal([]byte(resGetLast), &lastLocation); resErr != nil {
			logger.Error("[LAST LOCATION UNMARSHAL]", resErr.Error())
		}

		last, err := strconv.Atoi(lastLocation[0].ID)
		if err != nil {
			logger.Error("[LAST LOCATION UNMARSHAL] failed convert: ", err.Error())
		}

		now := time.Now()

		queryUpdated := fmt.Sprintf(`
			INSERT INTO last_updated (action, last_id, created_at)
			VALUE('%s', %d, '%s')
		`, "ss_dimension_location", last, now.Format("2006-01-02 15:04:05"))

		_, errs := conn.Store((queryUpdated))
		if errs != nil {
			logger.Error("[INSERT INTO LAST UPDATED]", errs.Error())
		}

		fmt.Println("[DIMENSION LOCATION]: Success update data")
	} else {
		fmt.Println("[DIMENSION LOCATION]: There is no new data")
	}

	conn.Close()
}
