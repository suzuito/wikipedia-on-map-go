package main

import (
	"context"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/common-go/clogger"
	"github.com/suzuito/wikipedia-on-map-go/web"
	"github.com/suzuito/wikipedia-on-map-go/web/router"
)

type Application interface {
	web.ApplicationWeb
}

type ApplicationImpl struct {
	*web.ApplicationWebImpl
}

func NewApplicationImpl(ctx context.Context) (*ApplicationImpl, error) {
	appWeb, err := web.NewApplicationWebImpl(ctx)
	if err != nil {
		return nil, err
	}
	app := ApplicationImpl{
		ApplicationWebImpl: appWeb,
	}
	return &app, nil
}

func (a *ApplicationImpl) Logger(ctx context.Context) clogger.Logger {
	return &clogger.LoggerPrint{}
}

func main() {
	ctx := context.Background()
	app, err := NewApplicationImpl(ctx)
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
	s2geo(app, root)
	root.Run()
}

func s2geo(app Application, root *gin.Engine) {
	s2g := root.Group("s2")
	s2g.GET("face/:faceID", router.GetCellFromFace(app))
}

func geo(app Application, root *gin.Engine) {
	groupGeo := root.Group("geo")
	groupGeo.GET("cells/children", router.GetGeoCellsChildren(app))
	groupGeo.GET("cells/convex", router.GetGeoCellsConvex(app))
	groupGeo.GET("caps", router.GetGeoCap(app))
	groupGeo.GET("locations", router.GetGeoLocations(app))
}

func usecors(app Application, root *gin.Engine) {
	root.Use(cors.New(cors.Config{
		AllowOrigins:     app.AllowedOrigins(),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
}
