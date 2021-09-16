package sqlapi

import (
	"context"
	"log"
	"projectbotticket/types/apitypes"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // here
	"github.com/pkg/errors"
)

type API struct {
	db *sqlx.DB
}

func NewAPI(db *sqlx.DB) *API {
	return &API{
		db: db,
	}
}

// Block test

// GetTestValue ...
func (api *API) GetTestValue(val int) int {
	var v = val * 2

	return v
}

// Block GET

// GetUserByID ...
func (api *API) GetUserByID(idUser string) (*apitypes.UserRow, error) {

	userRow := []apitypes.UserRow{}

	const query = `SELECT * FROM prj_user WHERE userid = $1;`

	err := api.db.Select(&userRow, query, idUser)
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

	const query = `SELECT * FROM prj_user WHERE nameuser = $1;`

	err := api.db.Select(&userRow, query, nameUser)
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

	const query = `SELECT * FROM prj_executor WHERE executorid = $1;`

	err := api.db.Select(&executerRow, query, executorId)
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

	const query = `SELECT * FROM prj_executor WHERE executorname = $1;`

	err := api.db.Select(&executerRow, query, executerName)
	if err != nil {
		log.Println("GetExecuterByName api.db.Select failed with an error: ", err.Error())
		return nil, err
	}
	if len(executerRow) == 1 {
		return &executerRow[0], err
	}

	return nil, err
}

