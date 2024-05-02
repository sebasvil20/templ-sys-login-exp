package users

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/sebasvil20/templ-sys-login-exp/models"
	"github.com/sebasvil20/templ-sys-login-exp/utils"
	"github.com/sebasvil20/templ-sys-login-exp/views/pages"
)

type IUserController interface {
	ListUserPage(w http.ResponseWriter, r *http.Request)
	AccountsPage(w http.ResponseWriter, r *http.Request)
	LoginAPI(w http.ResponseWriter, r *http.Request)
	SigninAPI(w http.ResponseWriter, r *http.Request)
}

type UserController struct {
	UserHandlerInyectado IUser
	decoder              *schema.Decoder
	store                *sessions.CookieStore
}

func InitUserController(userHandlerAInyectar IUser, decoderAInyectar *schema.Decoder, storeAInyectar *sessions.CookieStore) UserController {
	return UserController{
		UserHandlerInyectado: userHandlerAInyectar,
		decoder:              decoderAInyectar,
		store:                storeAInyectar,
		// Aca el resto de propiedades.....
	}
}

func (userControllerArg UserController) ListUserPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	userList, _ := userControllerArg.UserHandlerInyectado.GetUsers()
	pages.ListUser(userList).Render(r.Context(), w)
}

func (userControllerArg UserController) AccountsPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	pages.Accounts().Render(r.Context(), w)
}

func (userControllerArg UserController) LoginAPI(w http.ResponseWriter, r *http.Request) {
	validate, _ := r.Context().Value("validator").(*validator.Validate)

	reqUser := models.UserCredentials{}
	err := r.ParseMultipartForm(4096)
	if err != nil {
		utils.HandleReturnWithStatusCode(w, 400, map[string]string{"error": "Bad request body"})
		return
	}

	err = userControllerArg.decoder.Decode(&reqUser, r.PostForm)
	if err != nil {
		utils.HandleReturnWithStatusCode(w, 400, map[string]string{"error": "Bad request body"})
		return
	}

	err = validate.Struct(&reqUser)
	if err != nil {
		utils.HandleReturnWithStatusCode(w, 400, map[string]string{"error": fmt.Sprintf("Validation error: %s", err.Error())})
		return
	}

	err = userControllerArg.UserHandlerInyectado.Authenticate(reqUser)
	if err != nil {
		utils.HandleReturnWithStatusCode(w, 401, map[string]string{"error": "Invalid credentials"})
		return
	}
	session, _ := userControllerArg.store.Get(r, "session")
	session.Values["authenticated"] = true
	err = session.Save(r, w)
	if err != nil {
		utils.HandleReturnWithStatusCode(w, 500, map[string]string{"error": "Error saving session"})
		return
	}
	utils.HandleReturnWithStatusCode(w, 200, map[string]any{"data": nil})
}

func (userControllerArg UserController) SigninAPI(w http.ResponseWriter, r *http.Request) {
	validate, _ := r.Context().Value("validator").(*validator.Validate)
	reqUser := models.User{}

	err := r.ParseMultipartForm(4096)
	if err != nil {
		utils.HandleReturnWithStatusCode(w, 400, map[string]string{"error": "Bad request body"})
		return
	}

	err = userControllerArg.decoder.Decode(&reqUser, r.PostForm)
	if err != nil {
		utils.HandleReturnWithStatusCode(w, 400, map[string]string{"error": "Bad request body"})
		return
	}

	err = validate.Struct(&reqUser)
	if err != nil {
		utils.HandleReturnWithStatusCode(w, 400, map[string]string{"error": fmt.Sprintf("Validation error: %s", err.Error())})
		return
	}

	err = userControllerArg.UserHandlerInyectado.AddUser(reqUser)
	if err != nil {
		utils.HandleReturnWithStatusCode(w, 500, map[string]string{"error": "Error adding user " + err.Error()})
		return
	}

	session, _ := userControllerArg.store.Get(r, "session")
	session.Values["authenticated"] = true
	err = session.Save(r, w)
	if err != nil {
		utils.HandleReturnWithStatusCode(w, 500, map[string]string{"error": "Error saving session"})
		return
	}
	utils.HandleReturnWithStatusCode(w, 201, map[string]any{"data": nil})
}
