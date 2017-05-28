package userdb

import (
	"crypto"
	"fmt"

	"github.com/DistributedSolutions/twofactor"
)

func (u *User) Create2FA() ([]byte, error) {
	if u.Has2FA {
		return nil, fmt.Errorf("2FA is already enabled")
	}

	tfa, err := twofactor.NewTOTP(u.Username, "HodlZone", crypto.SHA256, 8)
	if err != nil {
		return nil, err
	}

	u.User2FA = tfa
	return u.User2FA.QR()
}

// Validate2FA will return if succusful, but also modify the underlying struct.
// Remember to save the user back to disk after calling
func (u *User) Validate2FA(token string) error {
	return u.User2FA.Validate(token)
}

// func test() {
// 	otp, err := twofactor.NewTOTP("info@sec51.com", "Sec51", crypto.SHA256, 8)
// 	if err != nil {
// 		return err
// 	}

// 	qrBytes, err := otp.QR()
// 	if err != nil {
// 		return err
// 	}

// 	err := otp.Validate(USER_PROVIDED_TOKEN)
// 	if err != nil {
// 		return err
// 	}

// }