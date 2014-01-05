package types

import (
	"logger"
	"plugins/command"
	"plugins/cpu_usage"
	"plugins/fs_usage"
	"plugins/io_usage"
	"plugins/json_poll"
	"plugins/load_average"
	"plugins/mem_usage"
	"plugins/net_usage"
	"plugins/process_count"
	"plugins/record"
	"plugins/socket_usage"
)

type PluginResult interface{}
type PluginResultCollection map[string]PluginResult

var Plugins = map[string]func(interface{}, *logger.Logger) interface{}{
	"load_average":  load_average.GetMetric,
	"cpu_usage":     cpu_usage.GetMetric,
	"mem_usage":     mem_usage.GetMetric,
	"command":       command.GetMetric,
	"net_usage":     net_usage.GetMetric,
	"io_usage":      io_usage.GetMetric,
	"record":        record.GetMetric,
	"fs_usage":      fs_usage.GetMetric,
	"json_poll":     json_poll.GetMetric,
	"socket_usage":  socket_usage.GetMetric,
	"process_count": process_count.GetMetric,
}

/*
How this works:

Basically, interface is expected to be nil or an array of strings which are
single parameters passed to the Params section of each json. Each element of
the array is treated as a Params line and passed straight to the generated
json. A params of "" is treated as nil because go is kind of stupid about nils.

Returning nil means to not include the monitor plugin, but it is wiser to just
exclude the plugin from this map if we never want to use it. This is
principally for those who would dispatch to a detection method.

An example from below: In the load average case, our params are "", but we want
to always include it.
*/

var Detectors = map[string]func() []string{
	"load_average": func() []string { return []string{} },
	"cpu_usage":    func() []string { return []string{} },
	"mem_usage":    func() []string { return []string{} },
	"net_usage":    net_usage.Detect,
	"io_usage":     io_usage.Detect,
	"fs_usage":     fs_usage.Detect,
}

type ConfigMap struct {
	Type   string
	Params interface{}
}

type PluginConfig map[string]ConfigMap

type CirconusConfig struct {
	Listen       string
	Username     string
	Password     string
	Facility     string
	LogLevel     string
	PollInterval uint
	Plugins      PluginConfig
}
