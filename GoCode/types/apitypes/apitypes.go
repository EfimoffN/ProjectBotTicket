package apitypes

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// StoreDB ...
type StoreDB struct {
	DB *sqlx.DB
}

// UserRow
type UserRow struct {
	UserId   string `db:"userid"`
	NameUser string `db:"nameuser"`
	ChatId   string `db:"chatid"`
}

// ExecutorRow
type ExecutorRow struct {
	ExecutorId      string `db:"executorid"`
	ExecutorName    string `db:"executorname"`
	ExecutorPasword string `db:"executorpasword"`
}

// StatusRow
type StatusRow struct {
	StatusId          string `db:"statusid"`
	StatusCode        string `db:"statuscode"`
	StatusDescription string `db:"statusdescription"`
}

// OrderRow
type OrderRow struct {
	OrderId          string    `db:"orderid"`
	OrderNumber      int       `db:"ordernumber"`
	OrderDescription string    `db:"orderdescription"`
	StatusId         string    `db:"statusid"`
	OrderStartTime   time.Time `db:"orderstarttime"`
	OrderStopTime    time.Time `db:"orderstoptime"`
}

// UserOrderExecutorRow
type UserOrderExecutorRow struct {
	LinkId     string `db:"linkid"`
	UserId     string `db:"userid"`
	OrderId    string `db:"orderid"`
	ExecutorId string `db:"executorid"`
}
