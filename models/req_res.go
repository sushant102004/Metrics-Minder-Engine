package models

type GAReq struct {
	Dimensions    []Dimensions `json:"dimensions"`
	Metrics       []Metrics    `json:"metrics"`
	DateRanges    []DateRanges `json:"dateRanges"`
	OrderBys      []OrderBys   `json:"orderBys"`
	KeepEmptyRows bool         `json:"keepEmptyRows"`
}
type Dimensions struct {
	Name string `json:"name"`
}
type Metrics struct {
	Name string `json:"name"`
}
type DateRanges struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
type Dimension struct {
	OrderType     string `json:"orderType"`
	DimensionName string `json:"dimensionName"`
}
type OrderBys struct {
	Dimension Dimension `json:"dimension"`
}

type GAResp struct {
	DimensionHeaders []DimensionHeaders `json:"dimensionHeaders"`
	MetricHeaders    []MetricHeaders    `json:"metricHeaders"`
	Rows             []Rows             `json:"rows"`
	RowCount         int                `json:"rowCount"`
	Metadata         Metadata           `json:"metadata"`
	Kind             string             `json:"kind"`
}
type DimensionHeaders struct {
	Name string `json:"name"`
}
type MetricHeaders struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type DimensionValues struct {
	Value string `json:"value"`
}
type MetricValues struct {
	Value string `json:"value"`
}
type Rows struct {
	DimensionValues []DimensionValues `json:"dimensionValues"`
	MetricValues    []MetricValues    `json:"metricValues"`
}
type Metadata struct {
	CurrencyCode string `json:"currencyCode"`
	TimeZone     string `json:"timeZone"`
}

