package sqlapi

import (
	"log"
	"projectbotticket/types/apitypes"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // here
)

type API struct {
	db *sqlx.DB
}

func NewAPI(db *sqlx.DB) *API {
	return &API{
		db: db,
	}
}

// Block GET

// GetUserByID ...
func (api *API) GetUserByID(idUser string) (*apitypes.UserRow, error) {

	userRow := []apitypes.UserRow{}

	err := api.db.Select(&userRow, "SELECT * FROM prj_user WHERE userid = $1", idUser)
	if err != nil {
		log.Println("GetUserByID api.db.Select failed with an error: ", err.Error())
		return nil, err
	}

	if len(userRow) == 1 {
		return &userRow[0], err
	}
	return nil, err
}

// GetUserByName ...
func (api *API) GetUserByName(nameUser string) (*apitypes.UserRow, error) {

	userRow := []apitypes.UserRow{}

	err := api.db.Select(&userRow, "SELECT * FROM prj_user WHERE nameuser = $1", nameUser)
	if err != nil {
		log.Println("GetUserByName api.db.Select failed with an error: ", err.Error())
		return nil, err
	}
	if len(userRow) == 1 {
		return &userRow[0], err
	}

	return nil, err
}

// GetExecuterById ...
func (api *API) GetExecuterById(executorId string) (*apitypes.ExecutorRow, error) {

	executerRow := []apitypes.ExecutorRow{}

	err := api.db.Select(&executerRow, "SELECT * FROM prj_executor WHERE executorid = $1", executorId)
	if err != nil {
		log.Println("GetExecuterById api.db.Select failed with an error: ", err.Error())
		return nil, err
	}

	if len(executerRow) == 1 {
		return &executerRow[0], err
	}

	return nil, err
}

// GetExecuterByName ...
func (api *API) GetExecuterByName(executerName string) (*apitypes.ExecutorRow, error) {
	executerRow := []apitypes.ExecutorRow{}

	err := api.db.Select(&executerRow, "SELECT * FROM prj_executor WHERE executorname = $1", executerName)
	if err != nil {
		log.Println("GetExecuterByName api.db.Select failed with an error: ", err.Error())
		return nil, err
	}
	if len(executerRow) == 1 {
		return &executerRow[0], err
	}

	return nil, err
}

// GetOrderById ...
func (api *API) GetOrderById(orderId string) (*apitypes.OrderRow, error) {
	orderRow := []apitypes.OrderRow{}

	err := api.db.Select(&orderRow, "SELECT * FROM prj_order WHERE orderid = $1", orderId)
	if err != nil {
		log.Println("GetOrderById api.db.Select failed with an error: ", err.Error())
		return nil, err
	}

	if len(orderRow) == 1 {
		return &orderRow[0], err
	}

	return nil, err
}

// GetStatusById ...
func (api *API) GetStatusById(statusId string) (*apitypes.StatusRow, error) {
	statusRow := []apitypes.StatusRow{}

	err := api.db.Select(&statusRow, "SELECT * FROM prj_status WHERE statusid = $1", statusId)
	if err != nil {
		log.Println("GetStatusById api.db.Select failed with an error: ", err.Error())
		return nil, err
	}

	if len(statusRow) == 1 {
		return &statusRow[0], err
	}

	return nil, err
}

// GetUserOrderExecutorById ...
func (api *API) GetUserOrderExecutorById(userOrderExecutorId string) (*apitypes.UserOrderExecutorRow, error) {
	userOrderExecutorRow := []apitypes.UserOrderExecutorRow{}

	err := api.db.Select(userOrderExecutorRow, "SELECT * FROM link_userorderexecutor WHERE linkid = $1", userOrderExecutorId)
	if err != nil {
		log.Println("GetUserOrderExecutorById api.db.Select failed with an error: ", err.Error())
		return nil, err
	}

	if len(userOrderExecutorRow) == 1 {
		return &userOrderExecutorRow[0], err
	}

	return nil, err
}

// GetUserOrderExecutorByUserIdExecutorId ...
func (api *API) GetUserOrderExecutorByUserIdExecutorId(userId string, executorid string) (*apitypes.UserOrderExecutorRow, error) {
	userOrderExecutorRow := []apitypes.UserOrderExecutorRow{}

	err := api.db.Select(userOrderExecutorRow, "SELECT * FROM link_userorderexecutor WHERE userid = $1 AND executorid = $2", userId, executorid)
	if err != nil {
		log.Println("GetUserOrderExecutorByUserIdExecutorId api.db.Select failed with an error: ", err.Error())
		return nil, err
	}

	if len(userOrderExecutorRow) == 1 {
		return &userOrderExecutorRow[0], err
	}

	return nil, err
}

// Block SET

