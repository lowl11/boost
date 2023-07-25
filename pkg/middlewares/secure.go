package middlewares

import (
	"fmt"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/pkg/enums/headers"
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
			res.Header.Set(headers.HeaderXXSSProtection, config.XSSProtection)
		}

		if config.ContentTypeNosniff != "" {
			res.Header.Set(headers.HeaderXContentTypeOptions, config.ContentTypeNosniff)
		}

		if config.XFrameOptions != "" {
			res.Header.Set(headers.HeaderXFrameOptions, config.XFrameOptions)
		}

		if (ctx.IsTLS() || (type_helper.BytesToString(req.Header.Peek(headers.HeaderXForwardedProto)) == "https")) &&
			config.HSTSMaxAge != 0 {
			subdomains := ""
			if !config.HSTSExcludeSubdomains {
				subdomains = "; includeSubdomains"
			}
			if config.HSTSPreloadEnabled {
				subdomains = fmt.Sprintf("%s; preload", subdomains)
			}
			res.Header.Set(headers.HeaderStrictTransportSecurity, fmt.Sprintf("max-age=%d%s", config.HSTSMaxAge, subdomains))
		}

		if config.ContentSecurityPolicy != "" {
			if config.CSPReportOnly {
				res.Header.Set(headers.HeaderContentSecurityPolicyReportOnly, config.ContentSecurityPolicy)
			} else {
				res.Header.Set(headers.HeaderContentSecurityPolicy, config.ContentSecurityPolicy)
			}
		}

		if config.ReferrerPolicy != "" {
			res.Header.Set(headers.HeaderReferrerPolicy, config.ReferrerPolicy)
		}

		if err := ctx.Next(); err != nil {
			return err
		}

		return nil
	}
}
