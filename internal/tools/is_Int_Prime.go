package tools

func IsIntPrime(num int) bool {
    if num == 0 || num == 1 {
        return false
    }

    for i:=2; i<num; i++ {
        if num % i == 0 {
             return false
            }
    }
    return true
}

