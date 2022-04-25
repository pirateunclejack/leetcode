package main

////////////  dedumplicateArray
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/

// func removeDuplicates(nums []int) (int, []int) {
// 	for i := len(nums) - 1; i > 0; i-- {
// 		if nums[i] == nums[i-1] {
// 			nums = append(nums[:i], nums[i+1:]...)
// 		}
// 	}
// 	return len(nums), nums[:]
// }

// func main() {
// 	arr := [11]int{1, 1, 2, 2, 4, 4, 5, 6, 6, 6, 6}
// 	a, barr := removeDuplicates(arr[:])
// 	fmt.Println(a)
// 	for _, k := range barr {
// 		fmt.Print(k, ",")
// 	}
// }

//////////////////// get max profit
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/

// func maxProfit(prices []int) int {
// 	var max int

// 	for i := 1; i < len(prices); i++ {
// 		if prices[i]-prices[i-1] > 0 {
// 			max += prices[i] - prices[i-1]
// 		}
// 	}

// 	return max
// }

// func main() {
// 	arr := [6]int{7, 1, 5, 3, 6, 4}
// 	sum := maxProfit(arr[:])
// 	fmt.Println(sum)
// }

/////////////////// rotate array
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/

// func rotate(nums []int, k int) {
// 	leng := len(nums) - k
// 	for j := 1; j <= k; j++ {
// 		for i := 0; i < leng; i++ {
// 			fmt.Println(i)
// 			nums[i], nums[leng] = nums[leng], nums[i]
// 		}
// 	}
// }

// func rotate(nums []int, k int) {
// 	length := len(nums)
// 	k %= length
// 	reverse(nums, 0, length-1)
// 	reverse(nums, 0, k-1)
// 	reverse(nums, k, length-1)

// }

// func reverse(nums []int, start int, end int) {
// 	for start < end {
// 		nums[start], nums[end] = nums[end], nums[start]
// 		start++
// 		end--
// 	}
// }

// func rotate(nums []int, k int) {
// 	hold := nums[0]
// 	var index int
// 	length := len(nums)
// 	var visited []int

// 	for i := 0; i < length; i++ {
// 		index = (index + k) % length
// 		f := func(visited []int, index int) bool {
// 			for i := 0; i < len(visited); i++ {
// 				if visited[i] == index {
// 					return true
// 				}
// 			}
// 			return false
// 		}
// 		if f(visited, index) {
// 			index = (index + 1) % length
// 			hold = nums[index]
// 			i--
// 		} else {
// 			visited = append(visited, index)
// 			nums[index], hold = hold, nums[index]
// 		}
// 	}

// }

// func main() {
// 	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
// 	k := 4

// 	rotate(nums, k)

// 	fmt.Println(nums)
// }

////////////////////// if array have the same member
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x248f5/

// func containsDuplicate(nums []int) bool {
// 	leng := len(nums)
// 	for i := 1; i < leng; i++ {
// 		for j := 0; j < i; j++ {
// 			if nums[i] == nums[j] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func main() {
// 	arr := []int{7, 1, 5, 3, 1, 4}
// 	result := containsDuplicate(arr)
// 	fmt.Println(result)

// }

////////////////// get num only exist onece
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x21ib6/

// func singleNumber(nums []int) int {
// 	for i := 0; i < len(nums); i++ {
// 		var count int
// 		for j := 0; j < len(nums); j++ {
// 			if i == j {
// 				continue
// 			}
// 			if nums[i] == nums[j] {
// 				count++
// 			}
// 		}
// 		if count == 0 {
// 			return nums[i]
// 		}
// 	}
// 	return 0
// }

// func main() {
// 	arr := []int{2, 2, 6, 2, 2, 1, 4, 3, 5, 2}
// 	result := singleNumber(arr)
// 	fmt.Println(result)
// }

/////////////////// get intersection
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2y0c2/

// func intersect(nums1 []int, nums2 []int) []int {
// m1 := make(map[int]int)
// m2 := make(map[int]int)
// var nums []int
// for _, i := range nums1 {
// m1[i]++
// }
// for _, j := range nums2 {
// m2[j]++
// }
// for x := range m1 {
// if m2[x] != 0 {
// if m1[x] < m2[x] {
// for k := 0; k < m1[x]; k++ {
// nums = append(nums, x)
// }
// } else {
// for k := 0; k < m2[x]; k++ {
// nums = append(nums, x)
// }
//
// }
// }
// }
// return nums
// }

