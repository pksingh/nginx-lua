package main

var (
	ADD_HOST = GetEnvOrDefault("ADD_HOST", "localhost")
	ADD_PORT = GetEnvOrDefault("ADD_PORT", "8081")
	ADD_URI  = GetEnvOrDefault("ADD_URI", "/api/v1/add")

	SUB_HOST = GetEnvOrDefault("SUB_HOST", "localhost")
	SUB_PORT = GetEnvOrDefault("SUB_PORT", "8082")
	SUB_URI  = GetEnvOrDefault("SUB_URI", "/api/v1/sub")


	APP_VERSION = "v1"
	APP_SERVICE = "name: calcparser, version: " + APP_VERSION
)

// ResultResponse is the Final Response served
type ResultResponse struct {
	Input   string        `json:"input" binding:"required"`
	Origins []interface{} `json:"origins,omitempty" binding:"required"`
	// Operands         `json:"operands" binding:"required"`
	Result  int    `json:"result" binding:"required"`
	Service string `json:"service" binding:"required"`
}

// NodeResponse received from each of micro services
type NodeResponse struct {
	Operands [2]int `json:"operands" binding:"required"`
	Result   int    `json:"result" binding:"required"`
	Service  string `json:"service" binding:"required"`
}

type Node struct {
	host string
	port string
	uri  string
}
func GetEnvOrDefault(name string, def string) string {
	name, ok := os.LookupEnv(name)
	if !ok {
		return def
	}
	return name
}
