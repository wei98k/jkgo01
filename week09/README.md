
## 题目

1. 总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用

2. 实现一个从 socket connection 中解码出 goim 协议的解码器。


## 题目1: 解

1. 发送方和接收方规定固定大小的缓冲区，也就是发送和接收都使用固定大小的 byte[] 数组长度，当字符长度不够时使用空字符弥补

2. 以特殊的字符结尾，比如以“\n”结尾，这样我们就知道结束字符，从而避免了半包和粘包问题

3. 在 TCP 协议的基础上封装一层数据请求协议，既将数据包封装成数据头（存储数据正文大小）+ 数据正文的形式，这样在服务端就可以知道每个数据包的具体长度了，知道了发送数据的具体边界之后，就可以解决半包和粘包的问题


## 题目2: 解


```
const (
	MaxBodySize = uint32(1 << 12) 
)

const (
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	_maxPackSize   = MaxBodySize + uint32(_rawHeaderSize)

	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
	_seqOffset    = _opOffset + _opSize
	_bodyOffset   = _seqOffset + _seqSize
)

type Idecoder interface {
	PacketLen() uint32
	HeaderLen() uint16
	Version() uint16
	Operation() uint32
	Sequence() uint32
	Body() []byte
}

type Decoder struct {
	packetLen uint32
	headerLen uint16
	version   uint16
	operation uint32
	sequence  uint32
	body      []byte
}

func Decode(buf []byte) (Idecoder, error) {
	decoder := &Decoder{}

	decoder.packetLen = binary.BigEndian.Uint32(buf[_packOffset : _packOffset+_packSize])
	decoder.headerLen = binary.BigEndian.Uint16(buf[_headerOffset : _headerOffset+_headerSize])
	decoder.version = binary.BigEndian.Uint16(buf[_verOffset : _verOffset+_verSize])
	decoder.operation = binary.BigEndian.Uint32(buf[_opOffset : _opOffset+_opSize])
	decoder.sequence = binary.BigEndian.Uint32(buf[_seqOffset : _seqOffset+_seqSize])

	if decoder.packetLen > _maxPackSize {
		return nil, errors.New("error package length")
	}

	if _bodyLen := int(decoder.packetLen - uint32(decoder.headerLen)); _bodyLen > 0 {
		decoder.body = buf[_bodyOffset : _bodyOffset+_bodyLen]
	}

	return decoder, nil
}

func (d *Decoder) PacketLen() uint32 {
	return d.packetLen
}

func (d *Decoder) HeaderLen() uint16 {
	return d.headerLen
}

func (d *Decoder) Version() uint16 {
	return d.version
}

func (d *Decoder) Operation() uint32 {
	return d.operation
}

func (d *Decoder) Sequence() uint32 {
	return d.sequence
}

func (d *Decoder) Body() []byte {
	if d.body == nil {
		return []byte{}
	} else {
		return d.body
	}
}
```

https://github.com/Terry-Mao/goim/blob/e742c99ad76e626d5f6df8b33bc47ca005501980/api/protocol/protocol.go#L18
