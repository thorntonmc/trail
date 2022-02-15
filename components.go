package trail

import (
	"bytes"
	"fmt"
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
	}

	return []byte{}
}

func (c *Components) parseNextComponent() []byte {
	index := bytes.Index(c.trail, []byte(string(separator)))

	fmt.Println(index)
	if index < 0 {
		return c.trail
	}

	fmt.Println(index)
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
