// addTwo sums digits in reverse order between two linked lists
// and returns the sum as a linked list.
// Example:
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
package addTwo

import (
	"fmt"
	"strconv"
)

func main() {
	list1 := createList([]int{1, 2, 4})
	list2 := createList([]int{4, 5, 6})
	addTwo(list1, list2)
	//node := addTwo(list1, list2)
	//printList(node)
}

func addTwo(l1 *Node, l2 *Node) *Node {
	str1 := makeString(l1)
	str2 := makeString(l2)
	sum := makeDigit(str1) + makeDigit(str2)
	slc := []int{}
	for sum > 0 {
		slc = append(slc, sum%10)
		sum /= 10
	}
	return createList(slc)
}

func makeString(node *Node) string {
	var str string
	if node.next != nil {
		str = makeString(node.next)
	}
	str += strconv.Itoa(node.value)
	return str
}

func makeDigit(str string) int {
	digit, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return digit
}
