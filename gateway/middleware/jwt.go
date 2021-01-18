package middleware

import (

	//"github.com/davecgh/go-spew/spew"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/ast"
)

type postParams struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func systemError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code": -1,
		"msg":  "system error",
	})
	ctx.Abort()
	return
}

/*
获取查询的深度
*/
func getQueryFieldDeep(set ast.SelectionSet, deep int) int {
	if set == nil {
		return deep
	}
	deep++
	max := 0
	for _, v := range set {
		if tmp, ok := v.(*ast.Field); ok {
			d := getQueryFieldDeep(tmp.SelectionSet, deep)
			if d > max {
				max = d
			}
		}
	}
	return max
}

/*
	指定的请求路径是否在queryName里
*/
func checkQuery(c *gin.Context) bool {
	queryName := c.GetStringSlice("queryName")
	for _, v := range queryName {
		if v == "posts" || v == "signIn" || v == "signUp" {
			return true
		}
	}
	return false
}
