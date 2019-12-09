// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package gen

import (
	json "encoding/json"
	humus "github.com/Vliro/humus"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD2b7633eDecodeGithubComVliroHumusTesting(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "User.name":
			out.Name = string(in.String())
		case "User.email":
			out.Email = string(in.String())
		case "uid":
			out.Uid = humus.UID(in.String())
		case "dgraph.type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				in.Delim('[')
				if out.Type == nil {
					if !in.IsDelim(']') {
						out.Type = make([]string, 0, 4)
					} else {
						out.Type = []string{}
					}
				} else {
					out.Type = (out.Type)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Type = append(out.Type, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComVliroHumusTesting(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"User.name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.Email != "" {
		const prefix string = ",\"User.email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	if in.Uid != "" {
		const prefix string = ",\"uid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Uid))
	}
	if len(in.Type) != 0 {
		const prefix string = ",\"dgraph.type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Type {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComVliroHumusTesting(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComVliroHumusTesting(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComVliroHumusTesting(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComVliroHumusTesting(l, v)
}
func easyjsonD2b7633eDecodeGithubComVliroHumusTesting1(in *jlexer.Lexer, out *Question) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Question.title":
			out.Title = string(in.String())
		case "Question.from":
			if in.IsNull() {
				in.Skip()
				out.From = nil
			} else {
				if out.From == nil {
					out.From = new(User)
				}
				(*out.From).UnmarshalEasyJSON(in)
			}
		case "Question.comments":
			if in.IsNull() {
				in.Skip()
				out.Comments = nil
			} else {
				in.Delim('[')
				if out.Comments == nil {
					if !in.IsDelim(']') {
						out.Comments = make([]*Comment, 0, 8)
					} else {
						out.Comments = []*Comment{}
					}
				} else {
					out.Comments = (out.Comments)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *Comment
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(Comment)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Comments = append(out.Comments, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Post.text":
			out.Text = string(in.String())
		case "Post.datePublished":
			if in.IsNull() {
				in.Skip()
				out.DatePublished = nil
			} else {
				if out.DatePublished == nil {
					out.DatePublished = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.DatePublished).UnmarshalJSON(data))
				}
			}
		case "uid":
			out.Uid = humus.UID(in.String())
		case "dgraph.type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				in.Delim('[')
				if out.Type == nil {
					if !in.IsDelim(']') {
						out.Type = make([]string, 0, 4)
					} else {
						out.Type = []string{}
					}
				} else {
					out.Type = (out.Type)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.Type = append(out.Type, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComVliroHumusTesting1(out *jwriter.Writer, in Question) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Title != "" {
		const prefix string = ",\"Question.title\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	if in.From != nil {
		const prefix string = ",\"Question.from\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.From).MarshalEasyJSON(out)
	}
	if len(in.Comments) != 0 {
		const prefix string = ",\"Question.comments\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v6, v7 := range in.Comments {
				if v6 > 0 {
					out.RawByte(',')
				}
				if v7 == nil {
					out.RawString("null")
				} else {
					(*v7).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Text != "" {
		const prefix string = ",\"Post.text\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Text))
	}
	if in.DatePublished != nil {
		const prefix string = ",\"Post.datePublished\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.DatePublished).MarshalJSON())
	}
	if in.Uid != "" {
		const prefix string = ",\"uid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Uid))
	}
	if len(in.Type) != 0 {
		const prefix string = ",\"dgraph.type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.Type {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Question) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComVliroHumusTesting1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Question) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComVliroHumusTesting1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Question) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComVliroHumusTesting1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Question) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComVliroHumusTesting1(l, v)
}
func easyjsonD2b7633eDecodeGithubComVliroHumusTesting2(in *jlexer.Lexer, out *Post) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Post.text":
			out.Text = string(in.String())
		case "Post.datePublished":
			if in.IsNull() {
				in.Skip()
				out.DatePublished = nil
			} else {
				if out.DatePublished == nil {
					out.DatePublished = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.DatePublished).UnmarshalJSON(data))
				}
			}
		case "uid":
			out.Uid = humus.UID(in.String())
		case "dgraph.type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				in.Delim('[')
				if out.Type == nil {
					if !in.IsDelim(']') {
						out.Type = make([]string, 0, 4)
					} else {
						out.Type = []string{}
					}
				} else {
					out.Type = (out.Type)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.Type = append(out.Type, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComVliroHumusTesting2(out *jwriter.Writer, in Post) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Text != "" {
		const prefix string = ",\"Post.text\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Text))
	}
	if in.DatePublished != nil {
		const prefix string = ",\"Post.datePublished\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.DatePublished).MarshalJSON())
	}
	if in.Uid != "" {
		const prefix string = ",\"uid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Uid))
	}
	if len(in.Type) != 0 {
		const prefix string = ",\"dgraph.type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.Type {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Post) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComVliroHumusTesting2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Post) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComVliroHumusTesting2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Post) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComVliroHumusTesting2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Post) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComVliroHumusTesting2(l, v)
}
func easyjsonD2b7633eDecodeGithubComVliroHumusTesting3(in *jlexer.Lexer, out *Error) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Error.message":
			out.Message = string(in.String())
		case "Error.errorType":
			out.ErrorType = string(in.String())
		case "Error.time":
			if in.IsNull() {
				in.Skip()
				out.Time = nil
			} else {
				if out.Time == nil {
					out.Time = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Time).UnmarshalJSON(data))
				}
			}
		case "uid":
			out.Uid = humus.UID(in.String())
		case "dgraph.type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				in.Delim('[')
				if out.Type == nil {
					if !in.IsDelim(']') {
						out.Type = make([]string, 0, 4)
					} else {
						out.Type = []string{}
					}
				} else {
					out.Type = (out.Type)[:0]
				}
				for !in.IsDelim(']') {
					var v13 string
					v13 = string(in.String())
					out.Type = append(out.Type, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComVliroHumusTesting3(out *jwriter.Writer, in Error) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Message != "" {
		const prefix string = ",\"Error.message\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Message))
	}
	if in.ErrorType != "" {
		const prefix string = ",\"Error.errorType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ErrorType))
	}
	if in.Time != nil {
		const prefix string = ",\"Error.time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.Time).MarshalJSON())
	}
	if in.Uid != "" {
		const prefix string = ",\"uid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Uid))
	}
	if len(in.Type) != 0 {
		const prefix string = ",\"dgraph.type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v14, v15 := range in.Type {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.String(string(v15))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Error) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComVliroHumusTesting3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Error) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComVliroHumusTesting3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Error) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComVliroHumusTesting3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Error) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComVliroHumusTesting3(l, v)
}
func easyjsonD2b7633eDecodeGithubComVliroHumusTesting4(in *jlexer.Lexer, out *Comment) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Comment.from":
			if in.IsNull() {
				in.Skip()
				out.From = nil
			} else {
				if out.From == nil {
					out.From = new(User)
				}
				(*out.From).UnmarshalEasyJSON(in)
			}
		case "Post.text":
			out.Text = string(in.String())
		case "Post.datePublished":
			if in.IsNull() {
				in.Skip()
				out.DatePublished = nil
			} else {
				if out.DatePublished == nil {
					out.DatePublished = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.DatePublished).UnmarshalJSON(data))
				}
			}
		case "uid":
			out.Uid = humus.UID(in.String())
		case "dgraph.type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				in.Delim('[')
				if out.Type == nil {
					if !in.IsDelim(']') {
						out.Type = make([]string, 0, 4)
					} else {
						out.Type = []string{}
					}
				} else {
					out.Type = (out.Type)[:0]
				}
				for !in.IsDelim(']') {
					var v16 string
					v16 = string(in.String())
					out.Type = append(out.Type, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComVliroHumusTesting4(out *jwriter.Writer, in Comment) {
	out.RawByte('{')
	first := true
	_ = first
	if in.From != nil {
		const prefix string = ",\"Comment.from\":"
		first = false
		out.RawString(prefix[1:])
		(*in.From).MarshalEasyJSON(out)
	}
	if in.Text != "" {
		const prefix string = ",\"Post.text\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Text))
	}
	if in.DatePublished != nil {
		const prefix string = ",\"Post.datePublished\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.DatePublished).MarshalJSON())
	}
	if in.Uid != "" {
		const prefix string = ",\"uid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Uid))
	}
	if len(in.Type) != 0 {
		const prefix string = ",\"dgraph.type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v17, v18 := range in.Type {
				if v17 > 0 {
					out.RawByte(',')
				}
				out.String(string(v18))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Comment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComVliroHumusTesting4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Comment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComVliroHumusTesting4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Comment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComVliroHumusTesting4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Comment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComVliroHumusTesting4(l, v)
}
