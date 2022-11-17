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

	values := promptsValues.Load()
	if values == nil {
		executions := make(map[string]interface{}, 100)
		promptsValues.Store(executions)
	}

	executions := promptsValues.Load().(map[string]interface{})

	if executions[label] != nil {
		return executions[label]
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

	executions[label] = result

	promptsValues.Store(executions)

	return number

}

// ReadString Read input as string
func (p *Prompt) ReadString(label string) interface{} {

	mutex.Lock()
	defer mutex.Unlock()

	values := promptsValues.Load()
	if values == nil {
		executions := make(map[string]interface{}, 100)
		promptsValues.Store(executions)
	}

	executions := promptsValues.Load().(map[string]interface{})

	if executions[label] != nil {
		return executions[label]
	}

	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	executions[label] = result

	promptsValues.Store(executions)

	return result

}

// Select read input as string from a select
func (p *Prompt) Select(label string, options ...string) interface{} {
	mutex.Lock()
	defer mutex.Unlock()

	values := promptsValues.Load()
	if values == nil {
		executions := make(map[string]interface{}, 100)
		promptsValues.Store(executions)
	}

	executions := promptsValues.Load().(map[string]interface{})

	if executions[label] != nil {
		return executions[label]
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

	executions[label] = result

	promptsValues.Store(executions)

	return result

}
