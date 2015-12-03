package stringx

import (
	"github.com/drborges/rivers"
	"github.com/drborges/rivers/stream"
	"strings"
	"regexp"
)

type Pipeline struct {
	*rivers.Pipeline
}

func (pipeline *Pipeline) StartingWith(prefix string) *Pipeline {
	return &Pipeline{pipeline.Take(func(data stream.T) bool {
		return strings.HasPrefix(data.(string), prefix)
	})}
}

func (pipeline *Pipeline) EndingWith(suffix string) *Pipeline {
	return &Pipeline{pipeline.Take(func(data stream.T) bool {
		return strings.HasSuffix(data.(string), suffix)
	})}
}

func (pipeline *Pipeline) Matching(regex string) *Pipeline {
	return &Pipeline{pipeline.Take(func(data stream.T) bool {
		regex := regexp.MustCompile(regex)
		return regex.MatchString(data.(string))
	})}
}

func (pipeline *Pipeline) WithLength(length int) *Pipeline {
	return &Pipeline{pipeline.Take(func(data stream.T) bool {
		return len(data.(string)) == length
	})}
}

func (pipeline *Pipeline) ToLower() *Pipeline {
	return &Pipeline{pipeline.Map(func(data stream.T) stream.T {
		return strings.ToLower(data.(string))
	})}
}

func (pipeline *Pipeline) ToUpper() *Pipeline {
	return &Pipeline{pipeline.Map(func(data stream.T) stream.T {
		return strings.ToUpper(data.(string))
	})}
}

func (pipeline *Pipeline) Title() *Pipeline {
	return &Pipeline{pipeline.Map(func(data stream.T) stream.T {
		return strings.Title(data.(string))
	})}
}

func (pipeline *Pipeline) Replace(old, new string) *Pipeline {
	return &Pipeline{pipeline.Map(func(data stream.T) stream.T {
		return strings.Replace(data.(string), old, new, -1)
	})}
}

func (pipeline *Pipeline) Prepend(prefix string) *Pipeline {
	return &Pipeline{pipeline.Map(func(data stream.T) stream.T {
		return prefix + data.(string)
	})}
}

func (pipeline *Pipeline) Append(suffix string) *Pipeline {
	return &Pipeline{pipeline.Map(func(data stream.T) stream.T {
		return data.(string) + suffix
	})}
}

func (pipeline *Pipeline) Trim() *Pipeline {
	return &Pipeline{pipeline.Map(func(data stream.T) stream.T {
		return strings.TrimSpace(data.(string))
	})}
}

func (pipeline *Pipeline) Split() *Pipeline {
	return &Pipeline{pipeline.OnData(func(data stream.T, emmiter stream.Emitter) {
		for _, c := range strings.Split(data.(string), "") {
			emmiter.Emit(c)
		}
	})}
}

func (pipeline *Pipeline) SplitBy(sep string) *Pipeline {
	return &Pipeline{pipeline.OnData(func(data stream.T, emmiter stream.Emitter) {
		for _, c := range strings.Split(data.(string), sep) {
			emmiter.Emit(c)
		}
	})}
}

func (pipeline *Pipeline) Collect() []string {
	var items []string
	pipeline.Pipeline.CollectAs(&items)
	return items
}

func (pipeline *Pipeline) CollectFirst() string {
	var item string
	pipeline.Pipeline.CollectFirstAs(&item)
	return item
}

func (pipeline *Pipeline) CollectLast() string {
	var item string
	pipeline.Pipeline.CollectLastAs(&item)
	return item
}


func From(slice []string) *Pipeline {
	return &Pipeline{rivers.FromSlice(slice)}
}