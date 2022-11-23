package mygrep

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyGrep(t *testing.T) {

	test_path := []string{
		`test_data.txt`,
	}
	test_path1 := []string{
		`test_data1.txt`,
	}
	// exept_paths := []string{
	// 	`exept_data.txt`,
	// 	`exept_data_u.txt`,
	// }

	exept_A := []string{"4. 5e wetcz\r", "5. 2r qsfv\r", "6. 5f wetz\r"}
	testnameA := fmt.Sprintf("Тест -A номер %d", 1)
	t.Run(testnameA, func(t *testing.T) {
		testA := CreateGrep("2", "", "", false, false, false, false, false, test_path, "4. 5e ")
		err := testA.Run()
		assert.NoError(t, err)
		assert.EqualValues(t, exept_A, testA.result)
	})

	exept_B := []string{"3. 2 qsfv\r", "3. 2 qsfv\r", "4. 5e wetcz\r"}
	testnameB := fmt.Sprintf("Тест -B номер %d", 2)
	t.Run(testnameB, func(t *testing.T) {
		testB := CreateGrep("", "2", "", false, false, false, false, false, test_path, "4. 5e ")
		err := testB.Run()
		assert.NoError(t, err)
		assert.EqualValues(t, exept_B, testB.result)
	})

	exept_C := []string{"3. 2 qsfv\r", "3. 2 qsfv\r", "4. 5e wetcz\r", "5. 2r qsfv\r", "6. 5f wetz\r"}
	testnameC := fmt.Sprintf("Тест -C номер %d", 3)
	t.Run(testnameC, func(t *testing.T) {
		testC := CreateGrep("", "", "2", false, false, false, false, false, test_path, "4. 5e ")
		err := testC.Run()
		assert.NoError(t, err)
		assert.EqualValues(t, exept_C, testC.result)
	})

	exept_c := []string{"2"}
	testnamec := fmt.Sprintf("Тест -c номер %d", 4)
	t.Run(testnamec, func(t *testing.T) {
		testc := CreateGrep("", "", "", true, false, false, false, false, test_path, "3. 2 ")
		err := testc.Run()
		assert.NoError(t, err)
		assert.EqualValues(t, exept_c, testc.result)
	})

	exept_ignore_case := []string{"3. 2 qsfv\r", "3. 2 QsFV\r"}
	testnameIC := fmt.Sprintf("Тест -i номер %d", 5)
	t.Run(testnameIC, func(t *testing.T) {
		testIC := CreateGrep("", "", "", false, true, false, false, false, test_path1, "3. 2 qsfv")
		err := testIC.Run()
		assert.NoError(t, err)
		assert.EqualValues(t, exept_ignore_case, testIC.result)
	})

	exept_invert := []string{"1. 3 qwwrsfb\r", "2. 5 wetgzxc\r", "3. 2 QsFV\r", "4. 5e WWsSA\r", "5. 2r qsfv\r", "6. 5f wetz\r", "7. 25 qsfax\r", "8. 5a wetd\r"}
	testnameV := fmt.Sprintf("Тест -V номер %d", 6)
	t.Run(testnameV, func(t *testing.T) {
		testV := CreateGrep("", "", "", false, false, true, false, false, test_path1, "3. 2 qsfv")
		err := testV.Run()
		assert.NoError(t, err)
		assert.EqualValues(t, exept_invert, testV.result)
	})

	exept_L := []string{"3- 3. 2 qsfv\r", "4- 3. 2 qsfv\r"}
	testnameL := fmt.Sprintf("Тест -L номер %d", 7)
	t.Run(testnameL, func(t *testing.T) {
		testL := CreateGrep("", "", "", false, false, false, false, true, test_path, "3. 2 qsfv\r")
		err := testL.Run()
		assert.NoError(t, err)
		assert.EqualValues(t, exept_L, testL.result)
	})

}
