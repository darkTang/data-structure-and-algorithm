function searchBigSortedArray(reader, target) {
    let rangeTotal = 1
    while (reader[rangeTotal - 1] < target) {
        rangeTotal = rangeTotal * 2
    }
    let start = 0
    let end = rangeTotal - 1
    while (start + 1 < end) {
        let mid = Math.floor(start + (end - start) / 2)
        if (reader[mid] > target) {
            end = mid
        } else if (reader[mid] < target) {
            start = mid
        } else {
            end = mid
        }
    }
    if (reader[start] === target) {
        return start
    }
    if (reader[end] === target) {
        return end
    }
    return -1
}
