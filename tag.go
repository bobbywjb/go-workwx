package workwx

// Tag 标签结构体
type Tag struct {
	TagID   int64  `json:"tagid"`
	TagName string `json:"tagname"`
}

// TagUser 标签成员
type TagUser struct {
	UserID string `json:"userid"`
	Name   string `json:"name"`
}

// TagDepartment 标签部门
type TagDepartment struct {
	DepartmentID int64  `json:"departmentid"`
	Name         string `json:"name"`
}

// TagDetail 标签详情
type TagDetail struct {
	TagName      string          `json:"tagname"`
	UserList     []TagUser       `json:"userlist"`
	DepartmentID []TagDepartment `json:"partylist"`
}

// TagCreateRequest 创建标签请求
type TagCreateRequest struct {
	TagID     int64    `json:"tagid,omitempty"`
	TagName   string   `json:"tagname"`
	UserList  []string `json:"userlist,omitempty"`
	PartyList []int64  `json:"partylist,omitempty"`
}

// TagUpdateRequest 更新标签请求
type TagUpdateRequest struct {
	TagID     int64    `json:"tagid"`
	TagName   string   `json:"tagname,omitempty"`
	UserList  []string `json:"userlist,omitempty"`
	PartyList []int64  `json:"partylist,omitempty"`
}

// tagCreateResp 创建标签响应
type tagCreateResp struct {
	respCommon
	TagID int64 `json:"tagid"`
}

// tagListResp 标签列表响应
type tagListResp struct {
	respCommon
	TagList []Tag `json:"taglist"`
}

// tagUserListResp
type tagUserListResp struct {
	respCommon
	TagName   string    `json:"tagname"`
	UserList  []TagUser `json:"userlist,omitempty"`
	PartyList []int64   `json:"partylist,omitempty"`
}

// GetTagList 获取标签列表
func (c *WorkwxApp) GetTagList() ([]Tag, error) {
	resp, err := c.execTagList(reqTagList{})
	if err != nil {
		return nil, err
	}
	return resp.TagList, nil
}

// GetTagUserList 获取标签中包含的成员列表
func (c *WorkwxApp) GetTagUserList(TagID int64) ([]TagUser, error) {
	req := reqTagUserList{TagID: TagID}
	resp, err := c.execTagUserList(req)
	if err != nil {
		return nil, err
	}
	return resp.UserList, nil
}

// GetTagDepartmentList 获取含有该标签用户的部门列表
func (c *WorkwxApp) GetTagDepartmentList(TagID int64) ([]int64, error) {
	req := reqTagUserList{TagID: TagID}
	resp, err := c.execTagUserList(req)
	if err != nil {
		return nil, err
	}
	return resp.PartyList, nil
}
