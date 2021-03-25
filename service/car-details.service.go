package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	dto "github.com/kazuki0924/go-chi/dto/car"
)

var (
	ownerService     OwnerService = NewOwnerService()
	carService       CarService   = NewCarService()
	carDatachannel                = make(chan *http.Response)
	ownerDatachannel              = make(chan *http.Response)
)

type CarDetailsService interface {
	GetDetails() dto.CarDetails
}

type service struct{}

func NewCarDetailsService() CarDetailsService {
	return &service{}
}

func (*service) GetDetails() dto.CarDetails {
	go carService.FetchData()
	go ownerService.FetchData()

	car, _ := getCarData()
	owner, _ := getOwnerData()

	return dto.CarDetails{
		ID:             car.CarData.ID,
		Brand:          car.CarData.Brand,
		Model:          car.CarData.Model,
		Year:           car.CarData.Year,
		Vin:            car.CarData.Vin,
		OwnerFirstName: owner.OwnerData.FirstName,
		OwnerLastName:  owner.OwnerData.LastName,
		OwnerEmail:     owner.OwnerData.Email,
		OwnerJobTitle:  owner.OwnerData.JobTitle,
	}
}

func getCarData() (dto.MyFakeApiCar, error) {
	r1 := <-carDatachannel
	var car dto.MyFakeApiCar
	err := json.NewDecoder(r1.Body).Decode(&car)
	if err != nil {
		fmt.Print(err.Error())
		return car, err
	}
	return car, nil
}

func getOwnerData() (dto.MyFakeApiOwner, error) {
	r2 := <-ownerDatachannel
	var owner dto.MyFakeApiOwner
	err := json.NewDecoder(r2.Body).Decode(&owner)
	if err != nil {
		fmt.Print(err.Error())
		return owner, err
	}
	return owner, nil
}
