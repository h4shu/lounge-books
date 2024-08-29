package gateways

type (
	ImageSearchWeb interface {
		SearchImage()
	}
	ImageSearchAPI struct {
		web ImageSearchWeb
	}
)
