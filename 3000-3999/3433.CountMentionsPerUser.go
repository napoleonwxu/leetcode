type EVENT_TYPE int

const (
    ONLINE EVENT_TYPE = 0
    OFFLINE EVENT_TYPE = 1
    MESSAGE EVENT_TYPE = 2
)

type event struct {
    ev EVENT_TYPE
    ts int
    id string
}

func countMentions(numberOfUsers int, events [][]string) []int {
    evs := make([]event, 0, len(events))
    for _, e := range events {
        ev := MESSAGE
        if e[0] == "OFFLINE" {
            ev = OFFLINE
        }
        ts, _ := strconv.Atoi(e[1])
        evs = append(evs, event{
            ev: ev,
            ts: ts,
            id: e[2],
        })
        if e[0] == "OFFLINE" {
            evs = append(evs, event{
                ev: ONLINE,
                ts: ts+60,
                id: e[2],
            })
        }
    }
    sort.Slice(evs, func(i, j int) bool {
        if evs[i].ts == evs[j].ts {
            return evs[i].ev < evs[j].ev
        }
        return evs[i].ts < evs[j].ts
    })
    offline := make([]bool, numberOfUsers)
    ans := make([]int, numberOfUsers)
    all := 0
    for _, e := range evs {
        if e.ev == MESSAGE {
            if e.id == "ALL" {
                all++
            } else if e.id == "HERE" {
                for i, st := range offline {
                    if !st {
                        ans[i]++
                    }
                }
            } else {
                ids := strings.Split(e.id, " ")
                for _, id := range ids {
                    i, _ := strconv.Atoi(id[2:])
                    ans[i]++
                }
            }
        } else if e.ev == OFFLINE {
            i, _ := strconv.Atoi(e.id)
            offline[i] = true
        } else {
            i, _ := strconv.Atoi(e.id)
            offline[i] = false
        }
    }
    for i := range ans {
        ans[i] += all
    }
    return ans
}
