package main

import (
	"mockapi/datasource"
	"net/http"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
)

func checkValidAccount(bankid, accountid string) (bool, interface{}) {
	if bankMap, ok := datasource.Account[bankid]; ok {
		if account, ok := bankMap.(map[string]interface{})[accountid]; ok {
			return true, account
		} else {
			return false, nil
		}
	} else {
		return false, nil
	}
}

func checkValidAccountBiller(billerid, accountid string) (bool, *datasource.Biller) {
	if billerMap, ok := datasource.BillerAccount[billerid]; ok {
		if biller, ok := billerMap[accountid]; ok {
			return true, biller
		} else {
			return false, nil
		}
	} else {
		return false, nil
	}
}

func checkValidVa(VAId string) bool {
	for _, va := range datasource.VAList {
		if va.VAID == VAId {
			return true
		}
	}
	return false
}

func checkValidVaGenerated(VAId, VANumber string) (bool, *datasource.VAGen) {
	if vaMap, ok := datasource.VAGenerated[VAId]; ok {
		if va, ok := vaMap[VANumber]; ok {
			return true, va
		} else {
			return false, nil
		}
	} else {
		return false, nil
	}
}

func checkValidDeposito(depositoId string) (bool, *datasource.Deposito) {
	for _, dep := range datasource.DepositoMaster {
		if dep.DepositoID == depositoId {
			return true, &dep
		}
	}
	return false, nil
}

// https://dasarpemrogramangolang.novalagung.com/A-client-http-request-simple.html

var BillerCache []*datasource.Biller

