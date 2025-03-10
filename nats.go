package errors

import (
	"strings"
)

func Nats(v string) (err E) {
	switch {
	case strings.Contains(v, "account exists"):
		err = Exist(v)
	case strings.Contains(v, "account expired"):
		err = Expired(v)
	case strings.Contains(v, "account missing"):
		err = Empty(v)
	case strings.Contains(v, "account resolver missing"):
		err = Empty(v)
	case strings.Contains(v, "account resolver no new claims"):
		err = Try(v)
	case strings.Contains(v, "account resolver update too soon"):
		err = Try(v)
	case strings.Contains(v, "account validation failed"):
		err = Access(v)
	case strings.Contains(v, "attempted to connect to gateway port"):
		err = Connection(v)
	case strings.Contains(v, "attempted to connect to leaf node port"):
		err = Connection(v)
	case strings.Contains(v, "attempted to connect to route port"):
		err = Connection(v)
	case strings.Contains(v, "attempted to connect to wrong port"):
		err = Connection(v)
	case strings.Contains(v, "authentication error"):
		err = Access(v)
	case strings.Contains(v, "authentication expired"):
		err = Access(v)
	case strings.Contains(v, "authentication timeout"):
		err = Timeout(v)
	case strings.Contains(v, "bad account"):
		err = Invalid(v)
	case strings.Contains(v, "bad message header detected"):
		err = Invalid(v)
	case strings.Contains(v, "bad qualifier"):
		err = Invalid(v)
	case strings.Contains(v, "bad sampling percentage, should be 1-100"):
		err = Invalid(v)
	case strings.Contains(v, "bad service response type"):
		err = Invalid(v)
	case strings.Contains(v, "certificate not pinned"):
		err = Connection(v)
	case strings.Contains(v, "cluster name cannot contain spaces"):
		err = Invalid(v)
	case strings.Contains(v, "cluster name conflicts between cluster and gateway definitions"):
		err = Invalid(v)
	case strings.Contains(v, "cluster name from remote server conflicts"):
		err = Invalid(v)
	case strings.Contains(v, "connection closed"):
		err = Connection(v)
	case strings.Contains(v, "credentials have been revoked"):
		err = Access(v)
	case strings.Contains(v, "duplicate server name"):
		err = Exist(v)
	case strings.Contains(v, "function argument is invalid or in the wrong format"):
		err = Invalid(v)
	case strings.Contains(v, "gateway name cannot contain spaces"):
		err = Invalid(v)
	case strings.Contains(v, "import forms a cycle"):
		err = Invalid(v)
	case strings.Contains(v, "invalid client protocol"):
		err = Invalid(v)
	case strings.Contains(v, "invalid mapping destination"):
		err = Invalid(v)
	case strings.Contains(v, "invalid publish subject"):
		err = Invalid(v)
	case strings.Contains(v, "invalid subject"):
		err = Invalid(v)
	case strings.Contains(v, "invalid transform"):
		err = Invalid(v)
	case strings.Contains(v, "leafnode loop detected"):
		err = Invalid(v)
	case strings.Contains(v, "leafnodes disabled"):
		err = Invalid(v)
	case strings.Contains(v, "malformed subject"):
		err = Invalid(v)
	case strings.Contains(v, "maximum account active connections exceeded"):
		err = Connection(v)
	case strings.Contains(v, "maximum connections exceeded"):
		err = Limit(v)
	case strings.Contains(v, "maximum control line exceeded"):
		err = Limit(v)
	case strings.Contains(v, "maximum payload exceeded"):
		err = Limit(v)
	case strings.Contains(v, "maximum subscriptions exceeded"):
		err = Limit(v)
	case strings.Contains(v, "message headers not supported"):
		err = Invalid(v)
	case strings.Contains(v, "minimum version required"):
		err = Invalid(v)
	case strings.Contains(v, "no matching transforms available"):
		err = Invalid(v)
	case strings.Contains(v, "no responders requires headers support"):
		err = Invalid(v)
	case strings.Contains(v, "not enough arguments passed to the function"):
		err = Invalid(v)
	case strings.Contains(v, "not using all of the token wildcard(s)"):
		err = Invalid(v)
	case strings.Contains(v, "remote leafnode has same cluster name"):
		err = Exist(v)
	case strings.Contains(v, "reserved account"):
		err = Invalid(v)
	case strings.Contains(v, "reserved internal subject"):
		err = Invalid(v)
	case strings.Contains(v, "search cycle depth exhausted"):
		err = Invalid(v)
	case strings.Contains(v, "server is not running"):
		err = Server(v)
	case strings.Contains(v, "server name cannot contain spaces"):
		err = Invalid(v)
	case strings.Contains(v, "service import not authorized"):
		err = Access(v)
	case strings.Contains(v, "service missing"):
		err = NotFound(v)
	case strings.Contains(v, "stream import already exists"):
		err = Exist(v)
	case strings.Contains(v, "stream import not authorized"):
		err = Access(v)
	case strings.Contains(v, "stream import prefix can not contain wildcard tokens"):
		err = Invalid(v)
	case strings.Contains(v, "subject has exceeded number of tokens limit"):
		err = Limit(v)
	case strings.Contains(v, "subscribe permission violation"):
		err = Access(v)
	case strings.Contains(v, "system account not setup"):
		err = NotExist(v)
	case strings.Contains(v, "the only mapping function allowed for import transforms"):
		err = Invalid(v)
	case strings.Contains(v, "too many arguments passed to the function"):
		err = Invalid(v)
	case strings.Contains(v, "unknown function"):
		err = Try(v)
	case strings.Contains(v, "wildcard index out of range"):
		err = Invalid(v)
	case strings.Contains(v, "wrong gateway"):
		err = Connection(v)
	case strings.Contains(v, "no responders"):
		err = Unavailable(v)
	default:
		err = Try(v)
	}
	return
}
