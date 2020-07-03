package transformer

import (
	"encoding/json"
	"fmt"

	"github.com/nicosrgh/straw-hat/config/repository"
	"github.com/nicosrgh/straw-hat/lib/logger"
	"github.com/nicosrgh/straw-hat/model"
)

// GenderDimension ...
func GenderDimension(store *repository.Store) {
	fmt.Println("Get dimension product ...")

	query := "select DISTINCT gender from employee;"

	res, err := store.My.Read(query)
	if err != nil {
		logger.Error(err.Error())
	}

	var genders []model.Employee
	if resErr := json.Unmarshal([]byte(res), &genders); resErr != nil {
		logger.Error(resErr.Error())
	}

	length := len(genders)
	if length > 0 {
		i := 0
		for _, gender := range genders {
			queryExists := fmt.Sprintf("select * from ss_gender_dimension where name = '%s'", gender.Gender)

			res, err := store.My.Read(queryExists)
			if err != nil {
				logger.Error(err.Error())
			}

			var dimGenders []model.DimensionGender
			if resErr := json.Unmarshal([]byte(res), &dimGenders); resErr != nil {
				logger.Error(resErr.Error())
			}

			if len(dimGenders) <= 0 {
				queryStore := fmt.Sprintf("insert into ss_gender_dimension (name) values ('%s')", gender.Gender)
				_, err := store.My.Store(queryStore)
				if err != nil {
					logger.Error(err.Error())
				}

				i++
			}
		}
		fmt.Println("Inserted dimension gender: ", i)
	}
}
