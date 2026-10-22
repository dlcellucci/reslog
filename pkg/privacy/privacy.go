package privacy

import "regexp"

var (
	emailRegex = regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	ipRegex    = regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)
	cardRegex  = regexp.MustCompile(`\b(?:\d[ -]*?){13,16}\b`)
)

type Options struct {
	RedactEmails bool
	RedactIPs    bool
	RedactCards  bool
}

func RedactWithOptions(input string, opts Options) string {
	out := input
	if opts.RedactEmails {
		out = emailRegex.ReplaceAllString(out, "[REDACTED EMAIL]")
	}
	if opts.RedactIPs {
		out = ipRegex.ReplaceAllString(out, "[REDACTED IP]")
	}
	if opts.RedactCards {
		out = cardRegex.ReplaceAllString(out, "[REDACTED CARD]")
	}
	return out
}
