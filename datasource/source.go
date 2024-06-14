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

type Bank struct {
	Name   string
	BankID string
}

var BankMaster = []Bank{
	{
		Name:   "BRI",
		BankID: "BRI",
	},
	{
		Name:   "BCA",
		BankID: "BCA",
	},
}

var cacheCCGen = make(map[string]string)

func genAccountId(key string) string {
	if val, ok := cacheCCGen[key]; ok {
		return val
	} else {
		cacheCCGen[key] = faker.CCNumber()
		return cacheCCGen[key]
	}

}

var Account map[string]interface{} = map[string]interface{}{
	"BRI": map[string]interface{}{
		genAccountId("bri_1"): BankAccount{
			BankID:    "BRI",
			AccountID: genAccountId("bri_1"),
			Name:      "anggito",
		},
		genAccountId("bri_2"): BankAccount{
			BankID:    "BRI",
			AccountID: genAccountId("bri_2"),
			Name:      "radit",
		},
		genAccountId("bri_3"): BankAccount{
			BankID:    "BRI",
			AccountID: genAccountId("bri_3"),
			Name:      "zoel",
		},
		genAccountId("bri_4"): BankAccount{
			BankID:    "BRI",
			AccountID: genAccountId("bri_4"),
			Name:      "yushi",
		},
	},
	"BCA": map[string]interface{}{
		genAccountId("bca_1"): BankAccount{
			BankID:    "BCA",
			AccountID: genAccountId("bca_1"),
			Name:      "putri",
		},
		genAccountId("bca_2"): BankAccount{
			BankID:    "BCA",
			AccountID: genAccountId("bca_2"),
			Name:      "kevin",
		},
		genAccountId("bca_3"): BankAccount{
			BankID:    "BCA",
			AccountID: genAccountId("bca_3"),
			Name:      "daffa",
		},
		genAccountId("bca_4"): BankAccount{
			BankID:    "BCA",
			AccountID: genAccountId("bca_4"),
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
		genAccountId("012_1"): &Biller{
			BillerID:  "012",
			Name:      "PDAM",
			Amount:    float64(GetAmount()),
			AccountID: genAccountId("012_1"),
			Paid:      false,
		},
		genAccountId("012_2"): &Biller{
			BillerID:  "012",
			Name:      "PDAM",
			Amount:    float64(GetAmount()),
			AccountID: genAccountId("012_2"),
			Paid:      false,
		},
		genAccountId("012_3"): &Biller{
			BillerID:  "012",
			Name:      "PDAM",
			Amount:    float64(GetAmount()),
			AccountID: genAccountId("012_3"),
			Paid:      false,
		},
		genAccountId("012_4"): &Biller{
			BillerID:  "012",
			Name:      "PDAM",
			Amount:    float64(GetAmount()),
			AccountID: genAccountId("012_4"),
			Paid:      false,
		},
	},
	"30305": {
		genAccountId("30305_1"): &Biller{
			BillerID:  "30305",
			Name:      "PLN",
			Amount:    float64(GetAmount()),
			AccountID: genAccountId("30305_1"),
			Paid:      false,
		},
		genAccountId("30305_2"): &Biller{
			BillerID:  "30305",
			Name:      "PLN",
			Amount:    float64(GetAmount()),
			AccountID: genAccountId("30305_2"),
			Paid:      false,
		},
		genAccountId("30305_3"): &Biller{
			BillerID:  "30305",
			Name:      "PLN",
			Amount:    float64(GetAmount()),
			AccountID: genAccountId("30305_3"),
			Paid:      false,
		},
		genAccountId("30305_4"): &Biller{
			BillerID:  "30305",
			Name:      "PLN",
			Amount:    float64(GetAmount()),
			AccountID: genAccountId("30305_4"),
			Paid:      false,
		},
	},
	"014": {
		genAccountId("014_1"): &Biller{
			BillerID:  "014",
			Name:      "Telkom",
			Amount:    float64(GetAmount()),
			AccountID: genAccountId("014_1"),
			Paid:      false,
		},
		genAccountId("014_2"): &Biller{
			BillerID:  "014",
			Name:      "Telkom",
			Amount:    float64(GetAmount()),
			AccountID: genAccountId("014_2"),
			Paid:      false,
		},
		genAccountId("014_3"): &Biller{
			BillerID:  "014",
			Name:      "Telkom",
			Amount:    float64(GetAmount()),
			AccountID: genAccountId("014_3"),
			Paid:      false,
		},
		genAccountId("014_4"): &Biller{
			BillerID:  "014",
			Name:      "Telkom",
			Amount:    float64(GetAmount()),
			AccountID: genAccountId("014_4"),
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
