package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type portfolio struct {
	TotalBalance    float64         `json:"totalBalnce"`
	TotalInvestment float64         `json:"total Investment"`
	PortfolioList   []portfolioList `json:"PortfolioList"`
}

type portfolioList struct {
	CurrencyId int     `json:"currencyId"`
	Holdings   int     `json:"holdings"`
	Price      float64 `json:"price"`
	Gains      float64 `json:"gains"`
}

type portfolioTimeStamp struct {
	TimeStamp string    `json:"timeStamp"`
	Portfolio portfolio `json:"portfolio"`
}

var myPortfolioList = []portfolioList{
	{
		CurrencyId: 1,
		Holdings:   4,
		Price:      19231.0,
		Gains:      10.5,
	},
	{
		CurrencyId: 2,
		Holdings:   100,
		Price:      580.0,
		Gains:      1000.20,
	},
}

var myPortfolio = portfolio{
	TotalBalance: 1056.99, TotalInvestment: 300.0, PortfolioList: myPortfolioList}

var TimeStampList = []portfolioTimeStamp{
	{TimeStamp: "01-01-2021",
		Portfolio: myPortfolio},
	{TimeStamp: "01-02-2021",
		Portfolio: myPortfolio},
}

func GetPortfolio(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, myPortfolio)
}

func GetMyPortfolioValue(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, TimeStampList)
}

func main() {

	router := gin.Default()

	// portfolio
	router.GET("/portfolio/:uid", GetPortfolio)
	router.GET("/portfolio/:uid/timeline", GetMyPortfolioValue)

	//  watchlist
	router.GET("/watchlist/:uid", GetWatchList)
	router.PUT("/watchlist/:uid/:cid", UpdateWatchList)

	// transactions
	router.GET("/transactions/:userId", GetTransactions)
	router.GET("/transactions/:userId/:transactionId", GetTransactionById)
	router.POST("/transactions/:userId", AddTransaction)

	router.Run("localhost:3000")
}

// ------------------------ WatchList-----------------------------------

type CryptoCurrency struct {
	CryptoCurrencyId string `json:"cryptoCurrencyId"`
}

var watchlist = []CryptoCurrency{
	{CryptoCurrencyId: "BTC"},
	{CryptoCurrencyId: "ETH"},
}

func GetWatchList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, watchlist)
}

func UpdateWatchList(c *gin.Context) {

	var newCryptoCurrencyId CryptoCurrency

	// Call BindJSON to bind the received JSON to
	// newCryptoCurrencyId.
	if err := c.BindJSON(&newCryptoCurrencyId); err != nil {
		return
	}

	watchlist = append(watchlist, newCryptoCurrencyId)
	c.IndentedJSON(http.StatusOK, watchlist)
}

// ------------ Transactions ----------------------------------------

type Transaction struct {
	Id          string  `json:"id"`
	CurrencyId  string  `json:"currencyId"`
	Units       int     `json:"units"`
	TotalValue  float64 `json:"totalValue"`
	Status      bool    `json:"status"`
	TimeStamp   string  `json:"timeStamp"`
	PaymentType string  `json:"paymentType"`
	PaymentMode string  `json:"paymentMode"`
}

var TransactionList = []Transaction{
	{Id: "1", CurrencyId: "ETH", Units: 200, TotalValue: 1700.80, Status: false, TimeStamp: "01-01-2021", PaymentType: "Buy", PaymentMode: "CARD PAYMENT"},
	{Id: "2", CurrencyId: "XYZ", Units: 200, TotalValue: 1300.00, Status: false, TimeStamp: "01-03-2021", PaymentType: "Buy", PaymentMode: "ONLINE PAYMENT"},
}

func GetTransactions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, TransactionList)
}

func GetTransactionById(c *gin.Context) {

	id := c.Param("transactionId")
	for _, a := range TransactionList {
		if a.Id == id {
			fmt.Println(a)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

}
func AddTransaction(c *gin.Context) {
	var newTransaction Transaction

	// Call BindJSON to bind the received JSON to
	// newTransaction
	if err := c.BindJSON(&newTransaction); err != nil {
		return
	}

	newTransaction.Id = "3"
	// Add the new Transaction to the slice.
	TransactionList = append(TransactionList, newTransaction)
	c.IndentedJSON(http.StatusCreated, TransactionList)
}
