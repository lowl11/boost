package middlewares

import (
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/pkg/interfaces"
	"github.com/lowl11/boost/pkg/types"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type CORSConfig struct {
	AllowOrigins                             []string
	AllowOriginFunc                          func(origin string) (bool, error)
	AllowMethods                             []string
	AllowHeaders                             []string
	AllowCredentials                         bool
	UnsafeWildcardOriginWithAllowCredentials bool
	ExposeHeaders                            []string
	MaxAge                                   int
}

func defaultCorsConfig() CORSConfig {
	return CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
	}
}

func CORS() types.MiddlewareFunc {
	return CORSWithConfig(defaultCorsConfig())
}

func CORSWithConfig(config CORSConfig) types.MiddlewareFunc {
	defaultConfig := defaultCorsConfig()

	if len(config.AllowOrigins) == 0 {
		config.AllowOrigins = defaultConfig.AllowOrigins
	}

	hasCustomAllowMethods := true
	if len(config.AllowMethods) == 0 {
		hasCustomAllowMethods = false
		config.AllowMethods = defaultConfig.AllowMethods
	}

	allowOriginPatterns := make([]string, 0, len(config.AllowOrigins))
	for _, origin := range config.AllowOrigins {
		pattern := regexp.QuoteMeta(origin)
		pattern = strings.Replace(pattern, "\\*", ".*", -1)
		pattern = strings.Replace(pattern, "\\?", ".", -1)
		pattern = "^" + pattern + "$"
		allowOriginPatterns = append(allowOriginPatterns, pattern)
	}

	allowMethods := strings.Join(config.AllowMethods, ",")
	allowHeaders := strings.Join(config.AllowHeaders, ",")
	exposeHeaders := strings.Join(config.ExposeHeaders, ",")
	maxAge := strconv.Itoa(config.MaxAge)

	return func(ctx interfaces.Context) error {
		// before next

		req := ctx.Request()
		res := ctx.Response()

		origin := type_helper.BytesToString(req.Header.Peek(headerOrigin))
		var allowOrigin string

		res.Header.Set(headerVary, headerOrigin)

		// Preflight request is an OPTIONS request, using three HTTP request headers: Access-Control-Request-Method,
		// Access-Control-Request-Headers, and the Origin header. See: https://developer.mozilla.org/en-US/docs/Glossary/Preflight_request
		// For simplicity we just consider method type and later `Origin` header.
		preflight := ctx.Method() == http.MethodOptions

		routerAllowMethods := ""
		if preflight {
			tmpAllowMethods, ok := ctx.Get(contextKeyHeaderAllow).(string)
			if ok && tmpAllowMethods != "" {
				routerAllowMethods = tmpAllowMethods
				ctx.Response().Header.Set(headerAllow, routerAllowMethods)
			}
		}

		if err := ctx.Next(); err != nil {
			return err
		}

		// after next

		if config.AllowOriginFunc != nil {
			allowed, err := config.AllowOriginFunc(origin)
			if err != nil {
				return err
			}

			if allowed {
				allowOrigin = origin
			}
		} else {
			// Check allowed origins
			for _, o := range config.AllowOrigins {
				if o == "*" && config.AllowCredentials && config.UnsafeWildcardOriginWithAllowCredentials {
					allowOrigin = origin
					break
				}

				if o == "*" || o == origin {
					allowOrigin = o
					break
				}

				if matchSubdomain(origin, o) {
					allowOrigin = origin
					break
				}
			}

			checkPatterns := false
			if allowOrigin == "" {
				// to avoid regex cost by invalid (long) domains (253 is domain name max limit)
				if len(origin) <= (253+3+5) && strings.Contains(origin, "://") {
					checkPatterns = true
				}
			}

			if checkPatterns {
				for _, re := range allowOriginPatterns {
					if match, _ := regexp.MatchString(re, origin); match {
						allowOrigin = origin
						break
					}
				}
			}
		}

		// Origin not allowed
		if allowOrigin == "" {
			if !preflight {
				return ctx.Next()
			}

			return ctx.Status(http.StatusNoContent).Empty()
		}

		res.Header.Set(headerAccessControlAllowOrigin, allowOrigin)
		if config.AllowCredentials {
			res.Header.Set(headerAccessControlAllowCredentials, "true")
		}

		// Simple request
		if !preflight {
			if exposeHeaders != "" {
				res.Header.Set(headerAccessControlExposeHeaders, exposeHeaders)
			}

			return ctx.Next()
		}

		// Preflight request
		res.Header.Add(headerVary, headerAccessControlRequestMethod)
		res.Header.Add(headerVary, headerAccessControlRequestHeaders)

		if !hasCustomAllowMethods && routerAllowMethods != "" {
			res.Header.Set(headerAccessControlAllowMethods, routerAllowMethods)
		} else {
			res.Header.Set(headerAccessControlAllowMethods, allowMethods)
		}

		if allowHeaders != "" {
			res.Header.Set(headerAccessControlAllowHeaders, allowHeaders)
		} else {
			h := type_helper.BytesToString(req.Header.Peek(headerAccessControlRequestHeaders))
			if h != "" {
				res.Header.Set(headerAccessControlAllowHeaders, h)
			}
		}

		if config.MaxAge > 0 {
			res.Header.Set(headerAccessControlMaxAge, maxAge)
		}

		return ctx.Status(http.StatusNoContent).Empty()
	}
}

func matchScheme(domain, pattern string) bool {
	didx := strings.Index(domain, ":")
	pidx := strings.Index(pattern, ":")
	return didx != -1 && pidx != -1 && domain[:didx] == pattern[:pidx]
}

func matchSubdomain(domain, pattern string) bool {
	if !matchScheme(domain, pattern) {
		return false
	}
	didx := strings.Index(domain, "://")
	pidx := strings.Index(pattern, "://")
	if didx == -1 || pidx == -1 {
		return false
	}
	domAuth := domain[didx+3:]
	// to avoid long loop by invalid long domain
	if len(domAuth) > 253 {
		return false
	}
	patAuth := pattern[pidx+3:]

	domComp := strings.Split(domAuth, ".")
	patComp := strings.Split(patAuth, ".")
	for i := len(domComp)/2 - 1; i >= 0; i-- {
		opp := len(domComp) - 1 - i
		domComp[i], domComp[opp] = domComp[opp], domComp[i]
	}
	for i := len(patComp)/2 - 1; i >= 0; i-- {
		opp := len(patComp) - 1 - i
		patComp[i], patComp[opp] = patComp[opp], patComp[i]
	}

	for i, v := range domComp {
		if len(patComp) <= i {
			return false
		}
		p := patComp[i]
		if p == "*" {
			return true
		}
		if p != v {
			return false
		}
	}
	return false
}
