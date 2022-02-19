# trail 

Inspired by rust's [std::path](https://doc.rust-lang.org/stable/std/path/struct.Path.html)

## Components

Break a path into components:
```go
import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestComponentRead(t *testing.T) {
    var c *Components
    // with root
    c, _ = testTrail.Components()
    assert.Equal(t, string(rootDir), string(c.Next()))
    assert.Equal(t, "Users", string(c.Next()))
    assert.Equal(t, "name", string(c.Next()))
    assert.Equal(t, "personal", string(c.Next()))
    assert.Equal(t, "trail", string(c.Next()))
    assert.Equal(t, "", string(c.Next()))
}
```
## Ancestry

Break a path into its ancestors
```go
func TestAncestor(t *testing.T) {
    trail := testTrail
    a := trail.Ancestors()

    assert.Equal(t, "/Users/name/personal/trail", a.Next().Inner)
    assert.Equal(t, "/Users/name/personal", a.Next().Inner)
    assert.Equal(t, "/Users/name", a.Next().Inner)
    assert.Equal(t, "/Users", a.Next().Inner)
    assert.Nil(t, a.Next())
    assert.Nil(t, a.Next())
}
```
