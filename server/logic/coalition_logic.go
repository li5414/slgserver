package logic

import "slgserver/server/logic/mgr"


type coalitionLogic struct {

}

func getUnionId(rid int) int {
	attr, ok := mgr.RAttrMgr.Get(rid)
	if ok {
		return attr.UnionId
	}else{
		return 0
	}
}

func getUnionName(unionId int) string {
	if unionId <= 0{
		return ""
	}

	u, ok := mgr.UnionMgr.Get(unionId)
	if ok {
		return u.Name
	}else{
		return ""
	}
}

func getParentId(rid int) int {
	attr, ok := mgr.RAttrMgr.Get(rid)
	if ok {
		return attr.ParentId
	}else{
		return 0
	}
}

func getMainMembers(unionId int) []int {
	u, ok := mgr.UnionMgr.Get(unionId)
	r := make([]int, 0)
	if ok {
		if u.Chairman != 0{
			r = append(r, u.Chairman)
		}
		if u.ViceChairman != 0{
			r = append(r, u.ViceChairman)
		}
	}
	return r
}

func (this* coalitionLogic) MemberEnter(rid, unionId int)  {
	mgr.RAttrMgr.EnterUnion(rid, unionId)

	if rcs, ok := mgr.RCMgr.GetByRId(rid); ok {
		for _, rc := range rcs {
			rc.SyncExecute()
		}
	}
}

func (this* coalitionLogic) MemberExit(rid int) {

	if ra, ok := mgr.RAttrMgr.Get(rid); ok {
		ra.UnionId = 0
	}

	if rcs, ok := mgr.RCMgr.GetByRId(rid); ok {
		for _, rc := range rcs {
			rc.SyncExecute()
		}
	}

}
