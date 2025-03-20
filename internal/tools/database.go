package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	Username string
	AuthToken string
}

type CoinDetails struct {
	Coins int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoinDetails(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface,error){
	var database DatabaseInterface = &mockDb{}
	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil,err
	}
	return &database,nil
}