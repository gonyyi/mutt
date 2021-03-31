// (C) 2021 GON Y YI.
// https://gonyyi.com/copyright.txt

package mutt 

import (
	"encoding/json"
	"github.com/gonyyi/atype"
	"github.com/gonyyi/mutt"
	"time"
)

func NewUser(b []byte) (*User, error) {
	u := User{}
	err := u.Load(b)

	if u.Created < 1 {
		u.Created = atype.TimeInt(time.Now().Unix())
	}
	if u.OtherTokens == nil {
		u.OtherTokens = atype.NewMapStr()
	}
	if u.Groups == nil {
		u.Groups = atype.NewMapStr()
	}
	return &u, err
}

type User struct {
	ID           string         `json:"id"`
	Credential   UserCredential `json:"credential"`
	Name         UserName       `json:"name"`
	Email        UserEmail      `json:"email"`
	Enabled      bool           `json:"enabled"`
	Events       []UserEvent    `json:"events"`
	OtherTokens  atype.MapStr   `json:"otherTokens"` // OPTIONAL: user's token once logged in
	Groups       atype.MapStr   `json:"groups"`      // OPTIONAL: user's token once logged in
	Created      atype.TimeInt  `json:"created"`
	LastModified atype.TimeInt  `json:"lastModified"`
	LastLogin    atype.TimeInt  `json:"lastLogin"`
}

func (u *User) Reset() {
	u.ID = ""
	u.Name.Reset()
	u.Enabled = false

	if u.OtherTokens == nil {
		u.OtherTokens.Reset()
	} else {
		u.OtherTokens = atype.NewMapStr()
	}

	if u.Groups == nil {
		u.Groups.Reset()
	} else {
		u.Groups = atype.NewMapStr()
	}
}

func (u *User) Bytes() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) Load(d []byte) error {
	u.Reset()
	return json.Unmarshal(d, u)
}

type UserName struct {
	DisplayName string `json:"disp"`
	FirstName   string `json:"fn"`
	LastName    string `json:"ln"`
}

func (u *UserName) Reset() {
	u.DisplayName = ""
	u.FirstName = ""
	u.LastName = ""
}
func (u *UserName) Set(DispN, Fn, Ln string) {
	u.DisplayName = DispN
	u.FirstName = Fn
	u.LastName = Ln
}

type UserEvent struct {
	Time    atype.TimeInt `json:"time"`
	Message string        `json:"message"`
}

type UserCredential struct {
	Passwd string `json:"passwd"` // hashed/encrypted user's password
	Token  string `json:"token"`
}

func (c *UserCredential) SetPasswd(pwd string) error {
	b, err := mutt.PasswdHash([]byte(pwd))
	c.Passwd = string(b)
	return err
}

func (c *UserCredential) VerifyPasswd(pwd string) error {
	return mutt.PasswdCompare([]byte(c.Passwd), []byte(pwd))
}

func (c *UserCredential) NewToken() {
	c.Token = mutt.Random(32)
}

type UserEmail struct {
	Email      string        `json:"email"`
	Verified   bool          `json:"verified"`
	VerifiedOn atype.TimeInt `json:"verifiedOn"`
}

func (u *UserEmail) Reset() {
	u.Email = ""
	u.Verified = false
	u.VerifiedOn = 0
}

func (u *UserEmail) Verify(verified bool) {
	if verified {
		u.Verified = true
		u.VerifiedOn.Set() // Set current time
		return
	}
	u.Verified = false
	u.VerifiedOn = 0
}
