/**
 * @Author: Machao
 * @Date: 2019-12-02 10:25
 * @To:
 */
package ut

import (
	"fmt"
	"testing"
)

func TestNowTimeStr(t *testing.T) {
	fmt.Println(NowTimeStr(Day))
	fmt.Println(NowTimeStr(Hour))
	fmt.Println(NowTimeStr(Minute))
	fmt.Println(NowTimeStr(Second))
}

func TestStrtotime(t *testing.T) {
	timeA, _ := Strtotime("2019-11-30", Day)
	fmt.Println(timeA)
	timeA, _ = Strtotime("2019-11-30 12", Hour)
	fmt.Println(timeA)
	timeA, _ = Strtotime("2019-11-30 12:02", Minute)
	fmt.Println(timeA)
	timeA, _ = Strtotime("2019-11-30 12:02:03", Second)
	fmt.Println(timeA)
}
