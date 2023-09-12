function findPeak(a) {
    let start = 1
    let end = a.length-2
    while (start + 1 < end) {
        let mid = start + Math.floor((end-start)/2)
        if (a[mid] < a[mid-1]) {
            end = mid
        } else if (a[mid] < a[mid+1]) {
            start = mid
        } else {
            return mid
        }
    }
    return a[start] > a[end]? start:end
}
