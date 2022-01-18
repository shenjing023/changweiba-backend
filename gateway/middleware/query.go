package middleware

import (
	"bytes"
	"encoding/json"
	"gateway/common"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/shenjing023/llog"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/parser"
)

const queryNameKey = "queryName"

// QueryDeepMiddleware 检测请求查询字段深度的中间件
func QueryDeepMiddleware(queryDeep int) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			return
		}
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Errorf("read request body error: %+v", err)
			systemError(c)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // 关键点,不能去掉

		//解析query
		var param postParams
		err = json.Unmarshal(body, &param)
		if err != nil {
			log.Errorf("unmarshal post param error:%+v, body: %s", err, string(body))
			systemError(c)
		}

		doc, err_ := parser.ParseQuery(&ast.Source{Input: param.Query})
		//spew.Dump(err)
		if err_ != nil {
			log.Errorf("parse query error: %+v", err_)
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
						c.JSON(http.StatusBadRequest, []gqlError{{
							Message: "请求字段深度超出限制",
							Extensions: map[string]interface{}{
								"code": common.InvalidArgument,
							},
							Path: []string{tmp.Name},
						}})
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
		c.Set(queryNameKey, queryName)
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

type postParams struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func systemError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, []gqlError{{
		Message: "service system error",
		Extensions: map[string]interface{}{
			"code": common.Internal,
		},
	}})
	ctx.Abort()
}
