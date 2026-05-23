package server

import (
	"encoding/json"

	"github.com/WindowsSov8forUs/sonolus-core-go/core"
	coreerrors "github.com/WindowsSov8forUs/sonolus-core-go/errors"
)

type unionTypeHeader struct {
	Type string `json:"type"`
}

type ServerAuthenticateRequest struct {
	Type        string             `json:"type"`
	Address     string             `json:"address"`
	Time        float64            `json:"time"`
	UserProfile ServiceUserProfile `json:"userProfile"`
}

type ServerAuthenticateResponse struct {
	Session    string  `json:"session"`
	Expiration float64 `json:"expiration"`
}

type ServerMessage struct {
	Message string `json:"message,omitempty"`
}

type ServerConfiguration struct {
	Options []ServerOption `json:"options"`
}

type ServerForm struct {
	Type                string         `json:"type"`
	Title               core.Text      `json:"title"`
	Icon                core.Icon      `json:"icon,omitempty"`
	Description         string         `json:"description,omitempty"`
	Help                string         `json:"help,omitempty"`
	RequireConfirmation bool           `json:"requireConfirmation"`
	Options             []ServerOption `json:"options"`
}

type ServerOption interface {
	serverOption()
	OptionType() string
}

type ServerOptionBase struct {
	Query       string    `json:"query"`
	Name        core.Text `json:"name"`
	Description string    `json:"description,omitempty"`
	Required    bool      `json:"required"`
}

func (ServerOptionBase) serverOption() {}

type ServerTextOption struct {
	ServerOptionBase
	Type        string   `json:"type"`
	Def         string   `json:"def"`
	Placeholder string   `json:"placeholder"`
	Limit       int      `json:"limit"`
	Shortcuts   []string `json:"shortcuts"`
}

func (ServerTextOption) OptionType() string { return "text" }

type ServerTextAreaOption struct {
	ServerOptionBase
	Type        string   `json:"type"`
	Def         string   `json:"def"`
	Placeholder string   `json:"placeholder"`
	Limit       int      `json:"limit"`
	Shortcuts   []string `json:"shortcuts"`
}

func (ServerTextAreaOption) OptionType() string { return "textArea" }

type ServerSliderOption struct {
	ServerOptionBase
	Type string    `json:"type"`
	Def  float64   `json:"def"`
	Min  float64   `json:"min"`
	Max  float64   `json:"max"`
	Step float64   `json:"step"`
	Unit core.Text `json:"unit,omitempty"`
}

func (ServerSliderOption) OptionType() string { return "slider" }

type ServerToggleOption struct {
	ServerOptionBase
	Type string `json:"type"`
	Def  bool   `json:"def"`
}

func (ServerToggleOption) OptionType() string { return "toggle" }

type ServerSelectOption struct {
	ServerOptionBase
	Type   string                    `json:"type"`
	Def    string                    `json:"def"`
	Values []ServerSelectOptionValue `json:"values"`
}

func (ServerSelectOption) OptionType() string { return "select" }

type ServerSelectOptionValue struct {
	Name  string    `json:"name"`
	Title core.Text `json:"title"`
}

type ServerMultiOption struct {
	ServerOptionBase
	Type   string                    `json:"type"`
	Def    []bool                    `json:"def"`
	Values []ServerSelectOptionValue `json:"values"`
}

func (ServerMultiOption) OptionType() string { return "multi" }

type ServerServerItemOption struct {
	ServerOptionBase
	Type              string        `json:"type"`
	ItemType          core.ItemType `json:"itemType"`
	InfoType          string        `json:"infoType,omitempty"`
	Def               *core.Sil     `json:"def"`
	AllowOtherServers bool          `json:"allowOtherServers"`
}

func (ServerServerItemOption) OptionType() string { return "serverItem" }

type ServerServerItemsOption struct {
	ServerOptionBase
	Type              string        `json:"type"`
	ItemType          core.ItemType `json:"itemType"`
	InfoType          string        `json:"infoType,omitempty"`
	Def               []core.Sil    `json:"def"`
	AllowOtherServers bool          `json:"allowOtherServers"`
	Limit             int           `json:"limit"`
}

