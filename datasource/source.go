package datasource

import (
	"log"
	"reflect"

	faker "github.com/bxcodec/faker/v3"
)

type BankAccount struct {
	BankID    string
	AccountID string
	Name      string
}

var Account map[string]interface{} = map[string]interface{}{
	"BRI": map[string]interface{}{
		"8888888888": BankAccount{
			BankID:    "BRI",
			AccountID: "8888888888",
			Name:      "anggito",
		},
		"8888888889": BankAccount{
			BankID:    "BRI",
			AccountID: "8888888889",
			Name:      "radit",
		},
		"8888888890": BankAccount{
			BankID:    "BRI",
			AccountID: "8888888890",
			Name:      "zoel",
		},
		"8888888891": BankAccount{
			BankID:    "BRI",
			AccountID: "8888888891",
			Name:      "yushi",
		},
	},
	"BCA": map[string]interface{}{
		"9188888888": BankAccount{
			BankID:    "BCA",
			AccountID: "9188888888",
			Name:      "putri",
		},
		"9288888889": BankAccount{
			BankID:    "BCA",
			AccountID: "9288888889",
			Name:      "kevin",
		},
		"9388888890": BankAccount{
			BankID:    "BCA",
			AccountID: "9388888890",
			Name:      "daffa",
		},
		"9488888891": BankAccount{
			BankID:    "BCA",
			AccountID: "9488888891",
			Name:      "huda",
		},
	},
}

type Biller struct {
	BillerID  string
	Name      string
	Amount    float64
	AccountID string
	Paid      bool
}

type BillerMaster struct {
	BillerID string
	Name     string
}

var BillerList = []BillerMaster{
	{
		BillerID: "012",
		Name:     "PDAM",
	},
	{
		BillerID: "30305",
		Name:     "PLN",
	},
	{
		BillerID: "014",
		Name:     "Telkom",
	},
}

func GetAmount() float32 {
	forKind := new(float32)
	amount, err := faker.GetPrice().Amount(reflect.ValueOf(forKind).Elem())
	if err != nil {
		log.Fatal(err)
	}

	return amount.(float32)
}

var BillerAccount map[string]map[string]*Biller = map[string]map[string]*Biller{
	"012": {
		"110000": &Biller{
			BillerID:  "012",
			Name:      "PDAM",
			Amount:    float64(GetAmount()),
			AccountID: "110000",
			Paid:      false,
		},
		"120000": &Biller{
			BillerID:  "012",
			Name:      "PDAM",
			Amount:    float64(GetAmount()),
			AccountID: "120000",
			Paid:      false,
		},
		"130000": &Biller{
			BillerID:  "012",
			Name:      "PDAM",
			Amount:    float64(GetAmount()),
			AccountID: "130000",
			Paid:      false,
		},
		"140000": &Biller{
			BillerID:  "012",
			Name:      "PDAM",
			Amount:    float64(GetAmount()),
			AccountID: "140000",
			Paid:      false,
		},
	},
	"30305": {
		"1100001": &Biller{
			BillerID:  "30305",
			Name:      "PLN",
			Amount:    float64(GetAmount()),
			AccountID: "1100001",
			Paid:      false,
		},
		"1200001": &Biller{
			BillerID:  "30305",
			Name:      "PLN",
			Amount:    float64(GetAmount()),
			AccountID: "1200001",
			Paid:      false,
		},
		"1300001": &Biller{
			BillerID:  "30305",
			Name:      "PLN",
			Amount:    float64(GetAmount()),
			AccountID: "1300001",
			Paid:      false,
		},
		"1400001": &Biller{
			BillerID:  "30305",
			Name:      "PLN",
			Amount:    float64(GetAmount()),
			AccountID: "1400001",
			Paid:      false,
		},
	},
	"014": {
		"1100002": &Biller{
			BillerID:  "014",
			Name:      "Telkom",
			Amount:    float64(GetAmount()),
			AccountID: "1100002",
			Paid:      false,
		},
		"1200002": &Biller{
			BillerID:  "014",
			Name:      "Telkom",
			Amount:    float64(GetAmount()),
			AccountID: "1200002",
			Paid:      false,
		},
		"1300002": &Biller{
			BillerID:  "014",
			Name:      "Telkom",
			Amount:    float64(GetAmount()),
			AccountID: "1300002",
			Paid:      false,
		},
		"1400002": &Biller{
			BillerID:  "014",
			Name:      "Telkom",
			Amount:    float64(GetAmount()),
			AccountID: "1400002",
			Paid:      false,
		},
	},
}

type VAMaster struct {
	VAID string
	Name string
}

type VAGen struct {
	VAID     string
	VANumber string
	Amount   float64
	Paid     bool
}

var VAList = []VAMaster{
	{
		VAID: "88789",
		Name: "Travelagent",
	},
	{
		VAID: "88708",
		Name: "Tokodia",
	},
}

var VAGenerated map[string]map[string]*VAGen = make(map[string]map[string]*VAGen)

type Deposito struct {
	Name       string
	DepositoID string
}

type DepositoAccount struct {
	DepositoName    string
	DepositoID      string
	DepositoAccount string
	AccountID       string
	Balance         float64
}

var DepositoMaster = []Deposito{
	{
		Name:       "MAXI",
		DepositoID: "MAXI",
	},
	{
		Name:       "FLEXI",
		DepositoID: "FLEXI",
	},
}

var DepositoAccountList map[string]map[string][]*DepositoAccount = make(map[string]map[string][]*DepositoAccount)