func main() {
	r := gin.Default()

	apiGroup := r.Group("/v1/api")
	{
		transferGroup := apiGroup.Group("/transfer")
		{
			transferGroup.GET("/list-bank", func(c *gin.Context) {
				// no implemented
				c.JSON(200, gin.H{
					"data": datasource.BankMaster,
				})

				return
			})

			transferGroup.GET("/:bankid/:accountid", func(c *gin.Context) {
				// no implemented
				bankid := c.Param("bankid")
				accountid := c.Param("accountid")

				if ok, account := checkValidAccount(bankid, accountid); ok {
					c.JSON(200, gin.H{
						"account": account,
					})
				} else {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "account not found",
					})
					return
				}

			})

			transferGroup.POST("/", func(c *gin.Context) {
				// no implemented

				bodyStruct := struct {
					BankID    string
					AccountID string
					Amount    float64
				}{}

				err := c.BindJSON(&bodyStruct)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "error load payload",
						"error":   err.Error(),
					})

					return
				}

				if bodyStruct.Amount == 0 {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "amount cannot be zero!",
					})

					return
				}

				if ok, _ := checkValidAccount(bodyStruct.BankID, bodyStruct.AccountID); ok {
					c.JSON(200, gin.H{
						"message": "transfer success!",
					})
					return
				} else {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "account not found",
					})
					return
				}

			})
		}

		billerGroup := apiGroup.Group("/biller")
		{
			billerGroup.GET("/", func(ctx *gin.Context) {
				// no implemented
				ctx.JSON(200, gin.H{
					"data": datasource.BillerList,
				})
			})

			billerGroup.GET("/:biller_id/:biller_account", func(c *gin.Context) {
				// no implemented
				billerid := c.Param("biller_id")
				accountid := c.Param("biller_account")

				if ok, biller := checkValidAccountBiller(billerid, accountid); ok {
					c.JSON(200, gin.H{
						"biller": biller,
					})
				} else {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "biller not found",
					})
					return
				}
			})

			billerGroup.POST("/", func(c *gin.Context) {
				// no implemented
				bodyStruct := struct {
					BillerID  string
					AccountID string
					Amount    float64
				}{}

				err := c.BindJSON(&bodyStruct)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "error load payload",
						"error":   err.Error(),
					})

					return
				}

				if ok, biller := checkValidAccountBiller(bodyStruct.BillerID, bodyStruct.AccountID); ok {
					if biller.Paid {
						c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
							"message": "biller already paid",
						})
						return
					}

					if biller.Amount != bodyStruct.Amount {
						c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
							"message": "biller not found",
						})
						return
					}

					biller.Paid = true

					BillerCache = append(BillerCache, biller)

					c.JSON(200, gin.H{
						"message": "transfer success!",
					})
					return
				} else {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "biller not found",
					})
					return
				}
			})

			billerGroup.GET("/refresh", func(ctx *gin.Context) {
				for _, bill := range BillerCache {
					bill.Paid = false
				}

				// remove the data
				BillerCache = []*datasource.Biller{}
			})
		}

		vagroup := apiGroup.Group("/va")
		{
			vagroup.GET("/", func(ctx *gin.Context) {
				// no implement
				ctx.JSON(200, gin.H{
					"data": datasource.VAList,
				})
			})

			vagroup.POST("/:va_id", func(c *gin.Context) {
				// no implement
				va_id := c.Param("va_id")

				if checkValidVa(va_id) {
					va := faker.CCNumber()
					vaGenerated := &datasource.VAGen{
						VANumber: va,
						VAID:     va_id,
						Amount:   float64(datasource.GetAmount()),
					}

					if _, ok := datasource.VAGenerated[va_id]; !ok {
						datasource.VAGenerated[va_id] = make(map[string]*datasource.VAGen)
					}

					datasource.VAGenerated[va_id][va] = vaGenerated
					c.JSON(200, gin.H{
						"data": vaGenerated,
					})
					return
				} else {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "va not found",
					})
					return
				}
			})

			vagroup.POST("/", func(c *gin.Context) {
				bodyStruct := struct {
					VAID     string
					VANumber string
					Amount   float64
				}{}

				err := c.BindJSON(&bodyStruct)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "error load payload",
						"error":   err.Error(),
					})

					return
				}

				if ok, va := checkValidVaGenerated(bodyStruct.VAID, bodyStruct.VANumber); ok {

					if va.Paid {
						c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
							"message": "va already paid!",
						})
						return
					}

					if va.Amount != bodyStruct.Amount {
						c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
							"message": "va not found",
						})
						return
					}

					va.Paid = true

					c.JSON(200, gin.H{
						"message": "va paid success!",
					})
					return
				} else {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "va not found",
					})
					return
				}

			})

			vagroup.GET("/refresh", func(ctx *gin.Context) {

				// remove the data
				datasource.VAGenerated = make(map[string]map[string]*datasource.VAGen)
			})
		}

		depositogroup := apiGroup.Group("/deposito")
		{
			depositogroup.GET("/", func(ctx *gin.Context) {
				// no implement
				ctx.JSON(200, datasource.DepositoMaster)
			})

			depositogroup.GET("/:deposito_id/:account_id", func(ctx *gin.Context) {
				deposito_id := ctx.Param("deposito_id")
				account_id := ctx.Param("account_id")

				if deps, ok := datasource.DepositoAccountList[deposito_id][account_id]; ok {
					ctx.JSON(200, gin.H{
						"data": deps,
					})
					return
				}

				ctx.JSON(200, gin.H{
					"data": []datasource.DepositoAccount{},
				})
				return
				// no implement
			})

			depositogroup.POST("/", func(c *gin.Context) {
				bodyStruct := struct {
					DepositoID string
					AccountID  string
					Amount     float64
				}{}

				err := c.BindJSON(&bodyStruct)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "error load payload",
						"error":   err.Error(),
					})

					return
				}

				if bodyStruct.Amount == 0 {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "amount cannot be zero",
					})
					return
				}

				if ok, deposito := checkValidDeposito(bodyStruct.DepositoID); ok {

					account := &datasource.DepositoAccount{
						DepositoName:    deposito.Name,
						DepositoID:      deposito.DepositoID,
						AccountID:       bodyStruct.AccountID,
						DepositoAccount: faker.CCNumber(),
						Balance:         bodyStruct.Amount,
					}

					if _, ok := datasource.DepositoAccountList[account.DepositoID]; !ok {
						datasource.DepositoAccountList[account.DepositoID] = map[string][]*datasource.DepositoAccount{}
					}

					if _, ok := datasource.DepositoAccountList[account.DepositoID]; ok {

						datasource.DepositoAccountList[account.DepositoID][account.AccountID] = append(datasource.DepositoAccountList[account.DepositoID][account.AccountID], account)

					}

					c.JSON(200, gin.H{
						"message": "deposito created!",
					})
					return
				} else {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "deposito not found",
					})
					return
				}
			})
		}

	}

	r.Run(":5000") // listen and serve on 0.0.0.0:8080
}
