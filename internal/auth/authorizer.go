package auth

import (
	"fmt"

	casbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(_model, _policy string) *Authorizer {
	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")
	enforcer, _ := casbin.NewEnforcer(m)
	_, _ = enforcer.AddPermissionForUser("root", "*", "produce")
	_, _ = enforcer.AddPermissionForUser("root", "*", "consume")
	return &Authorizer{
		enforcer: enforcer,
	}
}

type Authorizer struct {
	enforcer *casbin.Enforcer
}

func (a *Authorizer) Authorize(subject, object, action string) error {
	if res, _ := a.enforcer.Enforce(subject, object, action); !res {
		msg := fmt.Sprintf(
			"%s not permitted to %s to %s.",
			subject,
			action,
			object,
		)
		st := status.New(codes.PermissionDenied, msg)
		return st.Err()
	}
	return nil
}
