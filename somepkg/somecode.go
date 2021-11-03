package somepkg

import "demomockery/somepkg/models"

type DataProvider interface {
	GetData(key int) ([]models.DataRec, error)
}

func SomeLogic(data DataProvider) ([]models.DataRec, error) {
	var rec []models.DataRec
	rec, err := data.GetData(1)
	if err != nil {
		return nil, err
	}
	return rec, nil
}
