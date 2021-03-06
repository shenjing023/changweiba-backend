package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/shenjing023/llog"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/parser"
)

// QueryDeepMiddleware 检测请求查询字段深度的中间件
func QueryDeepMiddleware(queryDeep int) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			return
		}
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Error("read request body error:", err.Error())
			systemError(c)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // 关键点,不能去掉

		//解析query
		var param postParams
		err = json.Unmarshal(body, &param)
		if err != nil {
			log.Error(fmt.Sprintf("unmarshal post param error:%s, body: %s", err.Error(), string(body)))
			systemError(c)
		}

		//陷阱，不能是doc,err:= 目前还不知原因
		doc, err_ := parser.ParseQuery(&ast.Source{Input: param.Query})
		//spew.Dump(err)
		if err_ != nil {
			log.Error("parse query error: ", err_)
			systemError(c)
		}
		var queryName []string //存储查询的接口名称
		ops := doc.Operations
		for _, v := range ops {
			for _, k := range v.SelectionSet {
				if tmp, ok := k.(*ast.Field); ok {
					//检查查询的字段深度,待优化，还有directive
					deep := getQueryFieldDeep(tmp.SelectionSet, 0)
					if deep > queryDeep {
						c.JSON(http.StatusBadRequest, gin.H{
							"status": -1,
							"msg":    "请求字段深度超出限制",
						})
						c.Abort()
						return
					}
					queryName = append(queryName, tmp.Name)
				} else {
					log.Error("selection change to ast.Field error")
					systemError(c)
				}
			}
		}
		c.Set("queryName", queryName)
		c.Next()
	}
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
