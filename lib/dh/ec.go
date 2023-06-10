package dh

import "github.com/klauspost/reedsolomon"

type EcEncoder struct {
	reedsolomon.Encoder
}

func NewEcEncoder(dataNum int, parityNum int) (EcEncoder, error) {
	encoder, err := reedsolomon.New(dataNum, parityNum)
	if err != nil {
		return EcEncoder{}, err
	}
	return EcEncoder{encoder}, nil
}

func (ec *EcEncoder) AssignData(data []byte, slaveAddrList []string) (map[string][]byte, error) {
	panic("implement me")
}
