package swaggerui

import (
	"strings"

	"github.com/linden-honey/linden-honey-sdk-go/errors"
)

// Validate validates a [Config] and returns an error if validation is failed.
func (cfg Config) Validate() error {
	if strings.TrimSpace(cfg.Title) == "" {
		return errors.NewInvalidValueError("Title", errors.ErrEmptyValue)
	}

	if strings.TrimSpace(cfg.SpecURL) == "" {
		return errors.NewInvalidValueError("SpecURL", errors.ErrEmptyValue)
	}

	if strings.TrimSpace(cfg.ScriptURL) == "" {
		return errors.NewInvalidValueError("ScriptURL", errors.ErrEmptyValue)
	}

	if strings.TrimSpace(cfg.PresetScriptURL) == "" {
		return errors.NewInvalidValueError("PresetScriptURL", errors.ErrEmptyValue)
	}

	if strings.TrimSpace(cfg.StylesURL) == "" {
		return errors.NewInvalidValueError("StylesURL", errors.ErrEmptyValue)
	}

	if strings.TrimSpace(cfg.Favicon32) == "" {
		return errors.NewInvalidValueError("Favicon32", errors.ErrEmptyValue)
	}

	if strings.TrimSpace(cfg.Favicon16) == "" {
		return errors.NewInvalidValueError("Favicon16", errors.ErrEmptyValue)
	}

	if strings.TrimSpace(cfg.Template) == "" {
		return errors.NewInvalidValueError("Template", errors.ErrEmptyValue)
	}

	return nil
}
