// Package xtest is the X client API for the XTEST extension.
package xtest

// This file is automatically generated from xtest.xml. Edit at your peril!

import (
	"github.com/robotn/xgb"

	"github.com/robotn/xgb/xproto"
)

// Init must be called before using the XTEST extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 5, "XTEST").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named XTEST could be found on on the server.")
	}

	c.ExtLock.Lock()
	c.Extensions["XTEST"] = reply.MajorOpcode
	c.ExtLock.Unlock()
	for evNum, fun := range xgb.NewExtEventFuncs["XTEST"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["XTEST"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	return nil
}

func init() {
	xgb.NewExtEventFuncs["XTEST"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["XTEST"] = make(map[int]xgb.NewErrorFun)
}

const (
	CursorNone    = 0
	CursorCurrent = 1
)

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Card32'

// CompareCursorCookie is a cookie used only for CompareCursor requests.
type CompareCursorCookie struct {
	*xgb.Cookie
}

// CompareCursor sends a checked request.
// If an error occurs, it will be returned with the reply by calling CompareCursorCookie.Reply()
func CompareCursor(c *xgb.Conn, Window xproto.Window, Cursor xproto.Cursor) CompareCursorCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XTEST"]; !ok {
		panic("Cannot issue request 'CompareCursor' using the uninitialized extension 'XTEST'. xtest.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(compareCursorRequest(c, Window, Cursor), cookie)
	return CompareCursorCookie{cookie}
}

// CompareCursorUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func CompareCursorUnchecked(c *xgb.Conn, Window xproto.Window, Cursor xproto.Cursor) CompareCursorCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XTEST"]; !ok {
		panic("Cannot issue request 'CompareCursor' using the uninitialized extension 'XTEST'. xtest.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(compareCursorRequest(c, Window, Cursor), cookie)
	return CompareCursorCookie{cookie}
}

// CompareCursorReply represents the data returned from a CompareCursor request.
type CompareCursorReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	Same     bool
}

// Reply blocks and returns the reply data for a CompareCursor request.
func (cook CompareCursorCookie) Reply() (*CompareCursorReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return compareCursorReply(buf), nil
}

// compareCursorReply reads a byte slice into a CompareCursorReply value.
func compareCursorReply(buf []byte) *CompareCursorReply {
	v := new(CompareCursorReply)
	b := 1 // skip reply determinant

	if buf[b] == 1 {
		v.Same = true
	} else {
		v.Same = false
	}
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	return v
}

// Write request to wire for CompareCursor
// compareCursorRequest writes a CompareCursor request to a byte slice.
func compareCursorRequest(c *xgb.Conn, Window xproto.Window, Cursor xproto.Cursor) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["XTEST"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Window))
	b += 4

	xgb.Put32(buf[b:], uint32(Cursor))
	b += 4

	return buf
}

// FakeInputCookie is a cookie used only for FakeInput requests.
type FakeInputCookie struct {
	*xgb.Cookie
}

// FakeInput sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func FakeInput(c *xgb.Conn, Type byte, Detail byte, Time uint32, Root xproto.Window, RootX int16, RootY int16, Deviceid byte) FakeInputCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XTEST"]; !ok {
		panic("Cannot issue request 'FakeInput' using the uninitialized extension 'XTEST'. xtest.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(fakeInputRequest(c, Type, Detail, Time, Root, RootX, RootY, Deviceid), cookie)
	return FakeInputCookie{cookie}
}

