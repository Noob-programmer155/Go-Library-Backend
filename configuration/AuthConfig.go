package configuration

import "net/http"

func CheckAuth(h *http.Request, role int) (auth bool, err error) {
	cookie, err := h.Cookie("JLMS_TOKEN")
	if cookie != nil {
		if claims, err := ParseToken(cookie.Value); err == nil {
			if claims.Role == role {
				return true, nil
			}
			return false, err
		}
		return false, err
	}
	return false, err
}

func GetAuth() {

}
