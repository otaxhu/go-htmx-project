package chi

import (
	"fmt"

	"github.com/elnormous/contenttype"
)

// MediaTypes
var (
	mtHtml      = contenttype.NewMediaType("text/html")
	mtJson      = contenttype.NewMediaType("application/json")
	mtForm      = contenttype.NewMediaType("application/x-www-form-urlencoded")
	mtMultiForm = contenttype.NewMediaType("multipart/form-data")

	// These are the available representations of this endpoint [/products]
	availableMediaTypes = []contenttype.MediaType{
		mtHtml,
		mtJson,
	}

	// This variable is intended to use as an error message if the Content-Type header
	// is not supported. Contains some information about the media types supported by the server
	unsupportedMediaTypeErr string = fmt.Sprintf("the supported media types are %s, %s, %s. Please use one of those as the Content-Type header of your request", mtJson.MIME(), mtForm.MIME(), mtMultiForm.MIME())
)
