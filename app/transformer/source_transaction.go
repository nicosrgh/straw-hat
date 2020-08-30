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

// TransactionSource ...
func TransactionSource() {
	fmt.Println("Transaction source transform ...")

	conn, err := repository.InitMysql()
	if err != nil {
		logger.Error(err.Error())
	}

	// QUERY GET LATEST
	queryLast := fmt.Sprintf(`SELECT * FROM last_updated 
		WHERE action = 'source_transaction'
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
		trx.id,
		trx.client_id,
		cl.name AS client,
		cl.industry AS industry,
		trx.product_id,
		CASE WHEN prt.name IS NOT NULL 
			THEN prt.name 
			ELSE 'INTERNAL' 
			END
		AS department,
		pr.name AS product,
		pr.revenue_category AS revenue_category,
		trx.amount,
		trx.created_at
	FROM transaction AS trx
	JOIN client AS cl ON cl.id = trx.client_id
	JOIN product AS pr ON pr.id = trx.product_id
	LEFT JOIN partner AS prt ON prt.id = cl.department_id
	WHERE trx.id > %d;`, lastID)

	resEmp, err := conn.Read(query)
	if err != nil {
		logger.Error(err.Error())
	}

	var transactions []model.Transaction
	if resErr := json.Unmarshal([]byte(resEmp), &transactions); resErr != nil {
		logger.Error(resErr.Error())
	}

	if len(transactions) != 0 {
		for _, transaction := range transactions {
			ID, err := strconv.Atoi(transaction.ID)
			if err != nil {
				logger.Error(err.Error())
			}

			clientID, err := strconv.Atoi(transaction.ClientID)
			if err != nil {
				logger.Error(err.Error())
			}
			productID, err := strconv.Atoi(transaction.ProductID)
			if err != nil {
				logger.Error(err.Error())
			}
			amount, err := strconv.Atoi(transaction.Amount)
			if err != nil {
				logger.Error(err.Error())
			}

			queryInsert := fmt.Sprintf(
				`INSERT INTO source_transaction (
					source_id,
					client_id,
					client,
					industry,
					department,
					product_id,
					product,
					revenue_category,
					amount,
					created_at
				) 
				values (%d, %d, '%s', '%s', '%s', %d, '%s', '%s', %d, '%s')`,
				ID,
				clientID,
				transaction.Client,
				transaction.Industry,
				transaction.Department,
				productID,
				transaction.Product,
				transaction.RevenueCategory,
				amount,
				transaction.CreatedAt,
			)

			_, errs := conn.Store(queryInsert)
			if errs != nil {
				logger.Error(errs.Error())
			}
		}

		queryGetLastID := fmt.Sprintf(`
		SELECT id
		FROM transaction
		WHERE id > %d
		ORDER BY id DESC
		LIMIT 1`,
			lastID)

		resGetLast, err := conn.Read(queryGetLastID)
		if err != nil {
			logger.Error("[GET LAST SOURCE TRANSACTION]", err.Error())
		}

		var lastTitle []model.Title
		if resErr := json.Unmarshal([]byte(resGetLast), &lastTitle); resErr != nil {
			logger.Error("[LAST SOURCE TRANSACTION UNMARSHAL]", resErr.Error())
		}

		last, err := strconv.Atoi(lastTitle[0].ID)
		if err != nil {
			logger.Error("[LAST SOURCE TRANSACTION UNMARSHAL] failed convert: ", err.Error())
		}

		now := time.Now()

		queryUpdated := fmt.Sprintf(`
			INSERT INTO last_updated (action, last_id, created_at)
			VALUE('%s', %d, '%s')
		`, "source_transaction", last, now.Format("2006-01-02 15:04:05"))

		_, errs := conn.Store((queryUpdated))
		if errs != nil {
			logger.Error("[INSERT INTO LAST UPDATED]", errs.Error())
		}

		fmt.Println("[SOURCE TRANSACTION]: Success update data")
	} else {
		fmt.Println("[SOURCE TRANSACTION]: There is no new data")
	}

	conn.Close()

}
