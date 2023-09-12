function woodCut(l, k) {
    if (l.length === 0) {
        return 0
    }
    let sum = l.reduce((pre, cur) => pre + cur)
    let start = 1
    let end = Math.min(Math.max(...l), Math.floor(sum / k))
    if (end < 1) {
        return 0
    }
    while (start + 1 < end) {
        let mid = start + Math.floor((end - start) / 2)
        if (getCount(l, mid) >= k) {
            start = mid
        } else {
            end = mid
        }
    }
    return getCount(l, end) >= k ? end : start
}


function getCount(l, length) {
    let sum = 0
    for (let v of l) {
        sum += Math.floor((v / length))
    }
    return sum
}

console.log(woodCut([4, 6, 7, 8], 10))
