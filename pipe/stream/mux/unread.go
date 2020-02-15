package mux

import (
	"io"

	"github.com/wzshiming/pipe/pipe/stream"
)

func UnwrapUnreadStream(rwc stream.Stream) (stream.Stream, []byte) {
	if us, ok := rwc.(*unreadStream); ok {
		_, prefix := UnwrapUnread(us.Reader)
		return us.Stream, prefix
	}
	return rwc, nil
}

func UnreadStream(rwc stream.Stream, prefix []byte) stream.Stream {
	if len(prefix) == 0 {
		return rwc
	}
	if us, ok := rwc.(*unreadStream); ok {
		us.Reader = Unread(us.Reader, prefix)
		return us
	}
	return &unreadStream{
		Reader: Unread(rwc, prefix),
		Stream: rwc,
	}
}

type unreadStream struct {
	io.Reader
	stream.Stream
}

func (c *unreadStream) Read(p []byte) (n int, err error) {
	return c.Reader.Read(p)
}

func UnwrapUnread(reader io.Reader) (io.Reader, []byte) {
	if u, ok := reader.(*unread); ok {
		return u.reader, u.prefix
	}
	return reader, nil
}

func Unread(reader io.Reader, prefix []byte) io.Reader {
	if len(prefix) == 0 {
		return reader
	}
	if ur, ok := reader.(*unread); ok {
		ur.prefix = append(prefix, ur.prefix...)
		return reader
	}
	return &unread{
		prefix: prefix,
		reader: reader,
	}
}

type unread struct {
	prefix []byte
	reader io.Reader
}

func (u *unread) Read(p []byte) (n int, err error) {
	if len(u.prefix) == 0 {
		return u.reader.Read(p)
	}
	n = copy(p, u.prefix)
	if n <= len(u.prefix) {
		u.prefix = u.prefix[n:]
		return n, nil
	}
	a, err := u.reader.Read(p[n:])
	if err == io.EOF {
		err = nil
	}
	n += a
	return n, err
}
