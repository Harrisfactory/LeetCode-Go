import "fmt"

func maxArea(height []int) int {

    var max_area = 0

    var left = 0

    var right = len(height) - 1

    var min_max_height = 0


    for left <= right {

        //calculate minimum maxmimum to avoid tilt
        if height[left] < height[right] {
            min_max_height = height[left]
        } else {
            min_max_height = height[right]
        }

        var area = (right - left) * min_max_height

        if area > max_area {
            max_area = area
        }

        if height[left] < height[right] {
            left++
        } else if height[left] > height[right] {
            right--
        } else {
            left++
            right--
        }
    }

    return max_area
}
