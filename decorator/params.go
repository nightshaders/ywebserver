package decorator

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)


type Params interface {
	Response() http.ResponseWriter
	Request() *http.Request
	Session() *sessions.Session
	SessionKey() string
	SetSession(*sessions.Session, string)
	Render(name, tp string, data interface{}) error
	RedirectToLogin()
	RedirectToLandingPage()
	SessionString(key string) string
}

type BaseParams struct {
	Params
	response http.ResponseWriter
	request *http.Request
	session *sessions.Session
	sessionKey string
}

func NewParams(res http.ResponseWriter, req *http.Request) Params {
	return &BaseParams{
		response: res,
		request: req,
	}
}
func (p *BaseParams) Response() http.ResponseWriter {
	return p.response
}
func (p *BaseParams) Request() *http.Request {
	return p.request
}
func (p *BaseParams) Session() *sessions.Session {
	return p.session
}
func (p *BaseParams) SetSession(s *sessions.Session, key string) {
	p.session = s
	p.sessionKey = key
}

func (p *BaseParams) SessionString(key string) string {
	v, ok := p.Session().Values[key]
	if ok {
		s, ok := v.(string)
		if ok {
			return s
		}
	}
	return ""
}

func (p *BaseParams) Render(name, tp string, data interface{}) error {

	tmpl, err := template.New(name).Parse(tp)

	if err != nil {
		return fmt.Errorf("Could not create template for name: %s", name)
	}

	err = tmpl.Execute(p.Response(), data)

	if err != nil {
		return fmt.Errorf("Could not execute template %s for given data", name)
	}

	return nil
}

func (p *BaseParams) RedirectToLogin() {
	http.Redirect(p.Response(), p.Request(), "/index.html", http.StatusTemporaryRedirect)
}

func (p *BaseParams) RedirectToLandingPage() {
	http.Redirect(p.Response(), p.Request(), "/batman", http.StatusTemporaryRedirect)
}

