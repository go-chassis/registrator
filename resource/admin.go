package resource

import (
	"encoding/json"
	"github.com/go-chassis/go-chassis/pkg/string"
	"github.com/go-chassis/go-chassis/server/restful"
	"github.com/go-mesh/registrator/reg"
	"net/http"
)

type Admin struct {
}

func (r *Admin) GetInstances(ctx *restful.Context) {
	m := reg.GetInstances()
	b, err := json.MarshalIndent(m, "", "   ")
	if err != nil {
		ctx.Write(stringutil.Str2bytes(err.Error()))
		ctx.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.Write(b)
	ctx.WriteHeader(http.StatusOK)
}
func (r *Admin) GetInstancesByName(ctx *restful.Context) {
	m := reg.GetInstances()
	b, err := json.MarshalIndent(m[ctx.ReadPathParameter("service")], "", "   ")
	if err != nil {
		ctx.Write(stringutil.Str2bytes(err.Error()))
		ctx.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.Write(b)
	ctx.WriteHeader(http.StatusOK)

}

//URLPatterns helps to respond for corresponding API calls
func (r *Admin) URLPatterns() []restful.Route {
	return []restful.Route{
		{Method: http.MethodGet, Path: "/instances/{service}", ResourceFuncName: "GetInstancesByName",
			Returns: []*restful.Returns{{Code: 200}}},
		{Method: http.MethodGet, Path: "/instances/", ResourceFuncName: "GetInstances",
			Returns: []*restful.Returns{{Code: 200}}},
	}
}
