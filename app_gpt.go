package errors

import (
	"context"
	"strings"

	errs "errors"

	"github.com/sashabaranov/go-openai"
)

func GPT(er error) (err E) {
	if er == nil {
		return nil
	}

	// OpenAI / ChatGPT SDK
	var oe *openai.APIError
	if errs.As(er, &oe) {
		return errorByStatus(oe.HTTPStatusCode, er)
	}

	// Generic SDK errors that expose HTTP status.
	type statusCoder interface {
		StatusCode() int
	}
	var sc statusCoder
	if errs.As(er, &sc) {
		return errorByStatus(sc.StatusCode(), er)
	}

	type httpStatusCoder interface {
		HTTPStatusCode() int
	}
	var hsc httpStatusCoder
	if errs.As(er, &hsc) {
		return errorByStatus(hsc.HTTPStatusCode(), er)
	}

	// Context / network
	if errs.Is(er, context.DeadlineExceeded) {
		return Timeout(er)
	}
	if errs.Is(er, context.Canceled) {
		return Try(er)
	}

	msg := strings.ToLower(er.Error())

	switch {
	// Timeout / network
	case strings.Contains(msg, "timeout"),
		strings.Contains(msg, "timed out"),
		strings.Contains(msg, "deadline exceeded"),
		strings.Contains(msg, "context deadline exceeded"),
		strings.Contains(msg, "connection reset"),
		strings.Contains(msg, "connection refused"),
		strings.Contains(msg, "no such host"),
		strings.Contains(msg, "temporary failure"):
		return Timeout(er)

	// Access / auth
	case strings.Contains(msg, "unauthorized"),
		strings.Contains(msg, "permission denied"),
		strings.Contains(msg, "forbidden"),
		strings.Contains(msg, "invalid api key"),
		strings.Contains(msg, "api key not valid"),
		strings.Contains(msg, "incorrect api key"),
		strings.Contains(msg, "authentication"),
		strings.Contains(msg, "auth error"),
		strings.Contains(msg, "invalid x-api-key"):
		return Access(er)

	// Rate limit / quota / overload
	case strings.Contains(msg, "rate limit"),
		strings.Contains(msg, "rate_limit"),
		strings.Contains(msg, "too many requests"),
		strings.Contains(msg, "quota exceeded"),
		strings.Contains(msg, "insufficient quota"),
		strings.Contains(msg, "resource exhausted"),
		strings.Contains(msg, "overloaded"),
		strings.Contains(msg, "server overloaded"),
		strings.Contains(msg, "capacity"),
		strings.Contains(msg, "service unavailable"),
		strings.Contains(msg, "temporarily unavailable"),
		strings.Contains(msg, "try again later"):
		return TooManyRequests(er)

	// Safety / policy / moderation
	case strings.Contains(msg, "content policy"),
		strings.Contains(msg, "safety"),
		strings.Contains(msg, "blocked"),
		strings.Contains(msg, "moderation"),
		strings.Contains(msg, "harmful content"),
		strings.Contains(msg, "prompt was blocked"),
		strings.Contains(msg, "response was blocked"):
		return Access(er)

	default:
		return Try(er)
	}
}

func errorByStatus(code int, er error) E {
	switch code {
	case 400, 404, 422:
		return Try(er)

	case 401, 403:
		return Access(er)

	case 408, 504:
		return Timeout(er)

	case 409, 425, 429, 500, 502, 503:
		return TooManyRequests(er)

	default:
		if code >= 500 {
			return TooManyRequests(er)
		}
		if code >= 400 {
			return Try(er)
		}
		return Try(er)
	}
}

/*
ChatGPT Errors:
401 - Invalid Authentication														Cause: Invalid Authentication 																			Solution: Ensure the correct API key and requesting organization are being used.
401 - Incorrect API key provided													Cause: The requesting API key is not correct. 															Solution: Ensure the API key used is correct, clear your browser cache, or generate a new one.
401 - You must be a member of an organization to use the API						Cause: Your account is not part of an organization. 													Solution: Contact us to get added to a new organization or ask your organization manager to invite you to an organization.
401 - IP not authorized																Cause: Your request IP does not match the configured IP allowlist for your project or organization. 	Solution: Send the request from the correct IP, or update your IP allowlist settings.
403 - Country, region, or territory not supported									Cause: You are accessing the API from an unsupported country, region, or territory. 					Solution: Please see this page for more information.

429 - Rate limit reached for requests												Cause: You are sending requests too quickly. 															Solution: Pace your requests. Read the Rate limit guide.
429 - You exceeded your current quota, please check your plan and billing details	Cause: You have run out of credits or hit your maximum monthly spend. 									Solution: Buy more credits or learn how to increase your limits.
500 - The server had an error while processing your request							Cause: Issue on our servers. 																			Solution: Retry your request after a brief wait and contact us if the issue persists. Check the status page.
503 - The engine is currently overloaded, please try again later					Cause: Our servers are experiencing high traffic. 														Solution: Please retry your requests after a brief wait.
503 - Slow Down																		Cause: A sudden increase in your request rate is impacting service reliability. 						Solution: Please reduce your request rate to its original level, maintain a consistent rate for at least 15 minutes, and then gradually increase it.
*/
