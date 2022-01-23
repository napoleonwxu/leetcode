func numUniqueEmails(emails []string) int {
    Map := make(map[string]bool)
    for _, email := range emails {
        addr := strings.Split(email, "@")
        local, domain := addr[0], addr[1]
        local = strings.SplitN(local, "+", 2)[0]
        local = strings.Replace(local, ".", "", -1)
        Map[local+"@"+domain] = true
    }
    return len(Map)
}