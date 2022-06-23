func twoSum(numbers []int, target int) []int {

    var answer[]int

    var i = 0

    var j = len(numbers) - 1

    for i < j {
        if numbers[i] + numbers[j] == target {
            answer = append(answer, i + 1, j + 1)
            return answer
        } else if numbers[i] + numbers[j] > target {
            j--
        } else {
            i++
        }
    }

    return answer
}
