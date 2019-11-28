package example

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/dxvgef/tsing"
)

func TestGET(t *testing.T) {
	log.SetFlags(log.Lshortfile)

	app := tsing.New(&tsing.Config{
		RedirectTrailingSlash: false,
		HandleOPTIONS:         false,
		Recover:               false,
		ErrorEvent:            false,
		NotFoundEvent:         false,
		MethodNotAllowedEvent: false,
		Trigger:               false,
		Trace:                 false,
		ShortPath:             false,
		EventHandler: func(event *tsing.Event) {
			log.Println(event.Message)
			log.Println(event.Trigger)
			log.Println(event.Trace)
		},
	})
	app.Config.EventHandler = func(event *tsing.Event) {
		log.Println(event.Message)
	}
	app.Router.GET("/", func(ctx *tsing.Context) error {
		t.Log(ctx.QueryValue("id"))
		return nil
	})

	r, err := http.NewRequest("GET", "/?id=abc", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	app.ServeHTTP(httptest.NewRecorder(), r)
}

func TestPOST(t *testing.T) {
	log.SetFlags(log.Lshortfile)

	app := tsing.New(&tsing.Config{
		RedirectTrailingSlash: false,
		HandleOPTIONS:         false,
		Recover:               false,
		ErrorEvent:            false,
		NotFoundEvent:         false,
		MethodNotAllowedEvent: false,
		Trigger:               false,
		Trace:                 false,
		ShortPath:             false,
		EventHandler: func(event *tsing.Event) {
			log.Println(event.Message)
			log.Println(event.Trigger)
			log.Println(event.Trace)
		},
	})
	app.Config.EventHandler = func(event *tsing.Event) {
		log.Println(event.Message)
	}
	app.Router.POST("/", func(ctx *tsing.Context) error {
		t.Log(ctx.PostValue("id"))
		return nil
	})

	v := url.Values{}
	v.Add("id", "abc")
	r, err := http.NewRequest("POST", "/", strings.NewReader(v.Encode()))
	if err != nil {
		t.Error(err.Error())
		return
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	app.ServeHTTP(httptest.NewRecorder(), r)
}

func TestRoute(t *testing.T) {
	log.SetFlags(log.Lshortfile)

	app := tsing.New(&tsing.Config{
		RedirectTrailingSlash: false,
		HandleOPTIONS:         false,
		Recover:               false,
		ErrorEvent:            false,
		NotFoundEvent:         false,
		MethodNotAllowedEvent: false,
		Trigger:               false,
		Trace:                 false,
		ShortPath:             false,
		EventHandler: func(event *tsing.Event) {
			log.Println(event.Message)
			log.Println(event.Trigger)
			log.Println(event.Trace)
		},
	})
	app.Config.EventHandler = func(event *tsing.Event) {
		log.Println(event.Message)
	}
	app.Router.GET("/:classID/:id", func(ctx *tsing.Context) error {
		t.Log(ctx.ParamValue("classID"))
		t.Log(ctx.ParamValue("id"))
		return nil
	})

	r, err := http.NewRequest("GET", "/1/2", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	app.ServeHTTP(httptest.NewRecorder(), r)
}