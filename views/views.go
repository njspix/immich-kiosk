package views

import (
	"net/url"

	"github.com/damongolding/immich-kiosk/config"
)

type PageData struct {
	// KioskVersion the current build version of Kiosk
	KioskVersion string
	// ImageID id of image to be displayed
	ImageID string
	// ImageData image as base64 data
	ImageData string
	// ImageData blurred image as base64 data
	ImageBlurData string
	// Date image date
	ImageDate string
	// URL queries
	Queries url.Values
	// instance config
	config.Config
}
