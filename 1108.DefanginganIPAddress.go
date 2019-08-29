func defangIPaddr(address string) string {
    return strings.Replace(address, ".", "[.]", -1)
    //return strings.Join(strings.Split(address, "."), "[.]")
}