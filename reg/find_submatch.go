package reg

import "regexp"

// FindStringSubmatchMap 将正则表达式中的命名分组和值放到map中
func FindStringSubmatchMap(r *regexp.Regexp, s string) map[string]string {
	captures := make(map[string]string)

	match := r.FindStringSubmatch(s)
	if match == nil {
		return captures
	}

	for i, name := range r.SubexpNames() {
		//Ignore the whole regexp match and unnamed groups
		if i == 0 || name == "" {
			continue
		}
		captures[name] = match[i]
	}
	return captures
}
