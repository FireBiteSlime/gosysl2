package mysort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySort(t *testing.T) {

	test_path := []string{
		`test_data.txt`,
	}
	test_path1 := []string{
		`test_data1.txt`,
	}
	exept_paths := []string{
		`exept_data.txt`,
		`exept_data_u.txt`,
	}

	exept_b := []string{"1. 3 qwwr", "2. 5 wet", "3. 2 qsf", "3. 7 qsag", "6. 1 bbr", "6. 8 uij", "6. 8 uij", "6. 8 uij", "8. 2 ax"}

	testnameK := fmt.Sprintf("Тест номер %d", 1)
	t.Run(testnameK, func(t *testing.T) {
		testKey := CreateSort("0", false, false, false, false, false, test_path)
		err := testKey.Run()
		exept, error := ReadFile(exept_paths[0], false)
		assert.NoError(t, err, error)
		assert.EqualValues(t, exept, testKey.result)
	})

	testnameB := fmt.Sprintf("Тест номер %d", 2)
	t.Run(testnameB, func(t *testing.T) {
		testB := CreateSort("", false, false, false, true, false, test_path)
		err := testB.Run()
		assert.NoError(t, err)
		assert.EqualValues(t, exept_b, testB.result)
	})

	testnameU := fmt.Sprintf("Тест номер %d", 3)
	t.Run(testnameU, func(t *testing.T) {
		testU := CreateSort("", false, false, true, false, false, test_path1)
		err := testU.Run()
		exept, error := ReadFile(exept_paths[1], false)
		assert.NoError(t, err, error)
		assert.EqualValues(t, exept, testU.result)
	})
}
