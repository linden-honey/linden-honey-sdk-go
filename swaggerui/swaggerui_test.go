package swaggerui

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		opts []ConfigOption
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "ok",
			want: &Config{
				Title:           DefaultTitle,
				SpecURL:         DefaultSpecURL,
				ScriptURL:       ScriptLatest,
				PresetScriptURL: PresetScriptLatest,
				StylesURL:       StylesLatest,
				Favicon16:       Favicon16Latest,
				Favicon32:       Favicon32Latest,
				Template:        DefaultTemplate,
			},
		},
		{
			name: "ok  with opts",
			args: args{
				opts: []ConfigOption{
					WithTitle("Test"),
					WithSpecURL("/test-swagger.json"),
					WithScriptURL("/test-swagger-ui.js"),
					WithPresetScriptURL("/test-swagger-ui-preset.js"),
					WithStylesURL("/test-swagger-ui.css"),
					WithFavicon32("/test-swagger-32x32.png"),
					WithFavicon16("/test-swagger-16x16.png"),
				},
			},
			want: &Config{
				Title:           "Test",
				SpecURL:         "/test-swagger.json",
				ScriptURL:       "/test-swagger-ui.js",
				PresetScriptURL: "/test-swagger-ui-preset.js",
				StylesURL:       "/test-swagger-ui.css",
				Favicon16:       "/test-swagger-32x32.png",
				Favicon32:       "/test-swagger-32x32.png",
				Template:        DefaultTemplate,
			},
		},
		{
			name: "err  invalid config",
			args: args{
				opts: []ConfigOption{
					WithTitle(""),
				},
			},
			wantErr: true,
		},
		{
			name: "err  invalid template  failed to parse",
			args: args{
				opts: []ConfigOption{
					WithTemplate("{{"),
				},
			},
			wantErr: true,
		},
		{
			name: "err  invalid template  failed to execute",
			args: args{
				opts: []ConfigOption{
					WithTemplate("{{ .Test }}"),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.opts...)
			if tt.wantErr {
				if err == nil {
					t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				}

				return
			}

			req := httptest.NewRequest("GET", "/", nil)
			rec := httptest.NewRecorder()
			got.ServeHTTP(rec, req)

			if got, want := rec.Code, http.StatusOK; got != want {
				t.Errorf("Response.Code = %v, want = %v", got, want)
			}

			if got, want := rec.Header().Get("Content-Type"), "text/html; charset=utf-8"; got != want {
				t.Errorf(`Response..Header().Get("Content-Type") = %v, want = %v`, got, want)
			}

			if !strings.Contains(rec.Body.String(), tt.want.Title) {
				t.Errorf("Response.Body doesn't contain .Title = %v", tt.want.Title)
			}

			if !strings.Contains(rec.Body.String(), tt.want.SpecURL) {
				t.Errorf("Response.Body doesn't contain .SpecURL = %v", tt.want.SpecURL)
			}

			if !strings.Contains(rec.Body.String(), tt.want.ScriptURL) {
				t.Errorf("Response.Body doesn't contain .ScriptURL = %v", tt.want.ScriptURL)
			}

			if !strings.Contains(rec.Body.String(), tt.want.PresetScriptURL) {
				t.Errorf("Response.Body doesn't contain .PresetScriptURL = %v", tt.want.PresetScriptURL)
			}

			if !strings.Contains(rec.Body.String(), tt.want.StylesURL) {
				t.Errorf("Response.Body doesn't contain .StylesURL = %v", tt.want.StylesURL)
			}

			if !strings.Contains(rec.Body.String(), tt.want.Favicon32) {
				t.Errorf("Response.Body doesn't contain .Favicon32 = %v", tt.want.Favicon32)
			}

			if !strings.Contains(rec.Body.String(), tt.want.Favicon16) {
				t.Errorf("Response.Body doesn't contain .Favicon16 = %v", tt.want.Favicon16)
			}
		})
	}
}
