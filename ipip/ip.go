package ipip

import (
	"encoding/binary"
	"io/ioutil"
	"net"
)

type Ipip struct {
	offset uint32
	index  []byte
	binary []byte
}

func NewIpip() *Ipip {
	return &Ipip{
		offset: 0,
	}
}

func (p *Ipip) Load(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	p.binary = b
	p.offset = binary.BigEndian.Uint32(b[:4])
	p.index = b[4:p.offset]
	return nil
}

func (p *Ipip) Find(ipstr string) (string, error) {
	ip := net.ParseIP(ipstr).To4()
	if ip == nil {
		return na, ErrInvalidIp
	}

	tmp_offset := uint32(ip[0]) * 4
	start := binary.LittleEndian.Uint32(p.index[tmp_offset : tmp_offset+4])

	nip := binary.BigEndian.Uint32(ip)
	var index_offset uint32 = 0
	var index_length uint32 = 0
	var max_comp_len uint32 = p.offset - 1024 - 4
	start = start*8 + 1024

	for start < max_comp_len {
		n := binary.BigEndian.Uint32(p.index[start : start+4])
		if n >= nip {
			tmp_index := []byte{0, 0, 0, 0}
			copy(tmp_index, p.index[start+4:start+7])
			index_offset = binary.LittleEndian.Uint32(tmp_index)
			index_length = uint32(p.index[start+7])
			break
		}
		start += 8
	}

	if index_offset == 0 {
		return na, ErrIpNotFound
	}

	res_offset := p.offset + index_offset - 1024
	var area = make([]byte, index_length)
	copy(area, p.binary[res_offset:res_offset+index_length])
	return bytes2str(area), nil
}
