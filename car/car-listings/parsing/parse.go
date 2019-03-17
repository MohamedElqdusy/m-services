package parsing

import (
	"car-listings/models"
	"strings"
	"mime/multipart"
	"strconv"
	"bufio"
	"car-listings/utils"
	)

func ReadCSVFromHttpRequest(file multipart.File) ([]models.Car, error) {
	var cars []models.Car  
	s := bufio.NewScanner(file)
	for s.Scan() {
		car, err:= parseCarFromCSV(s.Text())
		if ( err!=nil ){
			return nil, err
		}
		cars = append(cars, car)
    }
    return cars, nil
}

func parseCarFromCSV(line string) (models.Car, error){
		var car models.Car
		var err error

		columns := strings.Split(line, ",")
		car.Code = columns[0] 
		car.Make = columns[1]
		car.Model = columns[2]

		car.Kw, err= strconv.ParseInt(columns[3], 0, 64)
		utils.HandleError(err)

		car.Year, err = strconv.ParseInt(columns[4], 0, 64)
		utils.HandleError(err)

		car.Color = columns[5]

		car.Price, err = strconv.ParseInt(columns[6], 0, 64)
		utils.HandleError(err)
		return car, err
}