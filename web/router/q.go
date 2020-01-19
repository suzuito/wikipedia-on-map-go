package router

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// QueryFloat64 ...
func QueryFloat64(ctx *gin.Context, key string, dflt float64) float64 {
	v := ctx.Query(key)
	if v == "" {
		return dflt
	}
	r, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return dflt
	}
	return r
}

// QueryInt64 ...
func QueryInt64(ctx *gin.Context, key string, dflt int64) int64 {
	v := ctx.Query(key)
	if v == "" {
		return dflt
	}
	r, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return dflt
	}
	return r
}

// QueryUint64 ...
func QueryUint64(ctx *gin.Context, key string, dflt uint64) uint64 {
	v := ctx.Query(key)
	if v == "" {
		return dflt
	}
	r, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return dflt
	}
	return r
}