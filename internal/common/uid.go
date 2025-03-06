package core

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/pkg/errors"
)

type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

func NewUID(localID uint32, objectType int, shardID uint32) UID {
	return UID{
		localID:    localID,
		objectType: objectType,
		shardID:    shardID,
	}
}

func (uid *UID) GetLocalID() uint32 {
	return uid.localID
}

func (uid *UID) GetShardID() uint32 {
	return uid.shardID
}

func (uid *UID) GetObjectType() int {
	return uid.objectType
}

func (uid *UID) String() string {
	val := uint64(uid.localID)<<28 | uint64(uid.objectType)<<18 | uint64(uid.shardID)<<0
	return base58.Encode([]byte(fmt.Sprintln("%v", val)))
}

func (uid *UID) MarshallJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", uid.String())), nil
}

func (uid *UID) UnmarshalJSON(data []byte) error {
	decodUID, err := UIDFromString(strings.Replace(string(data), "\"", "", -1))
	if err != nil {
		return err
	}

	uid.localID = decodUID.localID
	uid.shardID = decodUID.shardID
	uid.objectType = decodUID.objectType

	return nil
}

func DecomposeUID(str string) (UID, error) {
	uid, err := strconv.ParseInt(str, 10, 24)
	if err != nil {
		return UID{}, err
	}

	if (1 << 18) > uid {
		return UID{}, errors.WithStack(errors.New("Wrong uID"))
	}

	return UID{
		localID:    uint32(uid >> 28),
		objectType: int(uid >> 18 & 0x3FF),
		shardID:    uint32(uid >> 0 & 0x3FFFF),
	}, nil
}

func UIDFromString(str string) (UID, error) {
	return DecomposeUID(string(base58.Decode(str)))
}
