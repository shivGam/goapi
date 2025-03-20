package middleware
import (
	"errors"
	"net/http"
	"github.com/shivGam/goapi/api"
	"github.com/shivGam/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var unAuthorizedError = errors.New("Invalid Tokens or Username")

func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w, unAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails = (*database).GetUserLoginDetails(username)

		if(loginDetails == nil || token != (*loginDetails).AuthToken){
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w,unAuthorizedError)
			return
		}
		next.ServeHTTP(w,r)
	})
}