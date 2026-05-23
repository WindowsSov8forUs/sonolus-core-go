package server

import "github.com/WindowsSov8forUs/sonolus-core-go/core"

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
