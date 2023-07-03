package modules

import (
	"errors"
	"github.com/purawaktra/semeru1-go/entities"
)

func (uc Argopuro1Usecase) SelectAccountByFirstName(firstName string, offset int) ([]Accounts, error) {
	// create check input on account-first name and offset
	if firstName == "" {
		return nil, errors.New("firstName can not be empty")
	}
	if offset < 0 {
		return nil, errors.New("offset can not be negative")
	}

	// convert input to entity
	account := entities.Accounts{FirstName: firstName}

	// call repo for the account first name
	accounts, err := uc.repo.SelectAccountByFirstName(account, uint(offset))

	// check for error on call usecase
	if err != nil {
		return nil, err
	}

	// convert entity to dto
	results := make([]Accounts, 0)
	for _, account := range accounts {
		// get city data
		// convert input to entity
		city := entities.Cities{CityId: account.City}
		cities, err := uc.repo.SelectCityById(city, 0)

		// check for error on call usecase
		if err != nil || cities == nil {
			cities = []entities.Cities{{Name: ""}}
		}

		// get province data
		// convert input to entity
		province := entities.Provinces{ProvinceId: account.Province}
		provinces, err := uc.repo.SelectProvinceById(province, 0)

		// check for error on call usecase
		if err != nil || provinces == nil {
			provinces = []entities.Provinces{{Name: ""}}
		}

		results = append(results, Accounts{
			AccountId:    account.AccountId,
			FirstName:    account.FirstName,
			LastName:     account.LastName,
			Address:      account.Address,
			City:         cities[0].Name,
			Province:     provinces[0].Name,
			Zipcode:      account.Zipcode,
			EmailAddress: account.EmailAddress,
			PhoneNumber:  account.PhoneNumber,
		})
	}
	// create return
	return results, nil
}

func (uc Argopuro1Usecase) SelectAccountByLastName(lastName string, offset int) ([]Accounts, error) {
	// create check input on account last name and offset
	if lastName == "" {
		return nil, errors.New("lastName can not be empty")
	}
	if offset < 0 {
		return nil, errors.New("offset can not be negative")
	}

	// convert input to entity
	account := entities.Accounts{LastName: lastName}

	// call repo for the account last name
	accounts, err := uc.repo.SelectAccountByLastName(account, uint(offset))

	// check for error on call usecase
	if err != nil {
		return nil, err
	}

	// convert entity to dto
	result := make([]Accounts, 0)
	for _, account := range accounts {
		// get city data
		// convert input to entity
		city := entities.Cities{CityId: account.City}
		cities, err := uc.repo.SelectCityById(city, 0)

		// check for error on call usecase
		if err != nil || cities == nil {
			cities = []entities.Cities{{Name: ""}}
		}

		// get province data
		// convert input to entity
		province := entities.Provinces{ProvinceId: account.Province}
		provinces, err := uc.repo.SelectProvinceById(province, 0)

		// check for error on call usecase
		if err != nil || provinces == nil {
			provinces = []entities.Provinces{{Name: ""}}
		}

		result = append(result, Accounts{
			AccountId:    account.AccountId,
			FirstName:    account.FirstName,
			LastName:     account.LastName,
			Address:      account.Address,
			City:         cities[0].Name,
			Province:     provinces[0].Name,
			Zipcode:      account.Zipcode,
			EmailAddress: account.EmailAddress,
			PhoneNumber:  account.PhoneNumber,
		})
	}
	// create return
	return result, nil
}

func (uc Argopuro1Usecase) SelectAllAccount(offset int) ([]Accounts, error) {
	// create check input offset
	if offset < 0 {
		return nil, errors.New("offset can not be negative")
	}

	// call repo for all accounts
	accounts, err := uc.repo.SelectAllAccount(uint(offset))

	// check for error on call usecase
	if err != nil {
		return nil, err
	}

	// convert entity to dto
	result := make([]Accounts, 0)
	for _, account := range accounts {
		// get city data
		// convert input to entity
		city := entities.Cities{CityId: account.City}
		cities, err := uc.repo.SelectCityById(city, 0)

		// check for error on call usecase
		if err != nil || cities == nil {
			cities = []entities.Cities{{Name: ""}}
		}

		// get province data
		// convert input to entity
		province := entities.Provinces{ProvinceId: account.Province}
		provinces, err := uc.repo.SelectProvinceById(province, 0)

		// check for error on call usecase
		if err != nil || provinces == nil {
			provinces = []entities.Provinces{{Name: ""}}
		}

		result = append(result, Accounts{
			AccountId:    account.AccountId,
			FirstName:    account.FirstName,
			LastName:     account.LastName,
			Address:      account.Address,
			City:         cities[0].Name,
			Province:     provinces[0].Name,
			Zipcode:      account.Zipcode,
			EmailAddress: account.EmailAddress,
			PhoneNumber:  account.PhoneNumber,
		})
	}
	// create return
	return result, nil
}

