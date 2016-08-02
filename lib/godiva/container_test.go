package godiva

import (
	"testing"
)

type TestTarget struct {
	TestProp1 int
	TestProp2 string
	TestProp3 TestChildObject
}

type TestChildObject struct {
}

// Test can be executed
func TestRegister(t *testing.T) {
	container := CreateContainer()
	container.Register(
		"testChildObj",
		func(c *Container) (interface{}, error) {
			return TestChildObject{}, nil
		},
	)
	container.Register(
		"testObj",
		func(c *Container) (interface{}, error) {
			childObj, _ := c.Create("testChildObj", false)
			t := TestTarget{
				1,
				"test",
				childObj.(TestChildObject),
			}
			return t, nil
		},
	)
}

// That the resources that have registered can be generated
func TestCreate(t *testing.T) {
	expected := 2
	container := CreateContainer()
	container.Register(
		"testChildObj",
		func(c *Container) (interface{}, error) {
			return TestChildObject{}, nil
		},
	)
	container.Register(
		"testObj",
		func(c *Container) (interface{}, error) {
			childObj, _ := c.Create("testChildObj", false)
			t := TestTarget{
				2,
				"test",
				childObj.(TestChildObject),
			}
			return t, nil
		},
	)
	testObject, err := container.Create("testObj", false)
	actual := testObject.(TestTarget).TestProp1
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if actual != expected {
		t.Errorf("actual: %v\n expected: %v", actual, expected)
		return
	}
}

// Use the cache at the time of generation
func testCreateByCache(t *testing.T) {
	container := CreateContainer()
	container.Register(
		"testChildObj",
		func(c *Container) (interface{}, error) {
			return TestChildObject{}, nil
		},
	)
	container.Register(
		"testObj",
		func(c *Container) (interface{}, error) {
			childObj, _ := c.Create("testChildObj", false)
			t := TestTarget{
				1,
				"cached",
				childObj.(TestChildObject),
			}
			return t, nil
		},
	)

	_, _ = container.Create("testObj", true)

	container.Register(
		"testObj",
		func(c *Container) (interface{}, error) {
			childObj, _ := c.Create("testChildObj", false)
			t := TestTarget{
				1,
				"nocached",
				childObj.(TestChildObject),
			}
			return t, nil
		},
	)

	obj, _ := container.Create("testObj", false)
	if obj.(TestTarget).TestProp2 != "cached" {
		t.Errorf("not moving normally cache function")
	}
}

// Ignoring the cache at the time of generation
func testCreateIgnoreCache(t *testing.T) {
	container := CreateContainer()
	container.Register(
		"testChildObj",
		func(c *Container) (interface{}, error) {
			return TestChildObject{}, nil
		},
	)
	container.Register(
		"testObj",
		func(c *Container) (interface{}, error) {
			childObj, _ := c.Create("testChildObj", false)
			t := TestTarget{
				1,
				"cached",
				childObj.(TestChildObject),
			}
			return t, nil
		},
	)

	_, _ = container.Create("testObj", true)

	container.Register(
		"testObj",
		func(c *Container) (interface{}, error) {
			childObj, _ := c.Create("testChildObj", false)
			t := TestTarget{
				1,
				"nocached",
				childObj.(TestChildObject),
			}
			return t, nil
		},
	)
	obj, _ := container.Create("testObj", true)
	if obj.(TestTarget).TestProp2 != "nocached" {
		t.Errorf("not moving normally ignore cache function")
	}
}

// When unregistered resource generation should return an error
func testCreateUnregisteredResource(t *testing.T) {
	container := CreateContainer()
	container.Register(
		"testChildObj",
		func(c *Container) (interface{}, error) {
			return TestChildObject{}, nil
		},
	)
	_, err := container.Create("unRegistered", true)
	if err == nil {
		t.Errorf("When unregistered resource generation should return an error")
	}
}
