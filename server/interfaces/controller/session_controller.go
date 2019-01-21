package controller

import (
	"github.com/Pluslab/gaia/server/domain"
	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
	"golang.org/x/crypto/bcrypt"
)

// SessionController model
type SessionController struct {
	Interactor usecase.SessionInteractor
}

// NewSessionController return the SessionController
func NewSessionController(SQLHandler database.SQLHandler) *SessionController {
	return &SessionController{
		Interactor: usecase.SessionInteractor{
			SessionRepository: &database.SessionRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Login is a function for logging in
func (controller *SessionController) Login(c Context) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Login(u.Email)
	if err != nil {
		c.JSON(500, err)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		c.JSON(500, err)
		return
	}

	// TODO: session info
	c.JSON(200, user)
}
