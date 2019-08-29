/*
Given a year Y and a month M, return how many days there are in that month.
Note:
1583 <= Y <= 2100
1 <= M <= 12
*/

func numberOfDays(Y int, M int) int {
    if M == 2 {
        if Y%400 == 0 || (Y%4 == 0 && Y%100 != 0) {
            return 29
        }
        return 28        
    }
    if M==4 || M==6 || M==9 || M==11 {
        return 30
    }
    return 31
}