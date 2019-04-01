package controller

import (
	"github.com/Pluslab/gaia/server/domain"
	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
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
func (controller *SessionController) Login(c *gin.Context) {
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
	session := sessions.Default(c)
	session.Set("alive", true)
	session.Set("UserID", user.ID)
	session.Set("UserName", user.Name)
	session.Save()
	c.JSON(200, session)
}

// Logout is a function for logging out
func (controller *SessionController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(200, session)
}
