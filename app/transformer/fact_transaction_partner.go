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

// FactTransactionPartner ...
func FactTransactionPartner() {
	fmt.Println("populate fact transaction partner ...")
	conn, err := repository.InitMysql()
	if err != nil {
		logger.Error(err.Error())
	}

	// get query latest
	queryLast := fmt.Sprintf(`SELECT * FROM last_updated 
		WHERE action = 'fact_transaction_partner'
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
		department AS partner,
		created_at,
		DAY(created_at) AS day,
		MONTH(created_at) AS month,
		YEAR(created_at) AS year
	FROM source_transaction
	WHERE id > %d
	GROUP BY 
		DAY(created_at),
		MONTH(created_at),
		YEAR(created_at),
		department`, lastID)

	res, err := conn.Read(query)
	if err != nil {
		logger.Error(err.Error())
	}

	var factTransactionPartner []model.FactTransactionPartner
	if resErr := json.Unmarshal([]byte(res), &factTransactionPartner); resErr != nil {
		logger.Error(resErr.Error())
	}

	length := len(factTransactionPartner)
	if length > 0 {
		i := 0
		for _, partner := range factTransactionPartner {
			count, err := strconv.Atoi(partner.Count)
			if err != nil {
				logger.Error("[Count to INT] failed convert: ", err.Error())
			}
			day, err := strconv.Atoi(partner.Day)
			if err != nil {
				logger.Error("[Day to INT] failed convert: ", err.Error())
			}
			month, err := strconv.Atoi(partner.Month)
			if err != nil {
				logger.Error("[Day to INT] failed convert: ", err.Error())
			}
			year, err := strconv.Atoi(partner.Year)
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

			// GET CLIENT DIMENSION
			queryPartnerDimension := fmt.Sprintf(`
				SELECT * FROM ss_dimension_partner
				WHERE name = '%s'`, partner.Partner)

			resDimensionPartner, err := conn.Read(queryPartnerDimension)
			if err != nil {
				logger.Error(err.Error())
			}
			var partnerDim []model.DimensionPartner
			if resErr := json.Unmarshal([]byte(resDimensionPartner), &partnerDim); resErr != nil {
				logger.Error(resErr.Error())
			}
			dimPartnerID, err := strconv.Atoi(partnerDim[0].ID)
			if err != nil {
				fmt.Println("1")
				logger.Error("[ERROR DIMENSION TIME]", err.Error())
			}

			queryStore := fmt.Sprintf(`
				INSERT INTO fact_transaction_partner 
				(time_id, partner_id, total_employee) 
				values (%d, %d, %d)`,
				dimTimeID, dimPartnerID, count)

			result, err := conn.Store(queryStore)
			if err != nil {
				logger.Error(err.Error())
			}
			logger.Info(result)
			i++
		}
		fmt.Println("Inserted fact employee partner: ", i)

		queryGetLastID := fmt.Sprintf(`
		SELECT id
		FROM source_transaction
		WHERE id > %d
		ORDER BY id DESC
		LIMIT 1`,
			lastID)

		resGetLast, err := conn.Read(queryGetLastID)
		if err != nil {
			logger.Error("[GET LAST LOCATION]", err.Error())
		}

		var lastEmployee []model.SourceEmployee
		if resErr := json.Unmarshal([]byte(resGetLast), &lastEmployee); resErr != nil {
			logger.Error("[LAST LOCATION UNMARSHAL]", resErr.Error())
		}

		last, err := strconv.Atoi(lastEmployee[0].ID)
		if err != nil {
			logger.Error("[LAST LOCATION UNMARSHAL] failed convert: ", err.Error())
		}

		now := time.Now()

		queryUpdated := fmt.Sprintf(`
			INSERT INTO last_updated (action, last_id, created_at)
			VALUE('%s', %d, '%s')
		`, "fact_transaction_partner", last, now.Format("2006-01-02 15:04:05"))

		_, errs := conn.Store((queryUpdated))
		if errs != nil {
			logger.Error("[INSERT INTO LAST UPDATED]", errs.Error())
		}

		fmt.Println("[FACT TRANSACTION PARTNER]: Success update data")
	} else {
		fmt.Println("[FACT TRANSACTION PARTNER]: There is no new data")
	}

	conn.Close()
}