func (ServerServerItemsOption) OptionType() string { return "serverItems" }

type ServerCollectionItemOption struct {
	ServerOptionBase
	Type     string        `json:"type"`
	ItemType core.ItemType `json:"itemType"`
}

func (ServerCollectionItemOption) OptionType() string { return "collectionItem" }

type ServerFileOption struct {
	ServerOptionBase
	Type       string                      `json:"type"`
	Def        string                      `json:"def"`
	Validation *ServerFileOptionValidation `json:"validation,omitempty"`
}

func (ServerFileOption) OptionType() string { return "file" }

type ServerFileOptionValidation struct {
	Type      string   `json:"type"`
	MinSize   *float64 `json:"minSize,omitempty"`
	MaxSize   *float64 `json:"maxSize,omitempty"`
	MinWidth  *float64 `json:"minWidth,omitempty"`
	MaxWidth  *float64 `json:"maxWidth,omitempty"`
	MinHeight *float64 `json:"minHeight,omitempty"`
	MaxHeight *float64 `json:"maxHeight,omitempty"`
	MinLength *float64 `json:"minLength,omitempty"`
	MaxLength *float64 `json:"maxLength,omitempty"`
}

type ServerFileOptionValidationFile = ServerFileOptionValidation

type ServerFileOptionValidationImage = ServerFileOptionValidation

type ServerFileOptionValidationAudio = ServerFileOptionValidation

type ServerFileOptionValidationZip = ServerFileOptionValidation

type ServerFileOptionValidationJson = ServerFileOptionValidation

type ServerInfo struct {
	Title         string              `json:"title"`
	Description   string              `json:"description,omitempty"`
	Buttons       []ServerInfoButton  `json:"buttons"`
	Configuration ServerConfiguration `json:"configuration"`
	Banner        *core.Srl           `json:"banner,omitempty"`
}

type ServerInfoButton interface {
	serverInfoButton()
	ButtonType() string
}

type ServerInfoAuthenticationButton struct {
	Type string `json:"type"`
}

func (ServerInfoAuthenticationButton) serverInfoButton()  {}
func (ServerInfoAuthenticationButton) ButtonType() string { return "authentication" }

type ServerInfoItemButton struct {
	Type       core.ItemType `json:"type"`
	Title      core.Text     `json:"title,omitempty"`
	Icon       core.Icon     `json:"icon,omitempty"`
	BadgeCount *float64      `json:"badgeCount,omitempty"`
	InfoType   string        `json:"infoType,omitempty"`
	ItemName   string        `json:"itemName,omitempty"`
}

func (ServerInfoItemButton) serverInfoButton() {}
func (button ServerInfoItemButton) ButtonType() string {
	return string(button.Type)
}

type ServerInfoConfigurationButton struct {
	Type string `json:"type"`
}

func (ServerInfoConfigurationButton) serverInfoButton()  {}
func (ServerInfoConfigurationButton) ButtonType() string { return "configuration" }

func DecodeServerOption(data []byte) (ServerOption, error) {
	var header unionTypeHeader
	if err := json.Unmarshal(data, &header); err != nil {
		return nil, err
	}

	switch header.Type {
	case "text":
		return decodeServerOptionAs[ServerTextOption](data)
	case "textArea":
		return decodeServerOptionAs[ServerTextAreaOption](data)
	case "slider":
		return decodeServerOptionAs[ServerSliderOption](data)
	case "toggle":
		return decodeServerOptionAs[ServerToggleOption](data)
	case "select":
		return decodeServerOptionAs[ServerSelectOption](data)
	case "multi":
		return decodeServerOptionAs[ServerMultiOption](data)
	case "serverItem":
		return decodeServerOptionAs[ServerServerItemOption](data)
	case "serverItems":
		return decodeServerOptionAs[ServerServerItemsOption](data)
	case "collectionItem":
		return decodeServerOptionAs[ServerCollectionItemOption](data)
	case "file":
		return decodeServerOptionAs[ServerFileOption](data)
	default:
		return nil, coreerrors.UnknownUnionTypeError{
			Union: "ServerOption",
			Type:  header.Type,
		}
	}
}

