package middlewares

import (
	"fmt"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/pkg/interfaces"
	"github.com/lowl11/boost/pkg/types"
)

type SecureConfig struct {
	XSSProtection         string
	ContentTypeNosniff    string
	XFrameOptions         string
	HSTSMaxAge            int
	HSTSExcludeSubdomains bool
	ContentSecurityPolicy string
	CSPReportOnly         bool
	HSTSPreloadEnabled    bool
	ReferrerPolicy        string
}

func defaultSecureConfig() SecureConfig {
	return SecureConfig{}
}

func Secure() types.MiddlewareFunc {
	return SecureWithConfig(defaultSecureConfig())
}

func SecureWithConfig(config SecureConfig) types.MiddlewareFunc {
	return func(ctx interfaces.Context) error {
		req := ctx.Request()
		res := ctx.Response()

		if config.XSSProtection != "" {
			res.Header.Set(headerXXSSProtection, config.XSSProtection)
		}

		if config.ContentTypeNosniff != "" {
			res.Header.Set(headerXContentTypeOptions, config.ContentTypeNosniff)
		}

		if config.XFrameOptions != "" {
			res.Header.Set(headerXFrameOptions, config.XFrameOptions)
		}

		if (ctx.IsTLS() || (type_helper.BytesToString(req.Header.Peek(headerXForwardedProto)) == "https")) && config.HSTSMaxAge != 0 {
			subdomains := ""
			if !config.HSTSExcludeSubdomains {
				subdomains = "; includeSubdomains"
			}
			if config.HSTSPreloadEnabled {
				subdomains = fmt.Sprintf("%s; preload", subdomains)
			}
			res.Header.Set(headerStrictTransportSecurity, fmt.Sprintf("max-age=%d%s", config.HSTSMaxAge, subdomains))
		}

		if config.ContentSecurityPolicy != "" {
			if config.CSPReportOnly {
				res.Header.Set(headerContentSecurityPolicyReportOnly, config.ContentSecurityPolicy)
			} else {
				res.Header.Set(headerContentSecurityPolicy, config.ContentSecurityPolicy)
			}
		}

		if config.ReferrerPolicy != "" {
			res.Header.Set(headerReferrerPolicy, config.ReferrerPolicy)
		}

		if err := ctx.Next(); err != nil {
			return err
		}

		return nil
	}
}