// func intersect(nums1 []int, nums2 []int) []int {
// 	sort.Ints(nums1)
// 	sort.Ints(nums2)

// 	var nums []int

// 	i := 0
// 	j := 0
// app:
// 	if nums1[i] == nums2[j] {
// 		nums = append(nums, nums1[i])
// 		i++
// 		j++
// 		if i == len(nums1) {
// 			return nums

// 		}
// 		if j == len(nums2) {
// 			return nums

// 		}
// 		goto app
// 	} else if nums1[i] < nums2[j] {
// 		i++
// 		if i == len(nums1) {
// 			return nums

// 		}
// 		goto app
// 	} else {
// 		j++
// 		if j == len(nums2) {
// 			return nums

// 		}

// 		goto app
// 	}
// }

// func main() {
// 	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
// }

////////// add one
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2cv1c/
// func plusOne(digits []int) []int {

// 	leng := len(digits)

// 	for i := leng - 1; i >= 0; i-- {
// 		if digits[i] != 9 {
// 			digits[i]++
// 			return digits
// 		} else {
// 			digits[i] = 0
// 		}
// 	}

// 	digits = append([]int{1}, digits...)

// 	return digits
// }

// func main() {

// 	result := plusOne([]int{1, 9, 8, 9})
// 	fmt.Println(result)

// }

////////// Move 0
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2ba4i/

// func moveZeroes(nums []int) {

// 	for i := len(nums) - 1; i > 0; i-- {
// 		for j := 0; j < i; j++ {
// 			if nums[j] == 0 && nums[j+1] != 0 {
// 				nums[j], nums[j+1] = nums[j+1], nums[j]
// 			}
// 		}
// 	}
// }

// func moveZeroes(nums []int) {
// 	i := 0
// 	for j := 0; j < len(nums); j++ {
// 		if nums[j] == 0 {
// 			i++
// 		} else if i != 0 {
// 			nums[j-i] = nums[j]
// 			nums[j] = 0
// 		}
// 	}
// }

// func main() {

// 	result := []int{1, 2, 0, 2, 1, 0, 3, 4}

// 	moveZeroes(result)
// 	// nums := []int{1, 9, 8, 9, 0, 0}
// 	// sort.Slice(nums[0:1], func(k, j int) bool {
// 	// 	return nums[k] < nums[j]
// 	// })
// 	fmt.Println(result)

// }

////////// Get sum
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2jrse/

// func twoSum(nums []int, target int) []int {

