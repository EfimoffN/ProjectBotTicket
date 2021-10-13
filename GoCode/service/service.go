package service

import (
	"context"
	"strconv"

	"projectbotticket/types"
	"projectbotticket/types/apitypes"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Storage ...
type Storage interface {
	//GetLastCommandByUserName(usename string) (*apitypes.LastUserCommand, error)
	//CheckingBotTicketWork(userN string) (bool, error)
	//StopBotTicketWork(userN string) error
	//AddNewUser(username string) error

	GetTestValue(val int) int
	GetUserByID(idUser string) (*apitypes.UserRow, error)
	GetUserByName(nameUser string) (*apitypes.UserRow, error)
	GetExecuterById(executorId string) (*apitypes.ExecutorRow, error)
	GetExecuterByName(executerName string) (*apitypes.ExecutorRow, error)
	GetExecuterByNamePassword(executerName, executerPassword string) (*apitypes.ExecutorRow, error)
	GetOrderById(orderId string) (*apitypes.OrderRow, error)
	GetStatusById(statusId string) (*apitypes.StatusRow, error)
	GetUserOrderExecutorById(userOrderExecutorId string) (*apitypes.UserOrderExecutorRow, error)
	GetUserOrderExecutorByUserIdExecutorId(userId string, executorid string) (*apitypes.UserOrderExecutorRow, error)
	SetNewUser(ctx context.Context, user apitypes.UserRow) error
	SetNewExecutor(ctx context.Context, executor apitypes.ExecutorRow) error
	SetNewOrder(ctx context.Context, orderRow apitypes.OrderRow) error
	SetUserOrderExecutor(ctx context.Context, orderExecutor apitypes.UserOrderExecutorRow) error
	UpdateUserName(ctx context.Context, userid, nameuser string) error
	UpdateExecuterName(ctx context.Context, executorId, executorName string) error
	UpdateExecuterPassword(ctx context.Context, nameExecuter, oldPassword, newPassword string) error
	UpdateOrderDescription(ctx context.Context, orderId, orderdescription string) error
	UpdateOrderStatus(ctx context.Context, orderId, statusId string) error
	UpdateOrderFinishStatus(ctx context.Context, orderId, statusId string) error
}

// BotSvc ...
type BotSvc struct {
	storage     Storage
	commandsBot types.Commands
}

// NewBotSvc ...
func NewBotSvc(s Storage, commandsBot types.Commands) *BotSvc {
	return &BotSvc{
		storage:     s,
		commandsBot: commandsBot,
	}
}

func (b *BotSvc) ProcessingCommands(message *tgbotapi.Message, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

}

func (b *BotSvc) checkingBotWork(user *tgbotapi.User) (bool, error) {
	userIn, err := b.storage.GetUserByID(strconv.Itoa(user.ID))
}
