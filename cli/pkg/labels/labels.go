package labels

const (
	OdigosSystemLabelKey   = "vision.middleware.io/system-object"
	OdigosSystemLabelValue = "true"
)

var OdigosSystem = map[string]string{
	OdigosSystemLabelKey: OdigosSystemLabelValue,
}
