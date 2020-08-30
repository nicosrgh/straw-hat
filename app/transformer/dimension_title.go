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

// TitleDimension ...
func TitleDimension() {
	fmt.Println("populate dimension title ...")
	conn, err := repository.InitMysql()
	if err != nil {
		logger.Error(err.Error())
	}

	// QUERY GET LATEST
	queryLast := fmt.Sprintf(`SELECT * FROM last_updated 
		WHERE action = 'ss_dimension_title'
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

	query := fmt.Sprintf(`select * from title WHERE id > %d`, lastID)

	res, err := conn.Read(query)
	if err != nil {
		logger.Error(err.Error())
	}

	var titles []model.Title
	if resErr := json.Unmarshal([]byte(res), &titles); resErr != nil {
		logger.Error(resErr.Error())
	}

	length := len(titles)
	if length > 0 {
		i := 0
		for _, title := range titles {
			queryStore := fmt.Sprintf(`insert into ss_dimension_title (name) values ('%s')`, title.Name)
			_, err := conn.Store(queryStore)
			if err != nil {
				logger.Error(err.Error())
			}

			i++
		}
		fmt.Println("Inserted dimension title: ", i)

		queryGetLastID := fmt.Sprintf(`
		SELECT id
		FROM title
		WHERE id > %d
		ORDER BY id DESC
		LIMIT 1`,
			lastID)

		resGetLast, err := conn.Read(queryGetLastID)
		if err != nil {
			logger.Error("[GET LAST TITLE]", err.Error())
		}

		var lastTitle []model.Title
		if resErr := json.Unmarshal([]byte(resGetLast), &lastTitle); resErr != nil {
			logger.Error("[LAST TITLE UNMARSHAL]", resErr.Error())
		}

		last, err := strconv.Atoi(lastTitle[0].ID)
		if err != nil {
			logger.Error("[LAST TITLE UNMARSHAL] failed convert: ", err.Error())
		}

		now := time.Now()

		queryUpdated := fmt.Sprintf(`
			INSERT INTO last_updated (action, last_id, created_at)
			VALUE('%s', %d, '%s')
		`, "ss_dimension_title", last, now.Format("2006-01-02 15:04:05"))

		_, errs := conn.Store((queryUpdated))
		if errs != nil {
			logger.Error("[INSERT INTO LAST UPDATED]", errs.Error())
		}

		fmt.Println("[DIMENSION TITLE]: Success update data")
	} else {
		fmt.Println("[DIMENSION TITLE]: There is no new data")
	}

	conn.Close()
}
