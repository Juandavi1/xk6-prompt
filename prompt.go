package prompt

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"sync/atomic"
)

// Used to control the access to a prompt method across Goroutines
var mutex sync.Mutex

// Used to share input/selected option across Goroutines
var promptsValues atomic.Value

type Prompt struct{}

// ReadInt Read input as Int
func (p *Prompt) ReadInt(label string) interface{} {

	mutex.Lock()
	defer mutex.Unlock()

	if value, ok := ReadInputFromAtomic[int64](label); ok {
		return value
	}

	validate := func(input string) error {
		_, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return errors.New("invalid number format")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return 0
	}

	number, _ := strconv.ParseInt(result, 10, 64)

	return LoadInputInAtomic[int64](label, number)

}

// ReadString Read input as string
func (p *Prompt) ReadString(label string) interface{} {

	mutex.Lock()
	defer mutex.Unlock()

	if value, ok := ReadInputFromAtomic[string](label); ok {
		return value
	}

	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return LoadInputInAtomic[string](label, result)

}

// Select read input as string from a select
func (p *Prompt) Select(label string, options ...string) interface{} {
	mutex.Lock()
	defer mutex.Unlock()

	if value, ok := ReadInputFromAtomic[string](label); ok {
		return value
	}

	prompt := promptui.Select{
		Label: label,
		Items: options,
	}

	_, result, err := prompt.Run()

	if err != nil {
		logrus.Errorf("Prompt failed %v\n", err)
		return ""
	}

	return LoadInputInAtomic[string](label, result)

}

// ReadInputFromAtomic read an input entered from a cached list
func ReadInputFromAtomic[T interface{}](label string) (T, bool) {
	values := promptsValues.Load()
	if values == nil {
		inputsCached := make(map[string]interface{}, 100)
		promptsValues.Store(inputsCached)
	}

	inputsCached := promptsValues.Load().(map[string]interface{})

	if value, found := inputsCached[label]; found {
		return value.(T), true
	}

	return *new(T), false
}

// LoadInputInAtomic write an input entered into an atomic map to be read by other Goroutines
func LoadInputInAtomic[T interface{}](label string, value T) T {
	values := promptsValues.Load()
	if values == nil {
		inputsCached := make(map[string]interface{}, 100)
		promptsValues.Store(inputsCached)
	}

	inputsCached := promptsValues.Load().(map[string]interface{})

	inputsCached[label] = value

	promptsValues.Store(inputsCached)

	return value
}
