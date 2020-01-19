package main

import (
	"context"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/wikipedia-on-map-go/application"
	"github.com/suzuito/wikipedia-on-map-go/gcp"
	"github.com/suzuito/wikipedia-on-map-go/web/router"
)

type ApplicationWeb struct {
	*gcp.ApplicationGCP
}

func NewApplicationWeb(ctx context.Context) (*ApplicationWeb, error) {
	appGcp, err := gcp.NewApplicationGCP(ctx)
	if err != nil {
		return nil, err
	}
	app := ApplicationWeb{
		ApplicationGCP: appGcp,
	}
	return &app, nil
}

// AllowedOrigins ...
func (a *ApplicationWeb) AllowedOrigins() []string {
	return strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
}

func main() {
	ctx := context.Background()
	app, err := NewApplicationWeb(ctx)
	if err != nil {
		panic(err)
	}
	root := gin.New()
	if os.Getenv("GAE_APPLICATION") == "" {
		root.Use(gin.Logger())
	}
	root.Use(gin.Recovery())
	usecors(app, root)
	geo(app, root)
	root.Run()
}

func geo(app application.Application, root *gin.Engine) {
	groupGeo := root.Group("geo")
	groupGeo.GET("cells/children", router.GetGeoCellsChildren(app))
	groupGeo.GET("cells/convex", router.GetGeoCellsConvex(app))
}

func usecors(app *ApplicationWeb, root *gin.Engine) {
	root.Use(cors.New(cors.Config{
		AllowOrigins:     app.AllowedOrigins(),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
}
