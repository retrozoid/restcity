package trac

// Represents the specific server metric.
type Metric struct {
	Name           string        `json:"name,omitempty"`
	Description    string        `json:"description,omitempty"`
	PrometheusName string        `json:"prometheusName,omitempty"`
	MetricValues   *MetricValues `json:"metricValues,omitempty"`
	MetricTags     *MetricTags   `json:"metricTags,omitempty"`
}

// Represents a list of Metric entities.
type Metrics struct {
	Count  int32    `json:"count,omitempty"`
	Metric []Metric `json:"metric,omitempty"`
}

// Represents a list of MetricValue entities.
type MetricValues struct {
	Count       int32         `json:"count,omitempty"`
	MetricValue []MetricValue `json:"metricValue,omitempty"`
}

// Represents a metric value.
type MetricValue struct {
	Name  string  `json:"name,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// Represents a list of MetricTag entities.
type MetricTags struct {
	Count     int32       `json:"count,omitempty"`
	MetricTag []MetricTag `json:"metricTag,omitempty"`
}

// Represents a metric tag.
type MetricTag struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
