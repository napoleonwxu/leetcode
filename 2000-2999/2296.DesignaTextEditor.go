type TextEditor struct {
    left  []byte // string will TLE
    right []byte
    //text []byte // string will TLE
    //cursor int
}

func Constructor() TextEditor {
    //return TextEditor{left: []byte{}, right: []byte{}}
    return TextEditor{}
}

func (this *TextEditor) AddText(text string) {
    this.left = append(this.left, []byte(text)...)
    //this.text = append(this.text[:this.cursor], append([]byte(text), this.text[this.cursor:]...)...)
    //this.cursor += len(text)
}

func (this *TextEditor) DeleteText(k int) int {
    d := min(k, len(this.left))
    this.left = this.left[:len(this.left)-d]
    //n := min(k, this.cursor)
    //this.text = append(this.text[:this.cursor-n], this.text[this.cursor:]...)
    //this.cursor -= n
    return d
}

func (this *TextEditor) CursorLeft(k int) string {
    n := len(this.left) - min(k, len(this.left))
    //this.right = append(this.left[n:], this.right...) wrong answer
    this.right = append(append([]byte{}, this.left[n:]...), this.right...)
    this.left = this.left[:n]
    return string(this.left[len(this.left)-min(10, len(this.left)):])
    //this.cursor -= min(k, this.cursor)
    //n := min(10, this.cursor)
    //return string(this.text[this.cursor-n:this.cursor])
}

func (this *TextEditor) CursorRight(k int) string {
    m := min(k, len(this.right))
    this.left = append(this.left, this.right[:m]...)
    this.right = this.right[m:]
    return string(this.left[len(this.left)-min(10, len(this.left)):])
    //this.cursor += min(k, len(this.text)-this.cursor)
    //n := min(10, this.cursor)
    //return string(this.text[this.cursor-n:this.cursor])
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

/**
* Your TextEditor object will be instantiated and called as such:
* obj := Constructor();
* obj.AddText(text);
* param_2 := obj.DeleteText(k);
* param_3 := obj.CursorLeft(k);
* param_4 := obj.CursorRight(k);
*/