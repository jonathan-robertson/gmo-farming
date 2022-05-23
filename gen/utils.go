package gen

import (
	"fmt"
	"os"
)

// Write is used to set up the underlying file that the given producer will produce strings to
func Write(producer Producer) error {
	file, err := getFile(producer.GetPath(), producer.GetFilename())
	if err != nil {
		return err
	}
	defer file.Close()
	c := make(chan string, 10)
	go producer.Produce(c)
	for line := range c {
		if _, err = file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func getFile(path, filename string) (*os.File, error) {
	pathAndFilename := fmt.Sprintf("%s%c%s", path, os.PathSeparator, filename)
	os.Remove(pathAndFilename) // ignore error; deleting if present only
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, err
	}
	return os.OpenFile(pathAndFilename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
}

func getEnhancedSeedEffectDescription() string {
	return `All Enhanced Seeds yield twice the standard number of crops.`
}

func getInitialEnhancementCraftingTip() string {
	return `This can be crafted by Hand or at a Workbench.`
}

func getHotBoxRequirementTip() string {
	return `A Hot Box is required to craft this.`
}
