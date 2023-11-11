package prompt

import (
	"errors"
	"github.com/manifoldco/promptui"
	"strconv"
	"sync"
	"sync/atomic"
)

// Used to control the access to a prompt method across Goroutines
var mutex sync.Mutex

// Used to share input/selected option across Goroutines
var promptsValues atomic.Value

type Prompt struct{}

// ReadInt Read input as int from a prompt
// and cache it for future use in other Goroutines (if needed)
func (p *Prompt) ReadInt(label string) interface{} {

	mutex.Lock()
	defer mutex.Unlock()

	if value, ok := readInputFromAtomic[int64](label); ok {
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
		panic(err)
	}

	number, _ := strconv.ParseInt(result, 10, 64)

	return loadInputInAtomic[int64](label, number)

}

// ReadString Read input as string from a prompt input field (text)
// and cache it for future use in other Goroutines (if needed)
func (p *Prompt) ReadString(label string) interface{} {

	mutex.Lock()
	defer mutex.Unlock()

	if value, ok := readInputFromAtomic[string](label); ok {
		return value
	}

	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	return loadInputInAtomic[string](label, result)

}

// Select Read input as string from a prompt input field (text)
// and cache it for future use in other Goroutines (if needed).
func (p *Prompt) Select(label string, options ...string) interface{} {
	mutex.Lock()
	defer mutex.Unlock()

	if value, ok := readInputFromAtomic[string](label); ok {
		return value
	}

	prompt := promptui.Select{
		Label: label,
		Items: options,
	}

	_, result, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	return loadInputInAtomic[string](label, result)

}

// ReadInputFromAtomic read an input from an atomic map written by other Goroutines
func readInputFromAtomic[T interface{}](label string) (T, bool) {
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
func loadInputInAtomic[T string | int64 | float32](label string, value T) T {
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
