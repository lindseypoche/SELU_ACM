package web

import (
	"os"

	router "github.com/lindseypoche/SELU_ACM/api/internal/adapters/http"
	"github.com/lindseypoche/SELU_ACM/api/internal/adapters/storage"
	"github.com/lindseypoche/SELU_ACM/api/internal/domain"
)

var (
	// messages/blogs/events
	messageRepository domain.MessageRepository = storage.NewMongoRepo("mongodb://127.0.0.1:27017", "acm", 5)
	messageService    domain.MessageService    = domain.NewMessageService(messageRepository)
	messageHandler    MessageController        = NewMessageController(messageService)

	// // users
	// userRepository domain.UserRepository = storage.NewMySQLRepository("127.0.0.1:3306", "users_db")
	// userService    domain.UserService    = domain.NewUserService(userRepository)
	// userHandler    UserController        = NewUserController(userService)

	// router
	httpRouter router.Context = router.NewGinRouter()
)

func getPort() string {
	if os.Getenv("PORT") != "" {
		return os.Getenv("PORT")
	}
	return "8080"
}

// StartApplication ...
func StartApplication() {

	mapUrls()

	httpRouter.SERVE(":" + getPort())
}
