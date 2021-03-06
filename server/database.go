package server

import (
	"errors"
	"time"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

const (
	DATE_FORMAT = "20060102"
)

var (
	scopes    = []string{endpoints.EmailScope}
	clientIds = []string{keys.WebClientId, keys.AndroidReleaseClientId, keys.AndroidDebugClientId, endpoints.APIExplorerClientID}
	audiences = []string{keys.WebClientId, keys.AndroidReleaseClientId, keys.AndroidDebugClientId}
)

type DeviceType int

const (
	// Source Enum
	UNDEFINED DeviceType = 0
	ANDROID              = 1
	PC                   = 2
)

type User struct {
	UserId string `datastore:"-"` // User.ID
}

type Device struct {
	UserId         string `datastore:"-"` // User.ID
	DeviceId       string `datastore:"-"` // Unique ID for a device
	DeviceName     string // Display name.
	AlertThreshold int32  // 0 - 100.
	DeviceType     DeviceType
	Disabled       bool

	// For API
	Batteries []Battery `database:"-"`
}

type Battery struct {
	Time     time.Time `json:"time"`    // timestamp
	Battery  int32     `json:"battery"` // 0 - 100.
	Charging bool      `json:"charging"`
}

type History struct {
	Batteries []Battery
}

type ByTime []Battery

func (b ByTime) Len() int {
	return len(b)
}

func (b ByTime) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByTime) Less(i, j int) bool {
	return b[i].Time.Before(b[j].Time)
}

func userKey(u *User, c context.Context) *datastore.Key {
	uk := datastore.NewKey(c, "User", u.UserId, 0, nil)
	log.Debugf(c, "%#v", uk)
	return uk
}

func deviceKey(u *User, d string, c context.Context) *datastore.Key {
	uk := userKey(u, c)
	dk := datastore.NewKey(c, "Device", d, 0, uk)
	log.Debugf(c, "%#v", dk)
	return dk
}

func historyKey(u *User, d string, t time.Time, c context.Context) *datastore.Key {
	dk := deviceKey(u, d, c)
	hk := datastore.NewKey(c, "History", toDate(t), 0, dk)
	log.Debugf(c, "%#v", hk)
	return hk
}

func getHistory(key *datastore.Key, c context.Context) (*History, error) {
	h := new(History)
	err := datastore.Get(c, key, h)
	if err == datastore.ErrNoSuchEntity {
		err = nil
	}
	return h, err
}

// toDate is the StringID of historyKey.
func toDate(t time.Time) string {
	return t.UTC().Format(DATE_FORMAT)
}

func populateKey(k *datastore.Key, b *Battery) {
	b.Time = time.Unix(k.IntID(), 0)
}

func populateDeviceId(k *datastore.Key, d *Device) {
	d.DeviceId = k.StringID()
}

// getCurrentUser retrieves a user associated with the request.
// If there's no user (e.g. no auth info present in the request) returns
// an "unauthorized" error.
func getCurrentUser(c context.Context) (*User, error) {
	u, err := endpoints.CurrentUser(c, scopes, audiences, clientIds)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("Unauthorized: Please, sign in.")
	}
	log.Debugf(c, "Current user: %#v", u)
	return toUser(u), nil
}

func toUser(u *user.User) *User {
	if u == nil {
		return nil
	}
	return &User{UserId: u.Email}
}
