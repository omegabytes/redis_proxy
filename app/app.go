package app

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"os"
	"redis_proxy/app/db"
	"redis_proxy/app/models"
	"redis_proxy/app/routes"
)

var (
	ADDR     = os.Getenv("REDIS_URL")
	APP_ADDR = ":" + os.Getenv("PORT")
	DEMO     = os.Getenv("DEMO")
)

type App struct {
	Router *chi.Mux
}

func (a *App) Initialize() {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/post", routes.Routes())
	})
	a.Router = router
}

func (a *App) Run() {
	WalkRoutes(a.Router)

	log.Println("[DEBUG] app address: http://localhost", APP_ADDR)

	if DEMO != "" {
		log.Println("Thanks for checking out my code! Initializing data store with some values...")
		InitDBVals()
		log.Println("Done!")
	}

	log.Fatal(http.ListenAndServe(APP_ADDR, a.Router))
}

func WalkRoutes(router *chi.Mux) {
	walkFunc := func(method string, route string, hander http.Handler, middlewares ...func(handler http.Handler) http.Handler) error {
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf(err.Error())
	}
}

func InitDBVals() {
	example := []*models.Post{
		{Title: "first", Body: "one"},
		{Title: "second", Body: "two"},
		{Title: "third", Body: "three"},
		{Title: "fourth", Body: "four"},
		{Title: "fifth", Body: "five"},
		{Title: "sixth", Body: "six"},
		{Title: "seventh", Body: "seven"},
		{Title: "eighth", Body: "eight"},
		{Title: "nineth", Body: "nine"},
		{Title: "tenth", Body: "ten"},
	}

	for _, item := range example {
		db.CreatePost(*item)
	}
}
