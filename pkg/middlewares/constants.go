package middlewares

const (
	headerAllow           = "Allow"
	headerOrigin          = "Origin"
	headerVary            = "Vary"
	headerXForwardedProto = "X-Forwarded-Proto"

	contextKeyHeaderAllow = "boost_header_allow"

	// Access control
	headerAccessControlRequestHeaders   = "Access-Control-Request-Headers"
	headerAccessControlRequestMethod    = "Access-Control-Request-Method"
	headerAccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	headerAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	headerAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	headerAccessControlAllowMethods     = "Access-Control-Allow-Methods"
	headerAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	headerAccessControlMaxAge           = "Access-Control-Max-Age"

	// Security
	headerStrictTransportSecurity         = "Strict-Transport-Security"
	headerXContentTypeOptions             = "X-Content-Type-Options"
	headerXXSSProtection                  = "X-XSS-Protection"
	headerXFrameOptions                   = "X-Frame-Options"
	headerContentSecurityPolicy           = "Content-Security-Policy"
	headerContentSecurityPolicyReportOnly = "Content-Security-Policy-Report-Only"
	headerXCSRFToken                      = "X-CSRF-Token"
	headerReferrerPolicy                  = "Referrer-Policy"
)
