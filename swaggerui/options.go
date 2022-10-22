package swaggerui

// WithTitle sets the title for the SwaggerUI [Config].
func WithTitle(title string) ConfigOption {
	return func(cfg *Config) {
		cfg.Title = title
	}
}

// WithSpecURL sets the spec url for the SwaggerUI [Config].
func WithSpecURL(specURL string) ConfigOption {
	return func(cfg *Config) {
		cfg.SpecURL = specURL
	}
}

// WithScriptURL sets the script url for the SwaggerUI [Config].
func WithScriptURL(scriptURL string) ConfigOption {
	return func(cfg *Config) {
		cfg.ScriptURL = scriptURL
	}
}

// WithPresetScriptURL sets the preset script url for the SwaggerUI [Config].
func WithPresetScriptURL(presetScriptURL string) ConfigOption {
	return func(cfg *Config) {
		cfg.PresetScriptURL = presetScriptURL
	}
}

// WithStylesURL sets the styles url for the SwaggerUI [Config].
func WithStylesURL(stylesURL string) ConfigOption {
	return func(cfg *Config) {
		cfg.StylesURL = stylesURL
	}
}

// WithFavicon32 sets the favicon32 for the SwaggerUI [Config].
func WithFavicon32(favicon32 string) ConfigOption {
	return func(cfg *Config) {
		cfg.Favicon32 = favicon32
	}
}

// WithFavicon16 sets the favicon16 for the SwaggerUI [Config].
func WithFavicon16(favicon16 string) ConfigOption {
	return func(cfg *Config) {
		cfg.Favicon16 = favicon16
	}
}

// WithTemplate sets the template for the SwaggerUI [Config].
func WithTemplate(template string) ConfigOption {
	return func(cfg *Config) {
		cfg.Template = template
	}
}
