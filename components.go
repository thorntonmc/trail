package trail

import (
	"bytes"
)

type state int64

const (
	startDir state = iota
	body
	done
)

type Components struct {
	hasPhysicalRoot bool
	trail           []byte
	front           state
	back            state
}

func (c *Components) Next() []byte {
	switch c.front {
	case startDir:
		c.front = body // move up one state
		if c.hasPhysicalRoot {
			c.trail = c.trail[1:]
			return []byte{rootDir}
		}
		return c.parseNextComponent()
	case body:
		return c.parseNextComponent()
	case done:
		return []byte{}
	}

	return []byte{}
}

func (c *Components) parseNextComponent() []byte {
	index := bytes.Index(c.trail, []byte(string(separator)))
	if index < 0 {
		c.front = done
		return c.trail
	}

	component := c.trail[:index]
	c.trail = c.trail[index+1:]
	return component

}

func (c *Components) finished() bool {
	return c.front == done || c.back == done || c.front > c.back
}

func isSepByte(b byte) bool {
	return b == '/'
}
