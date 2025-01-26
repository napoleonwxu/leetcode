func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
    return (coordinate1[0]+coordinate1[1])%2 == (coordinate2[0]+coordinate2[1])%2
    /*
    x1, y1 := coordinate1[0]-'a', coordinate1[1]-'0'
    x2, y2 := coordinate2[0]-'a', coordinate2[1]-'0'
    return (x1%2 == x2%2 && y1%2 == y2%2) || (x1%2 != x2%2 && y1%2 != y2%2)
    */
}