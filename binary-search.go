func search(nums []int, target int) int {

    var left = 0

    var right = len(nums) - 1

    for left <= right {

        var middle = (left + right)/2

        if nums[middle] == target {
            return middle
        } else if nums[middle] < target {
            left = middle + 1
        } else if nums[middle] > target {
            right = middle - 1
        }
    }

    return -1
}
