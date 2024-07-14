package response

import "github.com/ginger-core/gateway"

func parseContentType(t gateway.ContentType) string {
	switch t {
	case gateway.ContentTypeJson:
		return "application/json"
	case gateway.ContentTypeImageJpeg:
		return "image/jpeg"
	}
	return string(t)
}
