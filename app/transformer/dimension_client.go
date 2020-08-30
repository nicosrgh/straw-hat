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

// ClientDimension ...
func ClientDimension() {
	fmt.Println("populate dimension client ...")
	conn, err := repository.InitMysql()
	if err != nil {
		logger.Error(err.Error())
	}

	// get query latest

	queryLast := fmt.Sprintf(`SELECT * FROM last_updated 
		WHERE action = 'ss_dimension_client'
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
		SELECT client.id, client.name, 
		CASE WHEN partner.name IS NOT NULL 
			THEN partner.name 
			ELSE 'INTERNAL' 
			END
		AS partner
		FROM client 
		LEFT JOIN partner ON partner.id = client.department_id
		WHERE client.id > %d`,
		lastID)

	res, err := conn.Read(query)
	if err != nil {
		logger.Error(err.Error())
	}

	var clients []model.Client
	if resErr := json.Unmarshal([]byte(res), &clients); resErr != nil {
		logger.Error(resErr.Error())
	}

	length := len(clients)
	if length > 0 {
		i := 0
		for _, client := range clients {
			queryStore := fmt.Sprintf(`
			INSERT INTO ss_dimension_client (name, partner) values ('%s', '%s')`,
				client.Name, client.Partner)
			_, err := conn.Store(queryStore)
			if err != nil {
				logger.Error(err.Error())
			}

			i++
		}
		fmt.Println("Inserted dimension client: ", i)

		queryGetLastID := fmt.Sprintf(`
		SELECT id
		FROM client
		WHERE id > %d
		ORDER BY id DESC
		LIMIT 1`,
			lastID)

		resGetLast, err := conn.Read(queryGetLastID)
		if err != nil {
			logger.Error("[GET LAST DEPARTMENT]", err.Error())
		}

		var lastTitle []model.Title
		if resErr := json.Unmarshal([]byte(resGetLast), &lastTitle); resErr != nil {
			logger.Error("[LAST DEPARTMENT UNMARSHAL]", resErr.Error())
		}

		last, err := strconv.Atoi(lastTitle[0].ID)
		if err != nil {
			logger.Error("[LAST DEPARTMENT UNMARSHAL] failed convert: ", err.Error())
		}

		now := time.Now()

		queryUpdated := fmt.Sprintf(`
			INSERT INTO last_updated (action, last_id, created_at)
			VALUE('%s', %d, '%s')
		`, "ss_dimension_client", last, now.Format("2006-01-02 15:04:05"))

		_, errs := conn.Store((queryUpdated))
		if errs != nil {
			logger.Error("[INSERT INTO LAST UPDATED]", errs.Error())
		}

		fmt.Println("[DIMENSION DEPARTMENT]: Success update data")
	} else {
		fmt.Println("[DIMENSION DEPARTMENT]: There is no new data")
	}

	conn.Close()
}
