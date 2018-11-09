package collector

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	metricNameReplacer = strings.NewReplacer(
		"/", "_",
		"-", "_",
		"'", "",
	)
	labelValueReplacer = strings.NewReplacer(
		"'", "",
	)
	labelKeyReplacer = strings.NewReplacer(
		"-", "_",
		"'", "",
	)
)

type metric struct {
	name              string
	labels            []string
	descriptionLabels []string
	value             interface{}
}

type label struct {
	key   string
	value string
}

func (m *metric) promName() string {
	return metricNameReplacer.Replace(m.name)
}

func (m *metric) Key() string {
	ret := m.name + "{"
	for _, label := range m.labels {
		ret += label + ","
	}

	ret += "}"
	return ret
}

func (m *metric) describe() *prometheus.Desc {
	return prometheus.NewDesc(m.promName(), m.name, m.promLabelKeys(), nil)
}

func (m *metric) promLabelValues() []string {
	values := m.labelValues()
	res := make([]string, len(values))

	for i, v := range values {
		res[i] = labelValueReplacer.Replace(v)
	}

	return res
}

func (m *metric) labelValues() []string {
	res := make([]string, len(m.labels)+len(m.descriptionLabels))

	i := 0
	for _, l := range m.labels {
		parts := strings.Split(l, "=")
		if len(parts) != 2 {
			continue
		}

		res[i] = parts[1]
		i++
	}

	for _, l := range m.descriptionLabels {
		parts := strings.Split(l, "=")
		if len(parts) != 2 {
			continue
		}

		res[i] = parts[1]
		i++
	}

	return res
}

func (m *metric) promLabelKeys() []string {
	keys := m.labelKeys()
	res := make([]string, len(keys))

	for i, k := range keys {
		res[i] = labelKeyReplacer.Replace(k)
	}

	return res
}

func (m *metric) labelKeys() []string {
	res := make([]string, len(m.labels)+len(m.descriptionLabels))

	i := 0
	for _, l := range m.labels {
		parts := strings.Split(l, "=")
		if len(parts) != 2 {
			continue
		}

		res[i] = parts[0]
		i++
	}

	for _, l := range m.descriptionLabels {
		parts := strings.Split(l, "=")
		if len(parts) != 2 {
			continue
		}

		res[i] = parts[0]
		i++
	}

	return res
}
