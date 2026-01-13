package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PromptBool 提示用户输入 yes/no
func PromptBool(question string, defaultValue bool) bool {
	reader := bufio.NewReader(os.Stdin)

	defaultStr := "y/N"
	if defaultValue {
		defaultStr = "Y/n"
	}

	fmt.Printf("%s [%s]: ", question, defaultStr)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "" {
		return defaultValue
	}

	return input == "y" || input == "yes"
}

// PromptString 提示用户输入字符串
func PromptString(question string, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)

	if defaultValue != "" {
		fmt.Printf("%s [%s]: ", question, defaultValue)
	} else {
		fmt.Printf("%s: ", question)
	}

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		return defaultValue
	}

	return input
}

// PromptSelect 提示用户从选项中选择
func PromptSelect(question string, options []string, defaultIndex int) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(question)
		for i, option := range options {
			defaultMarker := " "
			if i == defaultIndex {
				defaultMarker = "*"
			}
			fmt.Printf("  %s %d. %s\n", defaultMarker, i+1, option)
		}

		fmt.Printf("请选择 [1-%d]: ", len(options))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			return defaultIndex
		}

		var choice int
		if _, err := fmt.Sscanf(input, "%d", &choice); err == nil {
			if choice >= 1 && choice <= len(options) {
				return choice - 1
			}
		}

		fmt.Println("❌ 无效的选择，请重试")
	}
}

// PromptMultiSelect 提示用户多选
func PromptMultiSelect(question string, options []string) []string {
	reader := bufio.NewReader(os.Stdin)
	selected := []int{}

	fmt.Println(question)
	fmt.Println("(输入选项编号，用逗号分隔，完成后按回车)")

	for {
		for i, option := range options {
			selectedMarker := " "
			for _, s := range selected {
				if s == i {
					selectedMarker = "x"
					break
				}
			}
			fmt.Printf("  [%s] %d. %s\n", selectedMarker, i+1, option)
		}

		fmt.Print("选择 (或直接回车完成): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			break
		}

		choices := strings.Split(input, ",")
		for _, choiceStr := range choices {
			var choice int
			if _, err := fmt.Sscanf(strings.TrimSpace(choiceStr), "%d", &choice); err == nil {
				if choice >= 1 && choice <= len(options) {
					idx := choice - 1
					// 切换选择状态
					found := false
					for i, s := range selected {
						if s == idx {
							selected = append(selected[:i], selected[i+1:]...)
							found = true
							break
						}
					}
					if !found {
						selected = append(selected, idx)
					}
				}
			}
		}
	}

	result := make([]string, len(selected))
	for i, idx := range selected {
		result[i] = options[idx]
	}

	return result
}