// GetExecuterByNamePassword ...
func (api *API) GetExecuterByNamePassword(executerName, executerPassword string) (*apitypes.ExecutorRow, error) {
	executerRow := []apitypes.ExecutorRow{}

	const query = `SELECT * FROM prj_executor WHERE executorname = $1 AND executorpasword = $2;`

	err := api.db.Select(&executerRow, query, executerName, executerPassword)
	if err != nil {
		log.Println("GetExecuterByNamePassword api.db.Select failed with an error: ", err.Error())
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

	const query = `SELECT * FROM prj_order WHERE orderid = $1;`

	err := api.db.Select(&orderRow, query, orderId)
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

	const query = `SELECT * FROM prj_status WHERE statusid = $1;`

	err := api.db.Select(&statusRow, query, statusId)
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

	const query = `SELECT * FROM link_userorderexecutor WHERE linkid = $1;`

	err := api.db.Select(userOrderExecutorRow, query, userOrderExecutorId)
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

	const query = `SELECT * FROM link_userorderexecutor WHERE userid = $1 AND executorid = $2;`

	err := api.db.Select(userOrderExecutorRow, query, userId, executorid)
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
func (api *API) SetNewUser(ctx context.Context, user apitypes.UserRow) error {
	const query = `INSERT INTO prj_user(userid, nameuser, chatid)
	VALUES (:userid, :nameuser, :chatid)
	ON CONFLICT DO NOTHING
	;`

	if _, err := api.db.NamedExecContext(ctx, query, user); err != nil {
		return errors.Wrap(err, "can't add new user")
	}

	return nil
}

// SetNewExecutor ...
func (api *API) SetNewExecutor(ctx context.Context, executor apitypes.ExecutorRow) error {
	const query = `INSERT INTO prj_executor(executorid, executorname, executorpasword)
	VALUES (:executorid, :executorname, :executorpasword)
	ON CONFLICT DO NOTHING
	;`

	if _, err := api.db.NamedExecContext(ctx, query, executor); err != nil {
		return errors.Wrap(err, "can't add new executor")
	}

	return nil
}

// SetNewOrder ...
func (api *API) SetNewOrder(ctx context.Context, orderRow apitypes.OrderRow) error {
	const query = `INSERT INTO prj_order(orderid, orderdescription, statusid, orderstarttime)
	VALUES (:orderid, :orderdescription, :statusId, :orderstarttime)
	ON CONFLICT DO NOTHING
	;`

	if _, err := api.db.NamedExecContext(ctx, query, orderRow); err != nil {
		return errors.Wrap(err, "can't add new order")
	}

	return nil
}

// SetUserOrderExecutor ...
func (api *API) SetUserOrderExecutor(ctx context.Context, orderExecutor apitypes.UserOrderExecutorRow) error {

	const query = `INSERT INTO link_userorderexecutor (linkid, userid, orderid, executorid) 
	VALUES (:linkid, :userid, :orderid, :executorid)
	ON CONFLICT DO NOTHING
	;`

	if _, err := api.db.NamedExecContext(ctx, query, orderExecutor); err != nil {
		return errors.Wrap(err, "can't add new order")
	}

	return nil
}

// Block UPDATE

// UpdateUserName ...
func (api *API) UpdateUserName(ctx context.Context, userid, nameuser string) (*apitypes.UserRow, error) {
	var userRow *apitypes.UserRow
	var err error

	work := func(ctx context.Context, db TxContext) error {
		userRow, err := api.GetUserByID(userid)
		if err != nil {
			log.Println("UpdateUserName api.db.MustExec failed with an error: ", err.Error())
			return err
		}
		if userRow == nil {
			return err
		}

		query := `UPDATE prj_user SET nameuser=$1 WHERE userid=$2`

		if _, err := db.ExecContext(ctx, query, nameuser, userid); err != nil {
			return errors.Wrap(err, "UpdateUserName UPDATE prj_user failed: %s")
		}

		return nil
	}

	if err := RunInTransaction(ctx, api.db, work); err != nil {
		return nil, err
	}

	userRow, err = api.GetUserByID(userid)
	if err != nil {
		log.Println("UpdateUserName api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}

	return userRow, err
}

// UpdateExecuterName ...
func (api *API) UpdateExecuterName(ctx context.Context, executorId, executorName string) (*apitypes.ExecutorRow, error) {
	var executorRow *apitypes.ExecutorRow
	var err error

	work := func(ctx context.Context, db TxContext) error {
		executorRow, err := api.GetExecuterById(executorId)
		if err != nil {
			log.Println("UpdateExecuterName api.db.MustExec failed with an error: ", err.Error())
			return err
		}
		if executorRow == nil {
			return err
		}

		query := `UPDATE prj_executor SET executorname=$1 WHERE executorid=$2`
		if _, err := db.ExecContext(ctx, query, executorName, executorId); err != nil {
			return errors.Wrap(err, "UpdateExecuterName UPDATE prj_executor failed: %s")
		}

		return nil
	}

	if err := RunInTransaction(ctx, api.db, work); err != nil {
		return nil, err
	}

	executorRow, err = api.GetExecuterById(executorId)
	if err != nil {
		log.Println("UpdateExecuterName api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}

	return executorRow, err
}

// UpdateExecuterPassword ...
func (api *API) UpdateExecuterPassword(ctx context.Context, nameExecuter, oldPassword, newPassword string) (*apitypes.ExecutorRow, error) {

	var executorRow *apitypes.ExecutorRow
	var err error

	work := func(ctx context.Context, db TxContext) error {
		executerRow, err := api.GetExecuterByNamePassword(nameExecuter, oldPassword)
		if err != nil {
			log.Println("UpdateExecuterPassword api.db.MustExec failed with an error: ", err.Error())
			return err
		}
		if executerRow == nil {
			return err
		}

		query := `UPDATE prj_executor SET executorpasword=$1 WHERE executorid=$2`
		if _, err := db.ExecContext(ctx, query, newPassword, executerRow.ExecutorId); err != nil {
			return errors.Wrap(err, "UpdateExecuterPassword UPDATE prj_executor failed: %s")
		}

		return nil
	}
	if err := RunInTransaction(ctx, api.db, work); err != nil {
		return nil, err
	}

	executorRow, err = api.GetExecuterByNamePassword(nameExecuter, newPassword)

	if err != nil {
		log.Println("UpdateExecuterPassword api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}

	return executorRow, err
}

// UpdateOrderDescription ...
func (api *API) UpdateOrderDescription(orderId, orderdescription string) (*apitypes.OrderRow, error) {
	orderRow, err := api.GetOrderById(orderId)
	if err != nil {
		log.Println("UpdateOrderDescription api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}
	if orderRow == nil {
		return nil, err
	}

	tx := api.db.MustBegin()
	tx.MustExec(`UPDATE prj_order SET orderdescription=$1 WHERE orderid=$2`, orderdescription, orderId)
	tx.Commit()

	orderRow, err = api.GetOrderById(orderId)

	if err != nil {
		log.Println("UpdateOrderDescription api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}

	return orderRow, err
}

// UpdateOrderStatus ...
func (api *API) UpdateOrderStatus(orderId, statusId string) (*apitypes.OrderRow, error) {
	orderRow, err := api.GetOrderById(orderId)
	if err != nil {
		log.Println("UpdateOrderStatus api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}
	if orderRow == nil {
		return nil, err
	}

	tx := api.db.MustBegin()
	tx.MustExec(`UPDATE prj_order SET statusid=$1 WHERE orderid=$2`, statusId, orderId)
	tx.Commit()

	orderRow, err = api.GetOrderById(orderId)

	if err != nil {
		log.Println("UpdateOrderStatus api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}

	return orderRow, err
}

// UpdateOrderFinishStatus ...
func (api *API) UpdateOrderFinishStatus(orderId string) (*apitypes.OrderRow, error) {

	orderRow, err := api.GetOrderById(orderId)
	if err != nil {
		log.Println("UpdateOrderFinishStatus api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}
	if orderRow == nil {
		return nil, err
	}

	// получать финишный ID
	var finishId = "123"
	today := time.Now()
	orderFinishTime := today.Add(10 * time.Minute).Format("2006/1/2 15:04")

	tx := api.db.MustBegin()
	tx.MustExec(`UPDATE prj_order SET statusid=$1, orderstoptime=$2 WHERE orderid=$3`, finishId, orderFinishTime, orderId)
	tx.Commit()

	orderRow, err = api.GetOrderById(orderId)

	if err != nil {
		log.Println("UpdateOrderFinishStatus api.db.MustExec failed with an error: ", err.Error())
		return nil, err
	}

	return orderRow, err
}

//переписать методы Set и UPDATE на транзакции
