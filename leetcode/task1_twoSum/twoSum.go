// twosum returns indices of the two numbers such that they add up to target
package twoSum

func twoSum(nums []int, target int) [2]int {
	var output [2]int
exit:
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[j]+nums[i] {
				output[0] = i
				output[1] = j
				break exit
			}
		}
	}
	return output
}

// func main() {
// 	newNums := [4]int{0, 0, 3, 4}
// 	fmt.Println(twoSum(newNums[:], 3))
// 	fmt.Println(twoSum(newNums[:], 4))
// 	fmt.Println(twoSum(newNums[:], 5))
// 	fmt.Println(twoSum(newNums[:], 7))
// }
