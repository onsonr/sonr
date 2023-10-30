package plugins_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pluginsconfig "github.com/sonr-io/core/config/plugins"
)

func TestPluginIsGlobal(t *testing.T) {
	assert.False(t, pluginsconfig.Plugin{}.IsGlobal())
	assert.True(t, pluginsconfig.Plugin{Global: true}.IsGlobal())
}

func TestPluginIsLocalPath(t *testing.T) {
	assert.False(t, pluginsconfig.Plugin{}.IsLocalPath())
	assert.False(t, pluginsconfig.Plugin{Path: "github.com/ignite/example"}.IsLocalPath())
	assert.True(t, pluginsconfig.Plugin{Path: "/home/bob/example"}.IsLocalPath())
}

func TestPluginHasPath(t *testing.T) {
	tests := []struct {
		name        string
		plugin      pluginsconfig.Plugin
		path        string
		expectedRes bool
	}{
		{
			name:        "empty both path",
			plugin:      pluginsconfig.Plugin{},
			expectedRes: false,
		},
		{
			name: "simple path",
			plugin: pluginsconfig.Plugin{
				Path: "github.com/ignite/example",
			},
			path:        "github.com/ignite/example",
			expectedRes: true,
		},
		{
			name: "plugin path with ref",
			plugin: pluginsconfig.Plugin{
				Path: "github.com/ignite/example@v1",
			},
			path:        "github.com/ignite/example",
			expectedRes: true,
		},
		{
			name: "plugin path with empty ref",
			plugin: pluginsconfig.Plugin{
				Path: "github.com/ignite/example@",
			},
			path:        "github.com/ignite/example",
			expectedRes: true,
		},
		{
			name: "both path with different ref",
			plugin: pluginsconfig.Plugin{
				Path: "github.com/ignite/example@v1",
			},
			path:        "github.com/ignite/example@v2",
			expectedRes: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.plugin.HasPath(tt.path)

			require.Equal(t, tt.expectedRes, res)
		})
	}
}

func TestPluginCanonicalPath(t *testing.T) {
	tests := []struct {
		name         string
		plugin       pluginsconfig.Plugin
		expectedPath string
	}{
		{
			name:         "empty both path",
			plugin:       pluginsconfig.Plugin{},
			expectedPath: "",
		},
		{
			name: "simple path",
			plugin: pluginsconfig.Plugin{
				Path: "github.com/ignite/example",
			},
			expectedPath: "github.com/ignite/example",
		},
		{
			name: "plugin path with ref",
			plugin: pluginsconfig.Plugin{
				Path: "github.com/ignite/example@v1",
			},
			expectedPath: "github.com/ignite/example",
		},
		{
			name: "plugin path with empty ref",
			plugin: pluginsconfig.Plugin{
				Path: "github.com/ignite/example@",
			},
			expectedPath: "github.com/ignite/example",
		},
		{
			name: "plugin local directory path",
			plugin: pluginsconfig.Plugin{
				Path: "/home/user/go/foo/bar",
			},
			expectedPath: "/home/user/go/foo/bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.plugin.CanonicalPath()
			require.Equal(t, tt.expectedPath, res)
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		configs  []pluginsconfig.Plugin
		expected []pluginsconfig.Plugin
	}{
		{
			name:     "do nothing for empty list",
			configs:  []pluginsconfig.Plugin(nil),
			expected: []pluginsconfig.Plugin(nil),
		},
		{
			name: "remove duplicates",
			configs: []pluginsconfig.Plugin{
				{
					Path: "foo/bar",
				},
				{
					Path: "foo/bar",
				},
				{
					Path: "bar/foo",
				},
			},
			expected: []pluginsconfig.Plugin{
				{
					Path: "foo/bar",
				},
				{
					Path: "bar/foo",
				},
			},
		},
		{
			name: "do nothing for no duplicates",
			configs: []pluginsconfig.Plugin{
				{
					Path: "foo/bar",
				},
				{
					Path: "bar/foo",
				},
			},
			expected: []pluginsconfig.Plugin{
				{
					Path: "foo/bar",
				},
				{
					Path: "bar/foo",
				},
			},
		},
		{
			name: "prioritize local plugins",
			configs: []pluginsconfig.Plugin{
				{
					Path:   "foo/bar",
					Global: true,
				},
				{
					Path:   "bar/foo",
					Global: true,
				},
				{
					Path:   "foo/bar",
					Global: false,
				},
				{
					Path:   "bar/foo",
					Global: false,
				},
			},
			expected: []pluginsconfig.Plugin{
				{
					Path:   "foo/bar",
					Global: false,
				},
				{
					Path:   "bar/foo",
					Global: false,
				},
			},
		},
		{
			name: "prioritize local plugins different versions",
			configs: []pluginsconfig.Plugin{
				{
					Path:   "foo/bar@v1",
					Global: true,
				},
				{
					Path:   "bar/foo",
					Global: true,
				},
				{
					Path:   "foo/bar@v2",
					Global: false,
				},
				{
					Path:   "bar/foo",
					Global: false,
				},
			},
			expected: []pluginsconfig.Plugin{
				{
					Path:   "foo/bar@v2",
					Global: false,
				},
				{
					Path:   "bar/foo",
					Global: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unique := pluginsconfig.RemoveDuplicates(tt.configs)
			require.EqualValues(t, tt.expected, unique)
		})
	}
}
