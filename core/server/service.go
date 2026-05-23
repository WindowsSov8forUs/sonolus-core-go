package server

type ServiceUserId string

type ServiceUserProfile struct {
	ID                    ServiceUserId `json:"id"`
	Handle                string        `json:"handle"`
	Name                  string        `json:"name"`
	AvatarType            string        `json:"avatarType"`
	AvatarForegroundType  string        `json:"avatarForegroundType"`
	AvatarForegroundColor string        `json:"avatarForegroundColor"`
	AvatarBackgroundType  string        `json:"avatarBackgroundType"`
	AvatarBackgroundColor string        `json:"avatarBackgroundColor"`
	BannerType            string        `json:"bannerType"`
	AboutMe               string        `json:"aboutMe"`
	Favorites             []string      `json:"favorites"`
}

type ServiceAuthenticateExternalRequest struct {
	Type        string             `json:"type"`
	URL         string             `json:"url"`
	Time        float64            `json:"time"`
	UserProfile ServiceUserProfile `json:"userProfile"`
}

type ServiceAuthenticateExternalResponse = ServerMessage

type SignaturePublicKeyJWK struct {
	KeyOps []string `json:"key_ops"`
	Ext    bool     `json:"ext"`
	Kty    string   `json:"kty"`
	X      string   `json:"x"`
	Y      string   `json:"y"`
	Crv    string   `json:"crv"`
}

func SignaturePublicKey() SignaturePublicKeyJWK {
	return SignaturePublicKeyJWK{
		KeyOps: []string{"verify"},
		Ext:    true,
		Kty:    "EC",
		X:      "d2B14ZAn-zDsqY42rHofst8rw3XB90-a5lT80NFdXo0",
		Y:      "Hxzi9DHrlJ4CVSJVRnydxFWBZAgkFxZXbyxPSa8SJQw",
		Crv:    "P-256",
	}
}
