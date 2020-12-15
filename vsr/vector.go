package vsr

type Vector struct {
    HashMap map[string]float64
}

func (v *Vector) Add(token string) {
    count, present := v.HashMap[token]
    if present {
        v.HashMap[token] = count + 1
    } else {
        v.HashMap[token] = 1
    }
}


