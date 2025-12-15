package main

func main() {

}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	left1, right1 := 0, len(nums1)-1
	left2, right2 := 0, len(nums2)-1
	for left1 <= right1 && left2 <= right2 {
		mid1 := (right1-left1)/2 + left1
		mid2 := (right2-left2)/2 + left2
	}
}
