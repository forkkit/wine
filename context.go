package wine

import (
	"github.com/justintan/gox/types"
	"html/template"
	"net/http"
)

type Context interface {
	HTTPRequest() *http.Request
	Params() types.M
	Header() http.Header

	Next()
	Responded() bool
	ResponseHeader() http.Header
	Status(status int)
	JSON(obj interface{})
	File(filePath string)
	HTML(htmlText string)
	Text(text string)
	TemplateHTML(templateName string, params types.M)
	ServeHTTP(h http.Handler)

	Set(key string, value interface{})
	Get(key string) interface{}

	Rebuild(http.ResponseWriter, *http.Request, []*template.Template, []Handler)
}
