package swaggerui

import (
	"strings"

	sdkerrors "github.com/linden-honey/linden-honey-sdk-go/errors"
)

// Validate validates a [Config] and returns an error if validation is failed.
func (cfg Config) Validate() error {
	if strings.TrimSpace(cfg.Title) == "" {
		return sdkerrors.NewInvalidValueError("Title", sdkerrors.ErrEmptyValue)
	}

	if strings.TrimSpace(cfg.SpecURL) == "" {
		return sdkerrors.NewInvalidValueError("SpecURL", sdkerrors.ErrEmptyValue)
	}

	if strings.TrimSpace(cfg.ScriptURL) == "" {
		return sdkerrors.NewInvalidValueError("ScriptURL", sdkerrors.ErrEmptyValue)
	}

	if strings.TrimSpace(cfg.PresetScriptURL) == "" {
		return sdkerrors.NewInvalidValueError("PresetScriptURL", sdkerrors.ErrEmptyValue)
	}

	if strings.TrimSpace(cfg.StylesURL) == "" {
		return sdkerrors.NewInvalidValueError("StylesURL", sdkerrors.ErrEmptyValue)
	}

	if strings.TrimSpace(cfg.Favicon32) == "" {
		return sdkerrors.NewInvalidValueError("Favicon32", sdkerrors.ErrEmptyValue)
	}

	if strings.TrimSpace(cfg.Favicon16) == "" {
		return sdkerrors.NewInvalidValueError("Favicon16", sdkerrors.ErrEmptyValue)
	}

	if strings.TrimSpace(cfg.Template) == "" {
		return sdkerrors.NewInvalidValueError("Template", sdkerrors.ErrEmptyValue)
	}

	return nil
}