// FakeInputChecked sends a checked request.
// If an error occurs, it can be retrieved using FakeInputCookie.Check()
func FakeInputChecked(c *xgb.Conn, Type byte, Detail byte, Time uint32, Root xproto.Window, RootX int16, RootY int16, Deviceid byte) FakeInputCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XTEST"]; !ok {
		panic("Cannot issue request 'FakeInput' using the uninitialized extension 'XTEST'. xtest.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(fakeInputRequest(c, Type, Detail, Time, Root, RootX, RootY, Deviceid), cookie)
	return FakeInputCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook FakeInputCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for FakeInput
// fakeInputRequest writes a FakeInput request to a byte slice.
func fakeInputRequest(c *xgb.Conn, Type byte, Detail byte, Time uint32, Root xproto.Window, RootX int16, RootY int16, Deviceid byte) []byte {
	size := 36
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["XTEST"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Type
	b += 1

	buf[b] = Detail
	b += 1

	b += 2 // padding

	xgb.Put32(buf[b:], Time)
	b += 4

	xgb.Put32(buf[b:], uint32(Root))
	b += 4

	b += 8 // padding

	xgb.Put16(buf[b:], uint16(RootX))
	b += 2

	xgb.Put16(buf[b:], uint16(RootY))
	b += 2

	b += 7 // padding

	buf[b] = Deviceid
	b += 1

	return buf
}

// GetVersionCookie is a cookie used only for GetVersion requests.
type GetVersionCookie struct {
	*xgb.Cookie
}

// GetVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling GetVersionCookie.Reply()
func GetVersion(c *xgb.Conn, MajorVersion byte, MinorVersion uint16) GetVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XTEST"]; !ok {
		panic("Cannot issue request 'GetVersion' using the uninitialized extension 'XTEST'. xtest.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(getVersionRequest(c, MajorVersion, MinorVersion), cookie)
	return GetVersionCookie{cookie}
}

// GetVersionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func GetVersionUnchecked(c *xgb.Conn, MajorVersion byte, MinorVersion uint16) GetVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XTEST"]; !ok {
		panic("Cannot issue request 'GetVersion' using the uninitialized extension 'XTEST'. xtest.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(getVersionRequest(c, MajorVersion, MinorVersion), cookie)
	return GetVersionCookie{cookie}
}

// GetVersionReply represents the data returned from a GetVersion request.
type GetVersionReply struct {
	Sequence     uint16 // sequence number of the request for this reply
	Length       uint32 // number of bytes in this reply
	MajorVersion byte
	MinorVersion uint16
}

// Reply blocks and returns the reply data for a GetVersion request.
func (cook GetVersionCookie) Reply() (*GetVersionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getVersionReply(buf), nil
}

// getVersionReply reads a byte slice into a GetVersionReply value.
func getVersionReply(buf []byte) *GetVersionReply {
	v := new(GetVersionReply)
	b := 1 // skip reply determinant

	v.MajorVersion = buf[b]
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.MinorVersion = xgb.Get16(buf[b:])
	b += 2

	return v
}

// Write request to wire for GetVersion
// getVersionRequest writes a GetVersion request to a byte slice.
func getVersionRequest(c *xgb.Conn, MajorVersion byte, MinorVersion uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["XTEST"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = MajorVersion
	b += 1

	b += 1 // padding

	xgb.Put16(buf[b:], MinorVersion)
	b += 2

	return buf
}

// GrabControlCookie is a cookie used only for GrabControl requests.
type GrabControlCookie struct {
	*xgb.Cookie
}

// GrabControl sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func GrabControl(c *xgb.Conn, Impervious bool) GrabControlCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XTEST"]; !ok {
		panic("Cannot issue request 'GrabControl' using the uninitialized extension 'XTEST'. xtest.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(grabControlRequest(c, Impervious), cookie)
	return GrabControlCookie{cookie}
}

// GrabControlChecked sends a checked request.
// If an error occurs, it can be retrieved using GrabControlCookie.Check()
func GrabControlChecked(c *xgb.Conn, Impervious bool) GrabControlCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XTEST"]; !ok {
		panic("Cannot issue request 'GrabControl' using the uninitialized extension 'XTEST'. xtest.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(grabControlRequest(c, Impervious), cookie)
	return GrabControlCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook GrabControlCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for GrabControl
// grabControlRequest writes a GrabControl request to a byte slice.
func grabControlRequest(c *xgb.Conn, Impervious bool) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["XTEST"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	if Impervious {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}
