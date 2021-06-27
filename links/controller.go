package links

type Controller struct {
	cu LinkCropUncrop
}

type LinkCropUncrop interface {
	CropLinc(link string) (*Link, error)
	UncropLink(link *Link) (*Link, error)
}

func NewController(cu LinkCropUncrop) *Controller {
	return &Controller{
		cu: cu,
	}
}

func (c *Controller) CropLinc(link string) (*Link, error) {
	return c.cu.CropLinc(link)
}

func (c *Controller) UncropLink(link *Link) (*Link, error) {
	return c.cu.UncropLink(link)
}
