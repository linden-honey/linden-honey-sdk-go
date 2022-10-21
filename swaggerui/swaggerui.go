package swaggerui

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
)

// New returns a new instance of [http.Handler] with SwaggerUI or an error.
func New(opts ...ConfigOption) (http.Handler, error) {
	cfg := &Config{
		Title:           DefaultTitle,
		SpecURL:         DefaultSpecURL,
		ScriptURL:       ScriptLatest,
		PresetScriptURL: PresetScriptLatest,
		StylesURL:       StylesLatest,
		Favicon32:       Favicon32Latest,
		Favicon16:       Favicon16Latest,
		Template:        DefaultTemplate,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	tmpl := template.New("swagger_ui")
	tmpl, err := tmpl.Parse(cfg.Template)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	buf := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buf, cfg); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	b := buf.Bytes()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(b)
	}), nil
}

// Config represents the SwaggerUI configuration.
type Config struct {
	Title   string
	SpecURL string

	ScriptURL       string
	PresetScriptURL string
	StylesURL       string

	Favicon32 string
	Favicon16 string

	Template string
}

// ConfigOption sets optional parameters for the SwaggerUI [Config].
type ConfigOption func(*Config)

const (
	DefaultTitle   = "API Documentation"
	DefaultSpecURL = "/swagger.json"

	ScriptLatest       = "https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js"
	PresetScriptLatest = "https://unpkg.com/swagger-ui-dist/swagger-ui-standalone-preset.js"
	StylesLatest       = "https://unpkg.com/swagger-ui-dist/swagger-ui.css"
	Favicon32Latest    = "https://unpkg.com/swagger-ui-dist/favicon-32x32.png"
	Favicon16Latest    = "https://unpkg.com/swagger-ui-dist/favicon-16x16.png"
)

//go:embed swaggerui.html.tmpl
var DefaultTemplate string
