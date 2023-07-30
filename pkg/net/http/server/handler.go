package server

import (
	"net/http"
	"strconv"
)

type Handler struct{}

func (c *Handler) Param(context *Context, key string) string {
	return context.Param(key)
}

func (c *Handler) Query(context *Context, key string) (string, bool) {
	return context.GetQuery(key)
}

func (c *Handler) QueryBool(context *Context, key string) (bool, bool) {
	v, ok := context.GetQuery(key)
	if ok {
		return v == "true", true
	}
	return false, false
}

func (c *Handler) QueryInt(context *Context, key string) (int, bool) {
	v, ok := context.GetQuery(key)
	if ok {
		if vInt, err := strconv.Atoi(v); err == nil {
			return vInt, true
		}
	}
	return 0, false
}

func (c *Handler) QueryInt64(context *Context, key string) (int64, bool) {
	v, ok := context.GetQuery(key)
	if ok {
		if vInt64, err := strconv.ParseInt(v, 10, 64); err == nil {
			return vInt64, true
		}
	}
	return 0, false
}

func (c *Handler) QueryFloat(context *Context, key string) (float64, bool) {
	v, ok := context.GetQuery(key)
	if ok {
		if vFloat, err := strconv.ParseFloat(v, 64); err == nil {
			return vFloat, true
		}
	}
	return 0, false
}

func (c *Handler) BindJSON(context *Context, obj interface{}) error {
	return context.BindJSON(obj)
}

func (c *Handler) JSON(context *Context, data interface{}, codeAndMsg ...interface{}) {
	code, message := 0, "ok"
	if len(codeAndMsg) > 0 {
		code = codeAndMsg[0].(int)
		if len(codeAndMsg) > 1 {
			message = codeAndMsg[1].(string)
		}
	}
	context.JSON(http.StatusOK, map[string]interface{}{"code": code, "message": message, "data": data})
}
