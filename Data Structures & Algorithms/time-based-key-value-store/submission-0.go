
type TimestampValue struct {
	value string
	timestamp int
}

type TimeMap struct {
	values map[string][]TimestampValue
}

func Constructor() TimeMap {
	return TimeMap{
		values: make(map[string][]TimestampValue),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.values[key]; !ok {
		this.values[key] = []TimestampValue{}
	}

	this.values[key] = append(this.values[key], TimestampValue{
		value: value,
		timestamp: timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	list, ok := this.values[key]
	if (!ok) {
		return ""
	}

	left := 0
	right := len(list) - 1
	result := ""
	for ;left <= right; {
		mid := (left + right) / 2
		if list[mid].timestamp == timestamp {
			return list[mid].value
		} else if list[mid].timestamp <= timestamp {
			// this timestamp is less than the required time stamp
			// so its value is ok
			result = list[mid].value
			// we still might want to find a bigger timestamp that fits
			left = mid + 1
		} else {
			// this timestamp is not ok, we should look for smaller
			right = mid - 1
		}
	}

	return result
}