// 	var f = func(s []int, e int) bool {
// 		for _, a := range s {
// 			if a == e {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	var checked []int
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if f(checked, i) || f(checked, j) {
// 				continue
// 			} else if nums[i]+nums[j] == target {
// 				checked = append(checked, i, j)
// 			}
// 		}
// 	}
// 	return checked
// }

// func main() {

// 	result := []int{1, 2, 0, 2, 1, 0, 3, 4}

// 	get := twoSum(result, 3)
// 	// nums := []int{1, 9, 8, 9, 0, 0}
// 	// sort.Slice(nums[0:1], func(k, j int) bool {
// 	// 	return nums[k] < nums[j]
// 	// })
// 	fmt.Println(get)

// }
////////// check valid Sudoku
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2f9gg/

// func isValidSudoku(board [][]byte) bool {
// 	var prt = func(s []byte){
//         fmt.Println(string(s))
// 	}
// 	var check = func(s []byte) bool {
// 		for i := 0; i < len(s); i++ {
// 			for j := i + 1; j < len(s); j++ {
// 				if s[i] == s[j] {
// 					return false
// 				}
// 			}
// 		}
// 		return true
// 	}
// 	var toCheck []byte

// 	for i := 0; i < 9; i++ {
// 		// row check
// 		for j := 0; j < 9; j++ {

// 			if string(board[i][j]) != "." {
// 				toCheck = append(toCheck, board[i][j])
// 			}

// 		}
// 		prt(toCheck)
// 		if check(toCheck) == false {
// 			return false
// 		}

// 		// column check
// 		toCheck = []byte{}
// 		for j := 0; j < 9; j++ {

// 			if string(board[j][i]) != "." {
// 				toCheck = append(toCheck, board[j][i])
// 			}

// 		}
// 		prt(toCheck)
// 		if check(toCheck) == false {
// 			return false
// 		}
// 		toCheck = []byte{}
// 	}

// 	for i := 0; i < 9; i += 3 {
// 		// box check
// 		for j := 0; j < 9; j += 3 {
// 			for k := 0; k < 3; k++ {
// 				for l := 0; l < 3; l++ {
// 					if string(board[i+k][j+l]) != "." {
// 						toCheck = append(toCheck, board[i+k][j+l])
// 					}
// 				}
// 			}
// 			prt(toCheck)
// 			if check(toCheck) == false {
// 				return false
// 			}
//             toCheck = []byte{}
// 		}

// 		for j := 0; j < 9; j += 3 {
// 			for k := 0; k < 3; k++ {
// 				for l := 0; l < 3; l++ {
// 					if string(board[j+k][i+l]) != "." {
// 						toCheck = append(toCheck, board[j+k][i+l])
// 					}
// 				}
// 			}
// 			prt(toCheck)
// 			if check(toCheck) == false {
// 				return false
// 			}
//             toCheck = []byte{}
// 		}
// 	}
// 	return true
// }

////////// rotate picture
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnhhkv/

// func rotate(matrix [][]int) {
// 	leng := len(matrix)
// 	for i := 0; i < leng; i++ {
// 		for j := 0; j < (leng+1)/2; j++ {
// 			matrix[i][j], matrix[i][leng-1-j] = matrix[i][leng-1-j], matrix[i][j]
// 		}
// 	}
// 	for i := leng - 2; i >= 0; i-- {
// 		for j := 0; j <= leng-i-2; j++ {
// 			matrix[i][j], matrix[leng-j-1][leng-i-1] = matrix[leng-j-1][leng-i-1], matrix[i][j]
// 		}
// 	}
// }

// func main() {
// 	// var matrix [][]int
// 	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

// 	rotate(matrix)
// 	fmt.Print(matrix)
// }

////////// reverse string
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnhbqj/

// func reverseString(s []byte)  {
// 	for i:=0;i< (len(s)+1)/2; i++{
// 		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
// 	}
// }

////////// reverse int
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnx13t/

// func reverse(x int) int {
// 	var result int
// 	var temp int = x
// 	for i := 1; ; i++ {
// 		remainder := temp % 10
// 		result = result*10 + remainder
// 		temp = temp / 10
// 		if temp == 0 {
// 			break
// 		}
// 	}

// 	if result > 2147483647 || result < -2147483648 {
// 		return 0
// 	}

// 	return result
// }

// func main() {
// 	var intNum int = -1230012
// 	result := reverse(intNum)
// 	fmt.Println(result)
// }

////////// first uniq char
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn5z8r/

// func firstUniqChar(s string) int {
// 	if len(s) == 1 {
// 		return 0
// 	}
// 	var alreadyExist []byte
// 	var f = func(alalreadyExist []byte, b byte) bool {
// 		for k := 0; k < len(alreadyExist); k++ {
// 			if b == alreadyExist[k] {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	for i := 0; i < len(s); i++ {
// 		if f(alreadyExist, s[i]) {
// 			continue
// 		}
// 		if i == len(s)-1 {
// 			return i
// 		}
// 		for j := i + 1; j < len(s); j++ {
// 			if s[i] == s[j] {
// 				alreadyExist = append(alreadyExist, s[i])
// 				break
// 			} else if j == len(s)-1 {
// 				return i
// 			}
// 		}
// 	}
// 	return -1
// }

// func main() {
// 	var str string = "loveleetcode"
// 	// var str2 []string = string(str)
// 	fmt.Println(firstUniqChar(str))

// }

////////// isAnagram
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn96us/

// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t){
// 		return false
// 	} else {
// 		m1 := make(map[byte]int)
// 		m2 := make(map[byte]int)
// 		leng := len(s)
// 		for i:=0;i<leng;i++{
// 			m1[s[i]]++
// 			m2[t[i]]++
// 		}
// 		if len(m1) == len(m2) {
// 			for k,_ := range m1 {
// 				if m1[k] != m2[k]{
// 					return false
// 				}
// 			}
// 			return true
// 		} else {
// 			return false
// 		}
// 	}
// }

////////// isPalindrome
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xne8id/

// func isPalindrome(s string) bool {
// 	var arr []rune
// 	for _, ch := range s {
// 		if ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' {
// 			arr = append(arr, ch)
// 		} else if ch >= 'A' && ch <= 'Z' {
// 			arr = append(arr, ch+('a'-'A'))
// 		} else {
// 			continue
// 		}
// 	}

// 	for i := 0; i < (len(arr)+1)/2; i++ {
// 		if arr[i] != arr[len(arr)-1-i] {
// 			return false
// 		}
// 	}
// 	return true

// }

// func main() {
// 	var s string = "`l;`` 1o1 ??;l`"
// 	fmt.Println(isPalindrome(s))
// }

////////// myAtoi
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnoilh/

// func myAtoi(s string) int {
// 	// remove space
// 	str := strings.TrimSpace(s)
// 	l := len(str)
// 	if l == 0 {
// 		return 0
// 	}
// 	if !isCode(str[0]) && !isDigit(str[0]) {
// 		return 0
// 	}

// 	f := 1
// 	if isCode(str[0]) {
// 		if str[0] == '-' {
// 			f = -1
// 		}
// 		str = str[1:]
// 		l = len(str)
// 	}
// 	res := 0
// 	for i := 0; i < l; i++ {
// 		if !isDigit(str[i]) {
// 			break
// 		}
// 		res += int(str[i] - '0')
// 		if f*res > math.MaxInt32 {
// 			return math.MaxInt32
// 		}
// 		if f*res < math.MinInt32 {
// 			return math.MinInt32
// 		}
// 		res *= 10

// 	}
// 	res = f * res / 10
// 	return res
// }

// func isDigit(c uint8) bool {
// 	return '0' <= c && c <= '9'
// }
// func isCode(c uint8) bool {
// 	return c == '+' || c == '-'
// }

////////// strStr
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnr003/

// func strStr(haystack string, needle string) int {
// 	if len(needle) == 0 && len(haystack) == 0 {
// 		return 0
// 	}
// 	if len(needle) == 0 {
// 		return 0
// 	}

// 	var start int
// 	for i := 0; i < len(haystack)-len(needle)+1; i++ {
// 		if haystack[i] == needle[0] {
// 			start = i
// 			for j := 0; j < len(needle); j++ {
// 				if haystack[i+j] != needle[j] {
// 					break
// 				} else if j == len(needle)-1 {
// 					return start
// 				}
// 			}
// 		}
// 	}
// 	return -1
// }

////////// countAndSay
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnpvdm/

// func countAndSay(n int) string {
// 	var res int = 1

// 	for i := 1; i <= n; i++ {
// 		if i == 1 {
// 			res = 1
// 		} else {
// 			res = getRes(res)

// 		}
// 	}
// 	return strconv.Itoa(res)
// }
// func getRes(n int) int {

// 	var num []int
// 	var count []int
// 	// num := make([]int, 10)
// 	// count := make([]int, 10)
// 	num = append(num, n%10)
// 	count = append(count, 1)
// 	n = n / 10

// 	for i := 0; ; i++ {
// 		if n == 0 {
// 			break
// 		}
// 		for j := 0; ; j++ {
// 			if n%10 == num[i] {
// 				count[i]++
// 				n = n / 10
// 			} else {
// 				num = append(num, n%10)
// 				// num[i+1] = n % 10
// 				count = append(count, 1)
// 				// count[i+1] = 1
// 				n = n / 10
// 				break
// 			}
// 			if n == 0 {
// 				break
// 			}
// 		}
// 	}
// 	// fmt.Println(num)
// 	// fmt.Println(count)

// 	var res int
// 	for k := len(num) - 1; k >= 0; k-- {
// 		res = res*100 + count[k]*10 + num[k]
// 	}

// 	return res
// }
///////////////////////////////
// func countAndSay(n int) string {

// 	var arr string = "1"
// 	if n == 1 {
// 		return arr
// 	}

// 	for i := 1; i < n; i++ {
// 		word := arr[0] // current number
// 		num := 1       // count
// 		temp := ""     // string

// 		for j := 1; j < len(arr); j++ {
// 			if word == arr[j] {
// 				num++
// 			} else if word != arr[j] {
// 				num_char := num
// 				temp += strconv.Itoa(num_char)
// 				temp += string(word)

// 				word = arr[j]
// 				num = 1
// 			}
// 		}
// 		num_char := num
// 		temp += strconv.Itoa(num_char)
// 		temp += string(word)

// 		arr = temp
// 	}
// 	return arr

// }

// func main() {

// 	var n int = 4
// 	res := countAndSay(n)
// 	fmt.Println(res)

// }

////////// longestCommonPrefix
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnmav1/

// func longestCommonPrefix(strs []string) string {

// 	var resSlice []byte
// 	if len(strs) == 1 {
// 		return strs[0]
// 	}
// 	length := len(strs[0])
// 	// var shortest int
// 	for k := 0; k < len(strs); k++ {
// 		if length > len(strs[k]) {
// 			length = len(strs[k])
// 			// shortest = k
// 		}
// 	}
// 	for i := 0; i < length; i++ {
// 		for j := 0; j < len(strs)-1; j++ {
// 			if strs[j][i] != strs[j+1][i] {
// 				return string(resSlice)
// 			} else if j == len(strs)-2 {
// 				resSlice = append(resSlice, strs[j][i])
// 			}
// 		}
// 	}
// 	return string(resSlice)
// }

// func main() {
// 	var strs = []string{"flower", "flow", "flight"}
// 	result := longestCommonPrefix(strs)
// 	fmt.Println(result)

// }

////////// deleteNode
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnarn7/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func deleteNode(node *ListNode) {
//     node.Val = node.Next.Val
// 	node.Next = node.Next.Next
// }

////////// removeNthFromEnd
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn2925/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	var fast *ListNode = head
// 	var slow *ListNode = head
// 	for i := 0; i < n; i++ {
// 		fast = fast.Next
// 	}
// 	if fast == nil {
// 		return head.Next
// 	}
// 	for {
// 		if fast.Next != nil {
// 			fast = fast.Next
// 			slow = slow.Next
// 		} else {
// 			break
// 		}
// 	}
// 	slow.Next = slow.Next.Next
// 	return head
// }

////////// reverseList
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnnhm6/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	var newHead *ListNode

// 	for head != nil {
// 		node := head.Next
// 		head.Next = newHead
// 		newHead = head
// 		head = node
// 	}
// 	return newHead

// }

// func reverseList(head *ListNode) *ListNode {
// 	var newHead *ListNode

// 	for head != nil {
// 		head, head.Next, newHead = head.Next, newHead, head
// 	}
// }
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	ret := reverseList(head.Next)
// 	head.Next.Next = head
// 	head.Next = nil
// 	return ret
// }

// func main() {

// 	L5 := ListNode{
// 		Val:  5,
// 		Next: nil,
// 	}
// 	L4 := ListNode{
// 		Val:  4,
// 		Next: &L5,
// 	}
// 	L3 := ListNode{
// 		Val:  3,
// 		Next: &L4,
// 	}
// 	L2 := ListNode{
// 		Val:  2,
// 		Next: &L3,
// 	}
// 	L1 := ListNode{
// 		Val:  1,
// 		Next: &L2,
// 	}

// 	result := reverseList(&L1)
// 	for result != nil {
// 		fmt.Println(result.Val)
// 		result = result.Next
// 	}
// }

////////// mergeTwoLists
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnnbp2/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

// 	if list1 == nil {
// 		return list2
// 	}

// 	if list2 == nil {
// 		return list1
// 	}

// 	var head, node *ListNode

// 	if list1.Val < list2.Val {
// 		head = list1
// 		node = list1
// 		list1 = list1.Next
// 	} else {
// 		head = list2
// 		node = list2
// 		list2 = list2.Next
// 	}

// 	for list1 != nil && list2 != nil {
// 		if list1.Val < list2.Val {
// 			node.Next = list1
// 			list1 = list1.Next
// 		} else {
// 			node.Next = list2
// 			list2 = list2.Next
// 		}
// 		node = node.Next
// 	}

// 	if list1 != nil {
// 		node.Next = list1
// 	}
// 	if list2 != nil {
// 		node.Next = list2
// 	}

// 	return head

// }

// func main() {
// 	//
// 	L5 := ListNode{
// 		Val:  5,
// 		Next: nil,
// 	}
// 	L4 := ListNode{
// 		Val:  4,
// 		Next: &L5,
// 	}
// 	L3 := ListNode{
// 		Val:  3,
// 		Next: &L4,
// 	}
// 	L2 := ListNode{
// 		Val:  2,
// 		Next: &L3,
// 	}
// 	L1 := ListNode{
// 		Val:  1,
// 		Next: &L2,
// 	}

// 	M2 := ListNode{
// 		Val:  2,
// 		Next: nil,
// 	}
// 	M1 := ListNode{
// 		Val:  1,
// 		Next: &M2,
// 	}
// 	//
// 	result := mergeTwoLists(&L1, &M1)
// 	for result != nil {
// 		fmt.Println(result.Val)
// 		result = result.Next
// 	}
// }

//////////