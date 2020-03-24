package globals

// ExceptMap Exception information mapping table
var ExceptMap map[string]string = map[string]string{
	"NoCommand": "Err: No command given.",
}

var (
	// NoCommand No parameters passed
	NoCommand string = ExceptMap["NoCommand"]
)
