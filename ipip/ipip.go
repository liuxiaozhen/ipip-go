package ipip

import (
	"encoding/binary"
	"errors"
	"io/ioutil"
	"net"
	"unsafe"
)

var (
	ErrInvalidIp  = errors.New("invalid ip")
	ErrIpNotFound = errors.New("ip not found")
	field_drt     = string("\t")
	cacheIndex    = [65536]uint32{}
)

const (
	na = "N/A"
)

type Ipipx struct {
	offset int
	index  []byte
	binary []byte
}

func NewIpipx() *Ipipx {
	return &Ipipx{
		offset: 0,
	}
}

func (p *Ipipx) Load(path string) error {
	all, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			k := i*256 + j
			cacheIndex[k] = binary.LittleEndian.Uint32(all[(k+1)*4 : (k+1)*4+4])
		}
	}

	p.binary = all
	p.offset = int(binary.BigEndian.Uint32(all[:4]))
	p.index = make([]byte, p.offset-4)
	copy(p.index, all[4:p.offset-4])
	return nil
}

func (p *Ipipx) Find(ipstr string) (string, error) {
	ip := net.ParseIP(ipstr).To4()
	if ip == nil {
		return na, ErrInvalidIp
	}
	nip := binary.BigEndian.Uint32(ip)
	var prefix = int(ip[0])*256 + int(ip[1])
	var start = int(cacheIndex[prefix])
	var maxValue = p.offset - 262144 - 4
	var b = make([]byte, 4)
	var indexOffset = -1
	var indexLength = -1

	for start = start*9 + 262144; start < maxValue; start += 9 {
		tmpInt := binary.BigEndian.Uint32(p.index[start : start+4])
		if tmpInt >= nip {
			b[1] = p.index[start+6]
			b[2] = p.index[start+5]
			b[3] = p.index[start+4]

			indexOffset = int(binary.BigEndian.Uint32(b))
			indexLength = 0xFF&int(p.index[start+7])<<8 + 0xFF&int(p.index[start+8])
			break
		}
	}

	if indexOffset == -1 || indexLength == -1 {
		return na, ErrIpNotFound
	}
	var area = make([]byte, indexLength)
	indexOffset = int(p.offset) + indexOffset - 262144
	copy(area, p.binary[indexOffset:indexOffset+indexLength])
	return bytes2str(area), nil
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
