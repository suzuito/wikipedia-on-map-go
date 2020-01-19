package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/suzuito/wikipedia-on-map-go/entity/model"
	"github.com/suzuito/wikipedia-on-map-go/gcp"
	"github.com/suzuito/wikipedia-on-map-go/geo"
)

var splitter = regexp.MustCompile(`<.*?>`)
var extractorCoord = regexp.MustCompile(`point> "(.*?)" <`)
var extractorName = regexp.MustCompile(`ja\.dbpedia\.org/resource/(.*?)>`)

func strip(s string) string {
	return strings.Trim(strings.TrimLeft(s, "<"), ">")
}

func extract(s string, r *regexp.Regexp) string {
	results := r.FindStringSubmatch(s)
	if len(results) <= 1 {
		return ""
	}
	return results[1]
}

func main() {
	ctx := context.Background()
	app, err := gcp.NewApplicationGCP(ctx)
	if err != nil {
		panic(err)
	}
	cli, err := app.DBClient(ctx)
	if err != nil {
		panic(err)
	}
	glocs := []*model.GeoLocation{}
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()
		if strings.HasPrefix(text, "#") {
			continue
		}
		fieldCoord := extract(text, extractorCoord)
		if fieldCoord == "" {
			continue
		}
		name := extract(text, extractorName)
		if name == "" {
			continue
		}
		var lat = float64(0)
		var lng = float64(0)
		fmt.Sscanf(fieldCoord, "%f %f", &lat, &lng)
		cell := geo.GetCell(lat, lng, app.ApplicationBase.IndexLevel())
		gloc := model.NewGeoLocation(
			name,
			lat,
			lng,
			cell.ID().ToToken(),
		)
		glocs = append(glocs, gloc)
	}
	if err := cli.SetGeoLocations(ctx, app.ApplicationBase.IndexLevel(), &glocs); err != nil {
		panic(err)
	}
}
