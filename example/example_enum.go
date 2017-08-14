// Code generated by go-enum
// DO NOT EDIT!

package example

import (
	"fmt"
	"strings"
)

const (
	// AnimalCat is a Animal of type Cat
	AnimalCat Animal = iota
	// AnimalDog is a Animal of type Dog
	AnimalDog
	// AnimalFish is a Animal of type Fish
	AnimalFish
)

const _AnimalName = "CatDogFish"

var _AnimalMap = map[Animal]string{
	0: _AnimalName[0:3],
	1: _AnimalName[3:6],
	2: _AnimalName[6:10],
}

func (i Animal) String() string {
	if str, ok := _AnimalMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Animal(%d)", i)
}

var _AnimalValue = map[string]Animal{
	_AnimalName[0:3]:                   0,
	strings.ToLower(_AnimalName[0:3]):  0,
	_AnimalName[3:6]:                   1,
	strings.ToLower(_AnimalName[3:6]):  1,
	_AnimalName[6:10]:                  2,
	strings.ToLower(_AnimalName[6:10]): 2,
}

// ParseAnimal attempts to convert a string to a Animal
func ParseAnimal(name string) (Animal, error) {
	if x, ok := _AnimalValue[name]; ok {
		return Animal(x), nil
	}
	return Animal(0), fmt.Errorf("%s is not a valid Animal", name)
}

func (x *Animal) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *Animal) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseAnimal(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// ModelToyota is a Model of type Toyota
	ModelToyota Model = iota
	// Skipped value
	_
	// ModelChevy is a Model of type Chevy
	ModelChevy
	// Skipped value
	_
	// ModelFord is a Model of type Ford
	ModelFord
	// Skipped value
	_
	// ModelTesla is a Model of type Tesla
	ModelTesla
	// Skipped value
	_
	// ModelHyundai is a Model of type Hyundai
	ModelHyundai
	// Skipped value
	_
	// ModelNissan is a Model of type Nissan
	ModelNissan
	// Skipped value
	_
	// ModelJaguar is a Model of type Jaguar
	ModelJaguar
	// Skipped value
	_
	// ModelAudi is a Model of type Audi
	ModelAudi
	// Skipped value
	_
	// ModelBMW is a Model of type BMW
	ModelBMW
	// Skipped value
	_
	// ModelMercedes is a Model of type Mercedes
	ModelMercedes
	// Skipped value
	_
	// ModelVolkswagon is a Model of type Volkswagon
	ModelVolkswagon
)

const _ModelName = "ToyotaChevyFordTeslaHyundaiNissanJaguarAudiBMWMercedesVolkswagon"

var _ModelMap = map[Model]string{
	0:  _ModelName[0:6],
	2:  _ModelName[6:11],
	4:  _ModelName[11:15],
	6:  _ModelName[15:20],
	8:  _ModelName[20:27],
	10: _ModelName[27:33],
	12: _ModelName[33:39],
	14: _ModelName[39:43],
	16: _ModelName[43:46],
	18: _ModelName[46:54],
	20: _ModelName[54:64],
}

func (i Model) String() string {
	if str, ok := _ModelMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Model(%d)", i)
}

var _ModelValue = map[string]Model{
	_ModelName[0:6]:                    0,
	strings.ToLower(_ModelName[0:6]):   0,
	_ModelName[6:11]:                   2,
	strings.ToLower(_ModelName[6:11]):  2,
	_ModelName[11:15]:                  4,
	strings.ToLower(_ModelName[11:15]): 4,
	_ModelName[15:20]:                  6,
	strings.ToLower(_ModelName[15:20]): 6,
	_ModelName[20:27]:                  8,
	strings.ToLower(_ModelName[20:27]): 8,
	_ModelName[27:33]:                  10,
	strings.ToLower(_ModelName[27:33]): 10,
	_ModelName[33:39]:                  12,
	strings.ToLower(_ModelName[33:39]): 12,
	_ModelName[39:43]:                  14,
	strings.ToLower(_ModelName[39:43]): 14,
	_ModelName[43:46]:                  16,
	strings.ToLower(_ModelName[43:46]): 16,
	_ModelName[46:54]:                  18,
	strings.ToLower(_ModelName[46:54]): 18,
	_ModelName[54:64]:                  20,
	strings.ToLower(_ModelName[54:64]): 20,
}

// ParseModel attempts to convert a string to a Model
func ParseModel(name string) (Model, error) {
	if x, ok := _ModelValue[name]; ok {
		return Model(x), nil
	}
	return Model(0), fmt.Errorf("%s is not a valid Model", name)
}

func (x *Model) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *Model) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseModel(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
