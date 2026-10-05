package handler

import "net/url"

// decodePathParam decodes URL-escaped IMAP mailbox names.
// Example:
//   09.%20%D0%9D%D0%A3|09.1%20ESF
// becomes:
//   09. НУ|09.1 ESF
func decodePathParam(s string) string {
	v, err := url.PathUnescape(s)
	if err != nil {
		return s
	}
	return v
}
