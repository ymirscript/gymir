package gymir

import (
	b64 "encoding/base64"
	"encoding/json"
)

type SyntaxNode struct{}

type QueryParameterType string

const (
	Any      QueryParameterType = "any"
	String   QueryParameterType = "string"
	Int      QueryParameterType = "int"
	Float    QueryParameterType = "float"
	Bool     QueryParameterType = "bool"
	Date     QueryParameterType = "date"
	DateTime QueryParameterType = "datetime"
	Time     QueryParameterType = "time"
)

type QueryParameterNode struct {
	SyntaxNode
	Name string             `json:"name"`
	Type QueryParameterType `json:"type"`
}

type PathNode struct {
	SyntaxNode
	Path            string               `json:"path"`
	Alias           *string              `json:"alias"`
	QueryParameters []QueryParameterNode `json:"queryParameters"`
}

type RouterNode struct {
	SyntaxNode
	Path         PathNode                `json:"path"`
	Routers      []RouterNode            `json:"routers"`
	Routes       []RouteNode             `json:"routes"`
	Header       *MiddlewareOptions      `json:"header"`
	Body         *MiddlewareOptions      `json:"body"`
	Authenticate *AuthenticateClauseNode `json:"authenticate"`
}

type ScriptFileNode struct {
	RouterNode
}

type ProjectNode struct {
	ScriptFileNode
	Target      string                   `json:"target"`
	AuthBlocks  map[string]AuthBlockNode `json:"authBlocks"`
	Middlewares []MiddlewareNode         `json:"middlewares"`
}

type RouteNode struct {
	SyntaxNode
	Path         PathNode                `json:"path"`
	Method       Method                  `json:"method"`
	Header       *MiddlewareOptions      `json:"header"`
	Body         *MiddlewareOptions      `json:"body"`
	Authenticate *AuthenticateClauseNode `json:"authenticate"`
	Description  *string                 `json:"description"`
}

type MiddlewareOptions map[string]interface{}

type AuthBlockNode struct {
	SyntaxNode
	Alias                 *string 			`json:"alias"`
	Type                  AuthType 			`json:"type"`
	Source                string  			`json:"source"`
	Field                 string 			`json:"field"`
	IsDefaultAccessPublic *bool  			`json:"isDefaultAccessPublic"`
	IsAuthorizationInUse  bool 				`json:"isAuthorizationInUse"`
	options 			  MiddlewareOptions `json:"options"`

}

type AuthenticateClauseNode struct {
	SyntaxNode
	AuthBlock     string    `json:"authBlock"`
	Authorization *[]string `json:"authorization"`
}

type MiddlewareNode struct {
	SyntaxNode
	Name    string            `json:"name"`
	Options MiddlewareOptions `json:"options"`
}

type GlobalVariable struct {
	Name string   `json:"name"`
	Path []string `json:"path"`
}

type Method string

const (
	Get     Method = "GET"
	Post    Method = "POST"
	Put     Method = "PUT"
	Delete  Method = "DELETE"
	Patch   Method = "PATCH"
	Options Method = "OPTIONS"
	Head    Method = "HEAD"
)

type AuthType string

const (
	APIKey AuthType = "API-Key"
	Bearer AuthType = "Bearer"
)

type TargetConfig map[string]interface{}

type YmirData struct {
	Project ProjectNode  `json:"project"`
	Config  TargetConfig `json:"config"`
	Output  string       `json:"output"`
}

func GetYmirData(arg string) (*YmirData, error) {
	decoded, err := b64.StdEncoding.DecodeString(arg)
	if err != nil {
		return nil, err
	}

	var ymirData YmirData
	err = json.Unmarshal(decoded, &ymirData)

	if err != nil {
		return nil, err
	}

	return &ymirData, nil
}
