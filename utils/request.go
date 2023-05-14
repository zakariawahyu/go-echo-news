package utils

type Request struct {
	Keyword   string
	Page      int64
	PerPage   int64
	Offset    int64
	SortBy    string
	SortOrder string `default:"desc"`
	StartDate string
	EndDate   string
	Filters   map[string]interface{}
}
