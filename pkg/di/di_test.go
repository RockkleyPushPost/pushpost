package di

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"sync"
	"testing"
)

type testData struct {
	String1 string
	String2 string
	Integer int
	Complex struct{ name string }
}

var td = testData{
	String1: "test string",
	String2: "second string",
	Integer: 123,
	Complex: struct{ name string }{name: "test"},
}

type mockInterface interface {
	DoSomething() string
}

type mockImplementation struct{}

func (m *mockImplementation) DoSomething() string {
	return "mock"
}

type alternativeMockImpl struct{}

func (a *alternativeMockImpl) DoSomething() string {
	return "alternative"
}

func validateDependency(t *testing.T, di *DI, dep interface{}, wantLast bool, index, total int) {
	t.Helper()

	if dep == nil {
		return
	}

	depType := reflect.TypeOf(dep)
	stored, exists := di.dependencies[depType]
	if !exists {
		t.Errorf("Dependency of type %v was not stored", depType)
		return
	}

	if wantLast {
		if index == total-1 && stored != dep {
			t.Errorf("Last stored dependency %v does not match expected %v", stored, dep)
		}
	} else if stored != dep {
		t.Errorf("Stored dependency %v does not match registered dependency %v", stored, dep)
	}
}

func assertDependencyCount(t *testing.T, di *DI, expected int) {
	t.Helper()
	if len(di.dependencies) != expected {
		t.Errorf("Expected %d dependencies, got %d", expected, len(di.dependencies))
	}
}

func generateManyDependencies(count int) []interface{} {
	deps := make([]interface{}, count)
	for i := 0; i < count; i++ {
		switch i % 3 {
		case 0:
			deps[i] = i
		case 1:
			deps[i] = fmt.Sprintf("string-%d", i)
		case 2:
			deps[i] = struct{ value int }{value: i}
		}
	}
	return deps
}

func TestDI_Register(t *testing.T) {
	var mu sync.Mutex

	tests := []struct {
		name         string
		description  string
		dependencies []interface{}
		wantErr      bool
		wantLast     bool
		wantCount    int
	}{
		{
			name:         "register single dependency",
			description:  "Should successfully register a single string dependency",
			dependencies: []interface{}{td.String1},
			wantErr:      false,
			wantCount:    1,
		},
		{
			name:         "register multiple dependencies",
			description:  "Should handle multiple dependencies of different types",
			dependencies: []interface{}{td.Integer, td.String1, true},
			wantErr:      false,
			wantCount:    3,
		},
		{
			name:         "register nil dependency",
			description:  "Should return error when registering nil dependency",
			dependencies: []interface{}{nil},
			wantErr:      true,
			wantCount:    0,
		},
		{
			name:         "register empty dependencies",
			description:  "Should handle empty dependency list",
			dependencies: []interface{}{},
			wantErr:      false,
			wantCount:    0,
		},
		{
			name:         "register same type multiple times",
			description:  "Should keep only the last value when registering same type",
			dependencies: []interface{}{td.String1, td.String2},
			wantErr:      false,
			wantLast:     true,
			wantCount:    1,
		},
		{
			name:         "register complex struct",
			description:  "Should handle complex struct types",
			dependencies: []interface{}{td.Complex},
			wantErr:      false,
			wantCount:    1,
		},
		{
			name:         "register many dependencies",
			description:  "Should handle multiple different types of dependencies",
			dependencies: generateManyDependencies(15),
			wantErr:      false,
			wantLast:     true,
			wantCount:    3,
		},
		{
			name:         "register interface implementation",
			description:  "Should handle interface implementations",
			dependencies: []interface{}{&mockImplementation{}},
			wantErr:      false,
			wantCount:    1,
		},
		{
			name:         "register duplicate interface implementations",
			description:  "Should handle multiple implementations of same interface",
			dependencies: []interface{}{&mockImplementation{}, &alternativeMockImpl{}},
			wantErr:      false,
			wantLast:     true,
			wantCount:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			di := NewDI(&fiber.App{})
			t.Cleanup(func() {
				di.dependencies = nil
			})

			mu.Lock()
			err := di.Register(tt.dependencies...)
			mu.Unlock()

			if (err != nil) != tt.wantErr {
				t.Errorf("DI.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				mu.Lock()
				assertDependencyCount(t, di, tt.wantCount)
				for i, dep := range tt.dependencies {
					validateDependency(t, di, dep, tt.wantLast, i, len(tt.dependencies))
				}
				mu.Unlock()
			}
		})
	}
}

func BenchmarkDI_Register(b *testing.B) {
	di := NewDI(&fiber.App{})
	deps := generateManyDependencies(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		di.Register(deps...)
	}
}
func ExampleDI_Register() {
	di := NewDI(&fiber.App{})
	err := di.Register("example", 42)
	if err != nil {
		fmt.Println(err)
	}

}