func decodeServerOptionAs[T ServerOption](data []byte) (ServerOption, error) {
	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return nil, err
	}
	return value, nil
}

func DecodeServerOptions(data []byte) ([]ServerOption, error) {
	var rawOptions []json.RawMessage
	if err := json.Unmarshal(data, &rawOptions); err != nil {
		return nil, err
	}

	options := make([]ServerOption, len(rawOptions))
	for i, rawOption := range rawOptions {
		option, err := DecodeServerOption(rawOption)
		if err != nil {
			return nil, err
		}
		options[i] = option
	}
	return options, nil
}

func DecodeServerInfoButton(data []byte) (ServerInfoButton, error) {
	var header unionTypeHeader
	if err := json.Unmarshal(data, &header); err != nil {
		return nil, err
	}

	switch header.Type {
	case "authentication":
		return decodeServerInfoButtonAs[ServerInfoAuthenticationButton](data)
	case "configuration":
		return decodeServerInfoButtonAs[ServerInfoConfigurationButton](data)
	case string(core.ItemTypePost), string(core.ItemTypePlaylist), string(core.ItemTypeLevel):
		return decodeServerInfoButtonAs[ServerInfoItemButton](data)
	case string(core.ItemTypeSkin), string(core.ItemTypeBackground), string(core.ItemTypeEffect):
		return decodeServerInfoButtonAs[ServerInfoItemButton](data)
	case string(core.ItemTypeParticle), string(core.ItemTypeEngine), string(core.ItemTypeReplay):
		return decodeServerInfoButtonAs[ServerInfoItemButton](data)
	case string(core.ItemTypeRoom), string(core.ItemTypeUser):
		return decodeServerInfoButtonAs[ServerInfoItemButton](data)
	default:
		return nil, coreerrors.UnknownUnionTypeError{
			Union: "ServerInfoButton",
			Type:  header.Type,
		}
	}
}

func decodeServerInfoButtonAs[T ServerInfoButton](data []byte) (ServerInfoButton, error) {
	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return nil, err
	}
	return value, nil
}

func DecodeServerInfoButtons(data []byte) ([]ServerInfoButton, error) {
	var rawButtons []json.RawMessage
	if err := json.Unmarshal(data, &rawButtons); err != nil {
		return nil, err
	}

	buttons := make([]ServerInfoButton, len(rawButtons))
	for i, rawButton := range rawButtons {
		button, err := DecodeServerInfoButton(rawButton)
		if err != nil {
			return nil, err
		}
		buttons[i] = button
	}
	return buttons, nil
}

func (configuration *ServerConfiguration) UnmarshalJSON(data []byte) error {
	type serverConfiguration ServerConfiguration
	var raw struct {
		serverConfiguration
		Options []json.RawMessage `json:"options"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	options := make([]ServerOption, len(raw.Options))
	for i, rawOption := range raw.Options {
		option, err := DecodeServerOption(rawOption)
		if err != nil {
			return err
		}
		options[i] = option
	}

	*configuration = ServerConfiguration(raw.serverConfiguration)
	configuration.Options = options
	return nil
}

func (form *ServerForm) UnmarshalJSON(data []byte) error {
	type serverForm ServerForm
	var raw struct {
		serverForm
		Options []json.RawMessage `json:"options"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	options := make([]ServerOption, len(raw.Options))
	for i, rawOption := range raw.Options {
		option, err := DecodeServerOption(rawOption)
		if err != nil {
			return err
		}
		options[i] = option
	}

	*form = ServerForm(raw.serverForm)
	form.Options = options
	return nil
}

func (info *ServerInfo) UnmarshalJSON(data []byte) error {
	type serverInfo ServerInfo
	var raw struct {
		serverInfo
		Buttons []json.RawMessage `json:"buttons"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	buttons := make([]ServerInfoButton, len(raw.Buttons))
	for i, rawButton := range raw.Buttons {
		button, err := DecodeServerInfoButton(rawButton)
		if err != nil {
			return err
		}
		buttons[i] = button
	}

	*info = ServerInfo(raw.serverInfo)
	info.Buttons = buttons
	return nil
}
