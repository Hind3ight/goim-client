//
//	@Description
//	@return
//  @author hind3ight
//  @createdtime
//  @updatedtime

package ws

type TokenStruct struct {
	Mid      int    `json:"mid"`
	RoomId   string `json:"room_id"`
	Platform string `json:"platform"`
	Accepts  []int  `json:"accepts"`
}
