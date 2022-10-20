package swaggerui

import (
	"testing"
)

func TestConfig_Validate(t *testing.T) {
	type fields struct {
		Title           string
		SpecURL         string
		ScriptURL       string
		PresetScriptURL string
		StylesURL       string
		Favicon32       string
		Favicon16       string
		Template        string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Title:           "Test",
				SpecURL:         "/test-swagger.json",
				ScriptURL:       "/test-swagger-ui.js",
				PresetScriptURL: "/test-swagger-ui-preset.js",
				StylesURL:       "/test-swagger-ui.css",
				Favicon16:       "/test-swagger-32x32.png",
				Favicon32:       "/test-swagger-32x32.png",
				Template:        "{{ .Test }}",
			},
		},
		{
			name: "err  empty title",
			fields: fields{
				Title:           "",
				SpecURL:         "/test-swagger.json",
				ScriptURL:       "/test-swagger-ui.js",
				PresetScriptURL: "/test-swagger-ui-preset.js",
				StylesURL:       "/test-swagger-ui.css",
				Favicon16:       "/test-swagger-32x32.png",
				Favicon32:       "/test-swagger-32x32.png",
				Template:        "{{ .Test }}",
			},
			wantErr: true,
		},
		{
			name: "err  empty spec url",
			fields: fields{
				Title:           "Test",
				SpecURL:         "",
				ScriptURL:       "/test-swagger-ui.js",
				PresetScriptURL: "/test-swagger-ui-preset.js",
				StylesURL:       "/test-swagger-ui.css",
				Favicon16:       "/test-swagger-32x32.png",
				Favicon32:       "/test-swagger-32x32.png",
				Template:        "{{ .Test }}",
			},
			wantErr: true,
		},
		{
			name: "err  empty script url",
			fields: fields{
				Title:           "Test",
				SpecURL:         "/test-swagger.json",
				ScriptURL:       "",
				PresetScriptURL: "/test-swagger-ui-preset.js",
				StylesURL:       "/test-swagger-ui.css",
				Favicon16:       "/test-swagger-32x32.png",
				Favicon32:       "/test-swagger-32x32.png",
				Template:        "{{ .Test }}",
			},
			wantErr: true,
		},
		{
			name: "err  empty preset script url",
			fields: fields{
				Title:           "Test",
				SpecURL:         "/test-swagger.json",
				ScriptURL:       "/test-swagger-ui.js",
				PresetScriptURL: "",
				StylesURL:       "/test-swagger-ui.css",
				Favicon16:       "/test-swagger-32x32.png",
				Favicon32:       "/test-swagger-32x32.png",
				Template:        "{{ .Test }}",
			},
			wantErr: true,
		}, {
			name: "err  empty styles url",
			fields: fields{
				Title:           "Test",
				SpecURL:         "/test-swagger.json",
				ScriptURL:       "/test-swagger-ui.js",
				PresetScriptURL: "/test-swagger-ui-preset.js",
				StylesURL:       "",
				Favicon16:       "/test-swagger-32x32.png",
				Favicon32:       "/test-swagger-32x32.png",
				Template:        "{{ .Test }}",
			},
			wantErr: true,
		},
		{
			name: "err  empty favicon16",
			fields: fields{
				Title:           "Test",
				SpecURL:         "/test-swagger.json",
				ScriptURL:       "/test-swagger-ui.js",
				PresetScriptURL: "/test-swagger-ui-preset.js",
				StylesURL:       "/test-swagger-ui.css",
				Favicon16:       "",
				Favicon32:       "/test-swagger-32x32.png",
				Template:        "{{ .Test }}",
			},
			wantErr: true,
		},
		{
			name: "err  empty favicon32",
			fields: fields{
				Title:           "Test",
				SpecURL:         "/test-swagger.json",
				ScriptURL:       "/test-swagger-ui.js",
				PresetScriptURL: "/test-swagger-ui-preset.js",
				StylesURL:       "/test-swagger-ui.css",
				Favicon16:       "/test-swagger-32x32.png",
				Favicon32:       "",
				Template:        "{{ .Test }}",
			},
			wantErr: true,
		},
		{
			name: "err  empty template",
			fields: fields{
				Title:           "Test",
				SpecURL:         "/test-swagger.json",
				ScriptURL:       "/test-swagger-ui.js",
				PresetScriptURL: "/test-swagger-ui-preset.js",
				StylesURL:       "/test-swagger-ui.css",
				Favicon16:       "/test-swagger-32x32.png",
				Favicon32:       "/test-swagger-32x32.png",
				Template:        "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Config{
				Title:           tt.fields.Title,
				SpecURL:         tt.fields.SpecURL,
				ScriptURL:       tt.fields.ScriptURL,
				PresetScriptURL: tt.fields.PresetScriptURL,
				StylesURL:       tt.fields.StylesURL,
				Favicon32:       tt.fields.Favicon32,
				Favicon16:       tt.fields.Favicon16,
				Template:        tt.fields.Template,
			}
			if err := cfg.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
