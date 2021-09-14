package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"

	"gitlab.com/banter-bus/banter-bus-management-api/internal/games"
)

func GameRoutes(env *games.GameAPI, grp *gin.RouterGroup) {
	grp.POST("", checkJWT(env.Conf), tonic.Handler(env.AddGame, http.StatusCreated))
	grp.GET("", tonic.Handler(env.GetGames, http.StatusOK))
	grp.GET("/:game_name", tonic.Handler(env.GetGame, http.StatusOK))
	grp.DELETE("/:game_name", checkJWT(env.Conf), tonic.Handler(env.RemoveGame, http.StatusOK))
	grp.PUT("/:game_name/enable", checkJWT(env.Conf), tonic.Handler(env.EnableGame, http.StatusOK))
	grp.PUT("/:game_name/disable", checkJWT(env.Conf), tonic.Handler(env.DisableGame, http.StatusOK))
}