func (uc Argopuro1Usecase) InsertMultipleAccount(bodies []Accounts) ([]Accounts, []error) {
	// create check input on account last name and offset
	errs := make([]error, 0)
	for _, body := range bodies {
		if body.EmailAddress == "" {
			errs = append(errs, errors.New("emailAddress can not be empty"))
		} else {
			errs = append(errs, nil)
		}
	}

	results := make([]Accounts, 0)
	// proceed with nil error
	for index, err := range errs {
		// skip if err occurs
		// an insert nil account if err occurs
		if err != nil {
			results = append(results, Accounts{})
			continue
		}

		// convert city input to entity
		city := entities.Cities{Name: bodies[index].City}

		// convert city name to id
		cities, err := uc.repo.SelectCityByName(city, 0)

		// check for error on call usecase
		if err != nil || cities == nil {
			cities = []entities.Cities{{Name: ""}}
		}

		// convert province input to entity
		province := entities.Provinces{Name: bodies[index].Province}

		// convert province name to id
		provinces, err := uc.repo.SelectProvinceById(province, 0)

		// check for error on call usecase
		if err != nil || provinces == nil {
			provinces = []entities.Provinces{{Name: ""}}
		}

		// convert input to entity
		account := entities.Accounts{
			FirstName:    bodies[index].FirstName,
			LastName:     bodies[index].LastName,
			Address:      bodies[index].Address,
			City:         cities[0].CityId,
			Province:     provinces[0].ProvinceId,
			Zipcode:      bodies[index].Zipcode,
			EmailAddress: bodies[index].EmailAddress,
			PhoneNumber:  bodies[index].PhoneNumber,
		}

		// call repo for the insert account
		account, errInsert := uc.repo.InsertSingleAccount(account)

		// check for error on call usecase
		// skip if err occurs
		// an insert nil account if err occurs
		if errInsert != nil {
			errs[index] = errInsert
			results = append(results, Accounts{})
			continue
		}

		// convert entity to dto
		result := Accounts{
			AccountId:    account.AccountId,
			FirstName:    account.FirstName,
			LastName:     account.LastName,
			Address:      account.Address,
			City:         cities[0].Name,
			Province:     provinces[0].Name,
			Zipcode:      account.Zipcode,
			EmailAddress: account.EmailAddress,
			PhoneNumber:  account.PhoneNumber,
		}

		// add to array
		results = append(results, result)
	}

	// create return
	return results, errs
}

func (uc Argopuro1Usecase) UpdateMultipleAccountById(bodies []Accounts) ([]Accounts, []error) {
	// create check input on account last name and offset
	errs := make([]error, 0)
	for _, body := range bodies {
		if body.AccountId < 1 {
			errs = append(errs, errors.New("accountId can not be nil, negative or zero"))
		} else if body.EmailAddress == "" {
			errs = append(errs, errors.New("emailAddress can not be empty"))
		} else {
			errs = append(errs, nil)
		}
	}

	results := make([]Accounts, 0)
	// proceed with nil error
	for index, err := range errs {
		// skip if err occurs
		// an insert nil account if err occurs
		if err != nil {
			results = append(results, Accounts{})
			continue
		}

		// convert city input to entity
		city := entities.Cities{Name: bodies[index].City}

		// convert city name to id
		cities, err := uc.repo.SelectCityByName(city, 0)

		// check for error on call usecase
		if err != nil || cities == nil {
			cities = []entities.Cities{{Name: ""}}
		}

		// convert province input to entity
		province := entities.Provinces{Name: bodies[index].Province}

		// convert province name to id
		provinces, err := uc.repo.SelectProvinceById(province, 0)

		// check for error on call usecase
		if err != nil || provinces == nil {
			provinces = []entities.Provinces{{Name: ""}}
		}

		// convert input to entity
		account := entities.Accounts{
			FirstName:    bodies[index].FirstName,
			LastName:     bodies[index].LastName,
			Address:      bodies[index].Address,
			City:         cities[0].CityId,
			Province:     provinces[0].ProvinceId,
			Zipcode:      bodies[index].Zipcode,
			EmailAddress: bodies[index].EmailAddress,
			PhoneNumber:  bodies[index].PhoneNumber,
		}

		// call repo for the insert account
		account, errUpdate := uc.repo.UpdateSingleAccountById(account)

		// check for error on call usecase
		// skip if err occurs
		// an insert nil account if err occurs
		if errUpdate != nil {
			errs[index] = errUpdate
			results = append(results, Accounts{})
			continue
		}

		// convert entity to dto
		result := Accounts{
			AccountId:    account.AccountId,
			FirstName:    account.FirstName,
			LastName:     account.LastName,
			Address:      account.Address,
			City:         cities[0].Name,
			Province:     provinces[0].Name,
			Zipcode:      account.Zipcode,
			EmailAddress: account.EmailAddress,
			PhoneNumber:  account.PhoneNumber,
		}

		// add to array
		results = append(results, result)
	}

	// create return
	return results, errs
}