// SetNewUser ...
func (api *API) SetNewUser(userName, userId, chatId string) (*apitypes.UserRow, error) {
	user, err := api.GetUserByID(userId)

	if err != nil {
		log.Println("SetNewUser api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}
	if user != nil {
		return nil, err
	}

	tx := api.db.MustBegin()
	tx.MustExec(`INSERT INTO prj_user ("userid", "nameuser", "chatid") VALUES ($1, $2, $3)`,
		userId, userName, chatId)
	tx.Commit()

	user, err = api.GetUserByID(userId)

	if err != nil {
		log.Println("AddNewUser GetUserByID failed with an error: ", err.Error())
		return nil, err
	}

	return user, err
}

// SetNewExecutor ...
func (api *API) SetNewExecutor(executorid, executorname, executorpasword string) (*apitypes.ExecutorRow, error) {
	executor, err := api.GetExecuterById(executorid)

	if err != nil {
		log.Println("SetNewExecutor api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}
	if executor != nil {
		return nil, err
	}

	tx := api.db.MustBegin()
	tx.MustExec(`INSERT INTO prj_executor ("executorid", "executorname", "executorpasword") VALUES ($1, $2, $3)`,
		executorid, executorname, executorpasword)
	tx.Commit()

	executor, err = api.GetExecuterById(executorid)

	if err != nil {
		log.Println("SetNewExecutor GetUserByID failed with an error: ", err.Error())
		return nil, err
	}

	return executor, err
}

// SetNewOrder ...
func (api *API) SetNewOrder(orderid, description string) (*apitypes.OrderRow, error) {
	order, err := api.GetOrderById(orderid)

	if err != nil {
		log.Println("SetNewOrder api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}
	if order != nil {
		return nil, err
	}

	status, err := api.GetStatusById("")
	if err != nil {
		log.Println("SetNewOrder GetStatusById api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}
	if status == nil {
		return nil, err
	}

	today := time.Now()
	orderstarttime := today.Add(10 * time.Minute).Format("2006/1/2 15:04")

	tx := api.db.MustBegin()
	tx.MustExec(`INSERT INTO prj_order ("orderid", "orderdescription", "statusid", "orderstarttime") VALUES ($1, $2, $3, $4)`,
		orderid, description, status.StatusId, orderstarttime)
	tx.Commit()

	order, err = api.GetOrderById(orderid)

	if err != nil {
		log.Println("SetNewOrder GetUserByID failed with an error: ", err.Error())
		return nil, err
	}

	return order, err
}

// SetUserOrderExecutor ...
func (api *API) SetUserOrderExecutor(linkid, userid, orderid, executorid string) (*apitypes.UserOrderExecutorRow, error) {
	userOrderExecutor, err := api.GetUserOrderExecutorById(linkid)
	if err != nil {
		log.Println("SetUserOrderExecutor api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}
	if userOrderExecutor != nil {
		return nil, err
	}

	tx := api.db.MustBegin()
	tx.MustExec(`INSERT INTO link_userorderexecutor ("linkid", "userid", "orderid", "executorid") VALUES ($1, $2, $3, $4)`,
		linkid, userid, orderid, executorid)
	tx.Commit()

	userOrderExecutor, err = api.GetUserOrderExecutorById(linkid)

	if err != nil {
		log.Println("SetUserOrderExecutor GetUserByID failed with an error: ", err.Error())
		return nil, err
	}

	return userOrderExecutor, err
}

// Block UPDATE

// UpdateUserName ...
func (api *API) UpdateUserName(userid, nameuser string) (*apitypes.UserRow, error) {
	userRow, err := api.GetUserByID(userid)
	if err != nil {
		log.Println("UpdateUserName api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}
	if userRow == nil {
		return nil, err
	}

	tx := api.db.MustBegin()
	tx.MustExec(`UPDATE prj_user SET nameuser=$1 WHERE userid=$2`, nameuser, userid)
	tx.Commit()

	userRow, err = api.GetUserByID(userid)
	if err != nil {
		log.Println("UpdateUserName api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}

	return userRow, err
}

// UpdateExecuterName ...
func (api *API) UpdateExecuterName(executorId, executorName string) (*apitypes.ExecutorRow, error) {
	executorRow, err := api.GetExecuterById(executorId)
	if err != nil {
		log.Println("UpdateExecuterName api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}
	if executorRow == nil {
		return nil, err
	}

	tx := api.db.MustBegin()
	tx.MustExec(`UPDATE prj_executor SET executorname=$1 WHERE executorid=$2`, executorName, executorId)
	tx.Commit()

	executorRow, err = api.GetExecuterById(executorId)
	if err != nil {
		log.Println("UpdateExecuterName api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}

	return executorRow, err
}
