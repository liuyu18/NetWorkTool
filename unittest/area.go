package unittest

type box struct {
    length  int
    width   int
    height  int
    name    string
}

// 初始化一个结构体指针对象，后面使用结构体指针方法来设置和获取对象属性
func Newbox() (*box) {
    return &box{}
}

// 给结构体对象设置具体的属性(名称，规格大小)
// 注意: 在如下几个方法中，方法接受者为指针类型，而方法参数为值类型，因此在赋值时可能有人产生疑惑，这里其实是Golang底层做了优化(v.name = name 等同于(*v).name = name)
func (v *box) SetName(name string) {
    v.name = name
}
func (v *box) SetSize(l,w,h int) {
    v.length = l
    v.width = w
    v.height = h
}

// 获取对象的一些属性(名称和体积)
func (v *box) GetName() (string) {
    return v.name
}
func (v *box) GetVolume() (int) {
    return (v.length)*(v.width)*(v.height)
}